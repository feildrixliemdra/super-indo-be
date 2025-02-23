FROM golang:1.22.5-alpine as builder

ARG GO_BUILD_COMMAND="go build -tags static_all"

# Install some build deps + ssh tools for the setup below.
RUN apk update && apk --no-cache add  build-base  git bash  coreutils openssh  openssl


# this command if you get source from bitbucket repos
# Create the directory where the application will reside
RUN mkdir -p /go/src/github.com/feildrixliemdra/super-indo-be


WORKDIR /go/src/github.com/feildrixliemdra/super-indo-be

COPY . .


# application builder step
RUN go install -v github.com/swaggo/swag/cmd/swag@v1.7.8 && \
    swag init -g internal/router/router.go && \
    go mod download && \
    go mod tidy && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o super-indo-be .

# STEP 2 build a small image
# Set up the final (deployable/runtime) image.
FROM alpine:3.10.2


# setup package dependencies
RUN apk --no-cache update && apk --no-cache  add  ca-certificates bash jq curl

ENV BUILDDIR=/go/src/github.com/feildrixliemdra/super-indo-be
ENV PROJECT_DIR=/opt/github.com/feildrixliemdra/super-indo-be

# Setting timezone
ENV TZ=Asia/Jakarta
RUN apk add -U tzdata
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

#create project directory
RUN mkdir -p $PROJECT_DIR/database/migration

WORKDIR $PROJECT_DIR

COPY --from=builder $BUILDDIR/ $PROJECT_DIR/
COPY --from=builder $BUILDDIR/database/migration $PROJECT_DIR/database/migration


EXPOSE 8080
RUN chmod +x super-indo-be

# Add the current directory to PATH
ENV PATH="$PROJECT_DIR:$PATH"

ENTRYPOINT ["super-indo-be", "serve-http"]


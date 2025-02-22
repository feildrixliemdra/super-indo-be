-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id bigserial PRIMARY KEY ,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT current_timestamp,
    deleted_at timestamp DEFAULT NULL
);
CREATE UNIQUE INDEX idx_users_unique_email ON users(email) WHERE deleted_at IS NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd

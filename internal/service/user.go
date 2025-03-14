package service

import (
	"context"
	"super-indo-be/internal/dto"
	"super-indo-be/internal/errorcustom"
	"super-indo-be/internal/model"
	"super-indo-be/internal/payload"
	"super-indo-be/internal/repository"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Create(ctx context.Context, p payload.CreateUserRequest) (result payload.CreateUserResponse, err error)
	GetByID(ctx context.Context, id uint64) (result payload.GetUserDetailData, err error)
	GetByEmail(ctx context.Context, email string) (result payload.GetUserDetailData, err error)
}

type user struct {
	UserRepository repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &user{
		UserRepository: repo,
	}
}

func (s *user) Create(ctx context.Context, p payload.CreateUserRequest) (result payload.CreateUserResponse, err error) {
	// 1. make sure no duplicate email
	existUser, err := s.UserRepository.GetBy(ctx, model.User{Email: p.Email})
	if err != nil {
		log.Errorf("error get user by email: %v", err)
		return result, err
	}

	if existUser != nil {
		return result, errorcustom.ErrEmailAlreadyRegistered
	}

	// 2. transform create user request payload to user model
	usr := dto.CreateUserPayloadToUserModel(p)

	// 3. hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usr.Password), 10)
	if err != nil {
		log.Errorf("error hash password: %v", err)
		return result, err
	}

	usr.Password = string(hashedPassword)

	// 4. create user
	id, err := s.UserRepository.Create(ctx, usr)
	if err != nil {
		log.Errorf("error create user: %v", err)
		return result, err
	}

	result = payload.CreateUserResponse{
		ID:        id,
		Name:      usr.Name,
		Email:     usr.Email,
		CreatedAt: usr.CreatedAt,
	}

	return result, nil
}

func (s *user) GetByEmail(ctx context.Context, email string) (result payload.GetUserDetailData, err error) {
	usr, err := s.UserRepository.GetBy(ctx, model.User{Email: email})
	if err != nil {
		log.Errorf("error get user by email: %v", err)
		return
	}

	if usr == nil {
		err = errorcustom.ErrUserNotFound
		return
	}

	return dto.UserModelToUserDetailResponse(usr), nil
}

func (s *user) GetByID(ctx context.Context, id uint64) (result payload.GetUserDetailData, err error) {
	usr, err := s.UserRepository.GetBy(ctx, model.User{ID: id})
	if err != nil {
		log.Errorf("error get user by id: %v", err)
		return
	}

	if usr == nil {
		err = errorcustom.ErrUserNotFound
		return
	}

	return dto.UserModelToUserDetailResponse(usr), nil
}

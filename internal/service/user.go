package service

import (
	"context"
	"fmt"
	"super-indo-be/internal/constant"
	"super-indo-be/internal/dto"
	"super-indo-be/internal/model"
	"super-indo-be/internal/payload"
	"super-indo-be/internal/repository"

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
		return result, err
	}

	if existUser != nil {
		fmt.Println(existUser)
		return result, constant.ErrEmailAlreadyRegistered
	}

	// 2. transform create user request payload to user model
	usr := dto.CreateUserPayloadToUserModel(p)

	// 3. hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usr.Password), 10)
	if err != nil {
		return result, err
	}

	usr.Password = string(hashedPassword)

	// 4. create user
	id, err := s.UserRepository.Create(ctx, usr)
	if err != nil {
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
		return
	}

	if usr == nil {
		err = constant.ErrUserNotFound

		return
	}

	return dto.UserModelToUserDetailResponse(usr), nil
}

func (s *user) GetByID(ctx context.Context, id uint64) (result payload.GetUserDetailData, err error) {
	usr, err := s.UserRepository.GetBy(ctx, model.User{ID: id})
	if err != nil {
		return
	}

	if usr == nil {
		err = constant.ErrUserNotFound

		return
	}

	return dto.UserModelToUserDetailResponse(usr), nil
}

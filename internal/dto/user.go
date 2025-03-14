package dto

import (
	"super-indo-be/internal/model"
	"super-indo-be/internal/payload"
)

func CreateUserPayloadToUserModel(p payload.CreateUserRequest) model.User {
	return model.User{
		Name:  p.Name,
		Email: p.Email,
	}
}

func UpdateUserPayloadToUserModel(p payload.UpdateUserRequest) model.User {
	return model.User{
		ID:       p.ID,
		Name:     p.Name,
		Email:    p.Email,
		Password: p.Password,
	}
}

func UserModelToUserDetailResponse(u *model.User) payload.GetUserDetailData {
	return payload.GetUserDetailData{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func UserModelToUserListResponse(u []model.User) []payload.GetUserListData {
	res := []payload.GetUserListData{}

	for i := 0; i < len(u); i++ {
		res = append(res,
			payload.GetUserListData{
				ID:        u[i].ID,
				Name:      u[i].Name,
				Email:     u[i].Email,
				CreatedAt: u[i].CreatedAt,
				UpdatedAt: u[i].UpdatedAt,
			})
	}

	return res
}

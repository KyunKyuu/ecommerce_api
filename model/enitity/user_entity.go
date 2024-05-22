package entity

import "ecommerce_api/model/domain"

type UserEntity struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}

func ToUserEntity(user_id int, name string, email string, role string) UserEntity {
	return UserEntity{
		UserID: user_id,
		Name:   name,
		Email:  email,
		Role:   role,
	}
}

func ToUserListEntity(users []domain.User) []UserEntity {
	userData := []UserEntity{}

	for _, user := range users {
		userData = append(userData, ToUserEntity(user.UserID, user.Name, user.Email, user.Role))
	}

	return userData
}

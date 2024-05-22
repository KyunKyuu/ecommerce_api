package service

import (
	entity "ecommerce_api/model/enitity"
	"ecommerce_api/model/web"
)

type UserService interface {
	SaveUser(request web.UserServiceRequest) (map[string]interface{}, error)
	GetUser(userId int) (entity.UserEntity, error)
	GetUserList() ([]entity.UserEntity, error)
	GetUserDeleted(userId int) (entity.UserEntity, error)
	UpdateUser(request web.UserUpdateServiceRequest, pathId int) (map[string]interface{}, error)
	DeleteUser(userId int) error
	LoginUser(email string, password string) (map[string]interface{}, error)
}

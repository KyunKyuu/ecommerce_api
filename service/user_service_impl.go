package service

import (
	"ecommerce_api/helper"
	"ecommerce_api/model/domain"
	entity "ecommerce_api/model/enitity"
	"ecommerce_api/model/web"
	"ecommerce_api/repository"
	"errors"

	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	repository   repository.UserRepository
	tokenUseCase helper.TokenUseCase
}

func NewUserService(repository repository.UserRepository, token helper.TokenUseCase) *UserServiceImpl {
	return &UserServiceImpl{
		repository:   repository,
		tokenUseCase: token,
	}
}

func (service *UserServiceImpl) SaveUser(request web.UserServiceRequest) (map[string]interface{}, error) {

	passHash, errHash := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)

	if errHash != nil {
		return nil, errHash
	}

	userReq := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Role:     request.Role,
		Password: string(passHash),
	}

	saveUser, errSaveUser := service.repository.SaveUser(userReq)

	if errSaveUser != nil {
		return nil, errSaveUser
	}

	return helper.ResponseToJson{"name": saveUser.Name, "email": saveUser.Email, "role": saveUser.Role}, nil
}

func (service *UserServiceImpl) GetUser(userId int) (entity.UserEntity, error) {
	getUser, errGetUser := service.repository.GetUser(userId)

	if errGetUser != nil {
		return entity.UserEntity{}, errGetUser
	}

	return entity.ToUserEntity(getUser.UserID, getUser.Name, getUser.Email, getUser.Role), nil
}

func (service *UserServiceImpl) GetUserDeleted(userId int) (entity.UserEntity, error) {
	getUser, errGetUser := service.repository.GetUserDeleted(userId)

	if errGetUser != nil {
		return entity.UserEntity{}, errGetUser
	}

	return entity.ToUserEntity(getUser.UserID, getUser.Name, getUser.Email, getUser.Role), nil
}

func (service *UserServiceImpl) GetUserList() ([]entity.UserEntity, error) {
	getUserList, errGetUserList := service.repository.GetUsers()

	if errGetUserList != nil {
		return []entity.UserEntity{}, errGetUserList
	}

	return entity.ToUserListEntity(getUserList), nil
}

func (service *UserServiceImpl) UpdateUser(request web.UserUpdateServiceRequest, pathId int) (map[string]interface{}, error) {
	getUserById, err := service.repository.GetUser(pathId)

	if err != nil {
		return nil, err
	}

	if request.Name == "" {
		request.Name = getUserById.Name
	}

	if request.Email == "" {
		request.Email = getUserById.Email
	}

	if request.Role == "" {
		request.Role = getUserById.Role
	}

	userRequest := domain.User{
		UserID:   pathId,
		Name:     request.Name,
		Email:    request.Email,
		Role:     request.Role,
		Password: getUserById.Password,
	}

	updateUser, errUpdate := service.repository.UpdateUser(userRequest)

	if errUpdate != nil {
		return nil, errUpdate
	}

	return helper.ResponseToJson{"name": updateUser.Name, "email": updateUser.Email, "password": updateUser.Password, "role": updateUser.Role}, nil
}

func (service *UserServiceImpl) DeleteUser(userId int) error {
	errDeleteUser := service.repository.DeleteUser(userId)

	if errDeleteUser != nil {
		return errDeleteUser
	}

	return nil
}

func (service *UserServiceImpl) LoginUser(email string, password string) (map[string]interface{}, error) {
	user, err := service.repository.FindUserByEmail(email)

	if err != nil {
		return nil, errors.New("email tidak terdaftar")
	}

	errPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if errPass != nil {
		return nil, errors.New("password Salah")
	}

	expiredTime := time.Now().Local().Add(5 * time.Minute)

	claims := helper.JwtCustomClaims{
		ID:    strconv.Itoa(user.UserID),
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "rest gorm",
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token, errToken := service.tokenUseCase.GenerateAccessToken(claims)

	if errToken != nil {
		return nil, errors.New("ada kesalahan generate token")
	}

	return helper.ResponseToJson{"token": token, "expired": expiredTime}, nil
}

package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/zhuangalbert/boilerplate/src/api/v1/models"
	"github.com/zhuangalbert/boilerplate/src/api/v1/objects"
	"github.com/zhuangalbert/boilerplate/src/api/v1/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository *repositories.UserRepository
}

func UserServiceHandler() *UserService {
	service := &UserService{
		repository: repositories.UserRepositoryHandler(),
	}
	return service
}

func (service *UserService) Store(userData *objects.UserStoreObjectRequest) (int, error) {
	userPassword, err := Hash(userData.Password)
	if err != nil {
		return 0, err
	}
	userMasterData := &models.User{
		ID:       userData.ID,
		Username: userData.Username,
		Password: string(userPassword),
		Phone:    userData.Phone,
		Email:    userData.Email,
	}

	userId, err := service.repository.Store(userMasterData)

	if err != nil {
		return 0, err
	}
	return userId, nil
}

func (service *UserService) GetUserById(userId int) (*objects.UserFindObjectResponse, error) {
	userData, err := service.repository.GetUserById(userId)

	if err != nil {
		return nil, err
	}

	userMasterData := &objects.UserFindObjectResponse{
		ID:        userData.ID,
		Username:  userData.Username,
		Phone:     userData.Phone,
		Email:     userData.Email,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
		DeletedAt: userData.DeletedAt,
	}

	return userMasterData, nil
}

func (service *UserService) Update(userId int, userData *objects.UserUpdateObjectRequest) error {
	userMasterData := &models.User{
		Phone: userData.Phone,
		Email: userData.Email,
	}
	err := service.repository.Update(userId, userMasterData)

	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) Delete(userId int) error {

	err := service.repository.Delete(userId)

	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) Login(userData *objects.UserLoginObjectRequest) (*objects.UserLoginObjectResponse, error) {

	userLoginData, err := service.repository.GetUserByUsername(userData.Username)

	if err != nil {
		return nil, err
	}

	err = VerifyPassword(userLoginData.Password, userData.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, err
	}

	token, err := CreateToken(userLoginData.ID)

	if err != nil {
		return nil, err
	}

	userMasterData := &objects.UserLoginObjectResponse{
		Username:    userData.Username,
		ID:          userLoginData.ID,
		AccessToken: token,
		ExpiredIn:   3600,
	}
	return userMasterData, nil
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func CreateToken(user_id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token valid on 3600 second
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("secret"))
}

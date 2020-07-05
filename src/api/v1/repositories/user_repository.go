package repositories

import (
	"github.com/jinzhu/gorm"
	database "github.com/zhuangalbert/boilerplate/src/api/databases"
	"github.com/zhuangalbert/boilerplate/src/api/v1/models"
)

type UserRepository struct {
	DB        *gorm.DB
	TableName string
}

func UserRepositoryHandler() *UserRepository {
	repository := &UserRepository{DB: database.GetConnection(), TableName: "users"}
	return repository
}

func (repository *UserRepository) Store(user *models.User) (int, error) {
	query := repository.DB.Create(user)
	return int(user.ID), query.Error
}

func (repository *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var userModel models.User
	query := repository.DB.Table(repository.TableName)
	query = query.Where("username=?", username)
	query = query.First(&userModel)

	return &userModel, query.Error
}

func (repository *UserRepository) Update(userId int, userData *models.User) error {
	query := repository.DB.Model(userData)
	query = query.Where("id = ?", userId)
	query = query.Updates(userData)

	return query.Error
}

func (repository *UserRepository) Delete(userId int) error {
	data := repository.DB.Model(&models.User{})
	query := repository.DB.Model(data)
	query = query.Where("id = ?", userId)
	query = query.Delete(data)

	return query.Error
}

func (repository *UserRepository) GetUserById(userId int) (*models.User, error) {
	var userModel models.User
	query := repository.DB.Table(repository.TableName)
	query = query.Where("id=?", userId)
	query = query.First(&userModel)

	return &userModel, query.Error
}

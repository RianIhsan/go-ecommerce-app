package repository

import (
	"errors"
	"log"

	"github.com/RianIhsan/go-ecommerce-app/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(usr domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindById(id uint) (domain.User, error)
	UpdateUser(id uint, usr domain.User) (domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r userRepository) CreateUser(usr domain.User) (domain.User, error) {
	err := r.db.Create(&usr).Error
	if err != nil {
		log.Printf("create user error %v", err)
		return domain.User{}, errors.New("failed creating user")
	}

	return usr, nil
}

func (r userRepository) FindUser(email string) (domain.User, error){
	var user domain.User

	err := r.db.First(&user, "email = ?", email).Error

	if err != nil {
		log.Printf("error find user by email %v", err)
		return domain.User{}, errors.New("user not found")
	}

	return user, nil
}

func (r userRepository) FindById(id uint) (domain.User, error) {
	var user domain.User

	err := r.db.First(&user, "id = ?", id).Error

	if err != nil {
		log.Printf("error find user by id %v", err)
		return domain.User{}, errors.New("user not found")
	}

	return user, nil

}

func (r userRepository) UpdateUser(id uint, usr domain.User) (domain.User, error) {
	

	return domain.User{}, nil
}

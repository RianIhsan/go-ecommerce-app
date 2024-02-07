package service

import (
	"log"

	"github.com/RianIhsan/go-ecommerce-app/internal/domain"
	"github.com/RianIhsan/go-ecommerce-app/internal/dto"
)

type UserService struct {
	// something here
}

func (s UserService) SignUp(input dto.UserSignUp) (string, error) {
	log.Println(input)

	
	return "example-token-hehehe", nil
}

func (s UserService) findUserByEmail(email string) (domain.User, error) {
	// something here
	// business logic
	return domain.User{}, nil
}

func (s UserService) Login(input any) (string, error) {
	// something here
	// business logic
	return "", nil
}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {
	// something here
	// business logic
	return 0, nil
}

func (s UserService) VerifyCode(id uint, code int) error {
	// something here
	// business logic
	return nil
}

func (s UserService) CreateProfile(id uint, input any) error {
	// something here
	// business logic
	return nil
}

func (s UserService) GetProfile(id uint) (*domain.User, error) {
	// something here
	// business logic
	return nil, nil
}

func (s UserService) UpdateProfile(id uint, input any) error {
	// something here
	// business logic
	return nil
}

func (s UserService) BecomeSeller(id uint, input any) (string, error) {
	// something here
	// business logic
	return "", nil
}

func (s UserService) FindCart(id uint) ([]interface{}, error) {
	// something here
	// business logic
	return nil, nil
}

func (s UserService) CreateCart(input any, u domain.User) ([]interface{}, error) {
	// something here
	// business logic
	return nil, nil
}

func (s UserService) CreateOrder(u domain.User) (int, error) {
	// something here
	// business logic
	return 0, nil
}

func (s UserService) GetOrders(u domain.User) ([]interface{}, error) {
	// something here
	// business logic
	return nil, nil
}

func (s UserService) GerOrderById(id uint, uId uint) (interface{}, error) {
	// something here
	// business logic
	return nil, nil
}

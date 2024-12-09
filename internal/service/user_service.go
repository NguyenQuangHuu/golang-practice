package service

import (
	"awesomeProject/internal/model"
	"awesomeProject/internal/repository"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	GetUserByUsername(username string) (*model.User, error)
	PreprocessBeforeSaveUser(user *model.RegisterRequest) error
}

type UserService struct {
	userRepository repository.IUserRepository
}

func (u *UserService) PreprocessBeforeSaveUser(register *model.RegisterRequest) error {
	//TODO implement me
	//Validation
	//handle username exists
	checkUser, _ := u.userRepository.GetUserByUsername(register.Username)
	if checkUser != nil {
		err := errors.New("username is exists")
		return err
	}
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	user := &model.User{
		Username: register.Username,
		Password: string(hashPassword),
	}
	err := u.userRepository.SaveUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) GetUserByUsername(username string) (*model.User, error) {
	//TODO implement me
	result, err := u.userRepository.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{userRepository: userRepository}
}

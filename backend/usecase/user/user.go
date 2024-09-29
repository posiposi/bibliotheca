package user

import (
	"github.com/posiposi/project/backend/domain/model"
	"github.com/posiposi/project/backend/domain/repository/user"
)

type UserUsecase interface {
	GetList() ([]model.User, error)
}

type userUsecase struct {
	userRepository user.UserRepository
}

func NewUserUseCase(userRepository user.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (us *userUsecase) GetList() ([]model.User, error) {
	users, err := us.userRepository.GetList()
	if err != nil {
		return nil, err
	}
	return users, nil
}

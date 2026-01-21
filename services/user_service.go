package services

import (
	"gobackend/domain/dtos"
	"gobackend/domain/interfaces"
	"gobackend/domain/models"

	"github.com/google/uuid"
)

type userService struct {
	repo interfaces.UserRepository
}

func NewUserService(repo interfaces.UserRepository) interfaces.UserService {

	return &userService{repo: repo}

}

func (us *userService) CreateUser(createUserDto *dtos.CreateUserDto) (*dtos.UserResponseDto, error) {
	user := &models.User{
		Name:  createUserDto.Name,
		Email: createUserDto.Email,
	}

	if err := us.repo.CreateUser(user); err != nil {
		return nil, err
	}

	return &dtos.UserResponseDto{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (us *userService) GetUserById(id uuid.UUID) (*dtos.UserResponseDto, error) {

	user, err := us.repo.GetById(id)

	if err != nil {
		return nil, err
	}

	response := &dtos.UserResponseDto{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return response, nil

}

func (us *userService) UpdateUser(id uuid.UUID, updateUserDto *dtos.UpdateUserDto) (*dtos.UserResponseDto, error) {

	user, err := us.repo.GetById(id)

	if err != nil {
		return nil, err
	}

	if updateUserDto.Name != nil {
		user.Name = *updateUserDto.Name
	}
	if updateUserDto.Email != nil {
		user.Email = *updateUserDto.Email
	}

	updated, err := us.repo.UpdateUser(user)

	if err != nil {
		return nil, err
	}

	response := &dtos.UserResponseDto{
		Id:    updated.ID,
		Name:  updated.Name,
		Email: updated.Email,
	}

	return response, nil
}

func (us *userService) DeleteUser(id uuid.UUID) error {

	return us.repo.DeleteUserById(id)

}

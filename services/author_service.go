package services

import (
	"gobackend/domain/dtos"
	"gobackend/domain/interfaces"
	"gobackend/domain/models"

	"github.com/google/uuid"
)

type authorService struct {
	repo interfaces.AuthorRepository
}

func NewAuthorService(repo interfaces.AuthorRepository) interfaces.AuthorService {

	return &authorService{repo: repo}

}

func (us *authorService) CreateAuthor(createAuthorDto *dtos.CreateAuthorDto) (*dtos.AuthorResponseDto, error) {
	author := &models.Author{
		Name: createAuthorDto.Name,
		Bio:  createAuthorDto.Bio,
	}

	if err := us.repo.CreateAuthor(author); err != nil {
		return nil, err
	}

	return &dtos.AuthorResponseDto{
		Id:   author.Id,
		Name: author.Name,
		Bio:  author.Bio,
	}, nil
}

func (us *authorService) GetAuthorById(id uuid.UUID) (*dtos.AuthorResponseDto, error) {

	author, err := us.repo.GetById(id)

	if err != nil {
		return nil, err
	}

	response := &dtos.AuthorResponseDto{
		Id:   author.Id,
		Name: author.Name,
		Bio:  author.Bio,
	}

	return response, nil

}

func (us *authorService) UpdateAuthor(id uuid.UUID, updateAuthorDto *dtos.UpdateAuthorDto) (*dtos.AuthorResponseDto, error) {

	author, err := us.repo.GetById(id)

	if err != nil {
		return nil, err
	}

	if updateAuthorDto.Name != nil {
		author.Name = *updateAuthorDto.Name
	}
	if updateAuthorDto.Bio != nil {
		author.Bio = *updateAuthorDto.Bio
	}

	updated, err := us.repo.UpdateAuthor(author)

	if err != nil {
		return nil, err
	}

	response := &dtos.AuthorResponseDto{
		Id:   updated.Id,
		Name: updated.Name,
		Bio:  updated.Bio,
	}

	return response, nil
}

func (us *authorService) DeleteAuthor(id uuid.UUID) error {

	return us.repo.DeleteAuthor(id)

}

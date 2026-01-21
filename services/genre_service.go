package services

import (
	"gobackend/domain/dtos"
	"gobackend/domain/interfaces"
	"gobackend/domain/models"

	"github.com/google/uuid"
)

type genreService struct {
	repo interfaces.GenreRepository
}

func NewGenreService(repo interfaces.GenreRepository) interfaces.GenreService {

	return &genreService{repo: repo}

}

func (us *genreService) CreateGenre(createGenreDto *dtos.CreateGenreDto) (*dtos.GenreResponseDto, error) {
	genre := &models.Genre{
		Name: createGenreDto.Name,
	}

	if err := us.repo.CreateGenre(genre); err != nil {
		return nil, err
	}

	return &dtos.GenreResponseDto{
		Id:   genre.ID,
		Name: genre.Name,
	}, nil
}

func (us *genreService) GetGenreById(id uuid.UUID) (*dtos.GenreResponseDto, error) {

	genre, err := us.repo.GetById(id)

	if err != nil {
		return nil, err
	}

	response := &dtos.GenreResponseDto{
		Id:   genre.ID,
		Name: genre.Name,
	}

	return response, nil

}

func (us *genreService) UpdateGenre(id uuid.UUID, updateGenreDto *dtos.UpdateGenreDto) (*dtos.GenreResponseDto, error) {

	genre, err := us.repo.GetById(id)

	if err != nil {
		return nil, err
	}

	if updateGenreDto.Name != nil {

		genre.Name = *updateGenreDto.Name
	}

	updated, err := us.repo.UpdateGenre(genre)

	if err != nil {
		return nil, err
	}

	response := &dtos.GenreResponseDto{
		Id:   updated.ID,
		Name: updated.Name,
	}

	return response, nil
}

func (us *genreService) DeleteGenre(id uuid.UUID) error {

	return us.repo.DeleteGenreById(id)

}

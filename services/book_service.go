package services

import (
	"gobackend/domain/dtos"
	"gobackend/domain/interfaces"
	"gobackend/domain/models"

	"github.com/google/uuid"
)

type bookService struct {
	repo interfaces.BookRepository
}

func NewBookService(repo interfaces.BookRepository) interfaces.BookService {

	return &bookService{repo: repo}

}

func (bs *bookService) CreateBook(createBookDto *dtos.CreateBookDTO) (*dtos.BookResponseDTO, error) {
	var genres []models.Genre
	for _, id := range createBookDto.GenreIDs {
		genres = append(genres, models.Genre{ID: id})
	}

	book := &models.Book{
		Title:         createBookDto.Title,
		AuthorID:      createBookDto.AuthorID,
		PublishedYear: createBookDto.PublishedYear,
		Stock:         createBookDto.Stock,
		Genres:        genres,
	}

	if err := bs.repo.CreateBook(book); err != nil {
		return nil, err
	}

	return &dtos.BookResponseDTO{
		ID:            book.ID,
		Title:         book.Title,
		AuthorID:      book.AuthorID,
		PublishedYear: book.PublishedYear,
		Stock:         book.Stock,
	}, nil
}

func (bs *bookService) GetBookById(id uuid.UUID) (*dtos.BookResponseDTO, error) {

	book, err := bs.repo.GetById(id)

	if err != nil {
		return nil, err
	}

	var genreResponses []dtos.GenreResponseDto
	for _, genre := range book.Genres {
		genreResponses = append(genreResponses, dtos.GenreResponseDto{
			Id:   genre.ID,
			Name: genre.Name,
		})
	}

	response := &dtos.BookResponseDTO{
		ID:            book.ID,
		Title:         book.Title,
		AuthorID:      book.AuthorID,
		PublishedYear: book.PublishedYear,
		Stock:         book.Stock,
		Genres:        genreResponses,
	}

	return response, nil

}

func (bs *bookService) UpdateBook(id uuid.UUID, updateBookDto *dtos.UpdateBookDTO) (*dtos.BookResponseDTO, error) {

	book, err := bs.repo.GetById(id)

	if err != nil {

		return nil, err

	}

	if updateBookDto.Title != nil {
		book.Title = *updateBookDto.Title
	}
	if updateBookDto.AuthorID != nil {
		book.AuthorID = *updateBookDto.AuthorID
	}
	if updateBookDto.PublishedYear != nil {
		book.PublishedYear = *updateBookDto.PublishedYear
	}
	if updateBookDto.Stock != nil {
		book.Stock = *updateBookDto.Stock
		if updateBookDto.GenreIDs != nil {
			var genres []models.Genre
			for _, id := range updateBookDto.GenreIDs {
				genres = append(genres, models.Genre{ID: id})
			}
			book.Genres = genres
		}
	}

	updatedBook, err := bs.repo.UpdateBook(book)

	if err != nil {
		return nil, err
	}

	var genreResponses []dtos.GenreResponseDto
	for _, genre := range updatedBook.Genres {
		genreResponses = append(genreResponses, dtos.GenreResponseDto{
			Id:   genre.ID,
			Name: genre.Name,
		})
	}

	response := &dtos.BookResponseDTO{
		ID:            updatedBook.ID,
		Title:         updatedBook.Title,
		AuthorID:      updatedBook.AuthorID,
		PublishedYear: updatedBook.PublishedYear,
		Stock:         updatedBook.Stock,
		Genres:        genreResponses,
	}

	return response, nil
}

func (bs *bookService) DeleteBook(id uuid.UUID) error {
	return bs.repo.DeleteBook(id)
}

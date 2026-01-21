package repositories

import (
	"gobackend/domain/interfaces"
	"gobackend/domain/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) interfaces.BookRepository {
	return &bookRepository{db: db}
}

func (br *bookRepository) CreateBook(book *models.Book) error {

	return br.db.Create(book).Error

}

func (br *bookRepository) UpdateBook(book *models.Book) (*models.Book, error) {

	if err := br.db.Save(&book).Error; err != nil {

		return nil, err
	}

	var updated models.Book

	if err := br.db.Where("id = ?", book.ID).First(&updated).Error; err != nil {

		return nil, err

	}

	return &updated, nil
}

func (br *bookRepository) DeleteBook(id uuid.UUID) error {
	var book models.Book
	if err := br.db.First(&book, id).Error; err != nil {
		return err
	}

	if err := br.db.Select("Genres").Delete(&book).Error; err != nil {
		return err
	}

	return nil
}

func (br *bookRepository) GetById(id uuid.UUID) (*models.Book, error) {

	var book models.Book

	if err := br.db.Where("id = ?", id).First(&book).Error; err != nil {
		return nil, err
	}

	return &book, nil

}

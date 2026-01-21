package handlers

import (
	"gobackend/domain/dtos"
	"gobackend/domain/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GenreHandler struct {
	service interfaces.GenreService
}

func NewGenreHandler(service interfaces.GenreService) *GenreHandler {
	return &GenreHandler{service: service}
}

func (h *GenreHandler) CreateGenre(c *gin.Context) {

	var createGenreDto dtos.CreateGenreDto

	if err := c.ShouldBindJSON(&createGenreDto); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	createdGenre, err := h.service.CreateGenre(&createGenreDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdGenre)

}

func (h *GenreHandler) GetGenreById(c *gin.Context) {

	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	genre, err := h.service.GetGenreById(id)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{"error": "Genre not found"})
		return
	}

	response := dtos.GenreResponseDto{
		Id:   genre.Id,
		Name: genre.Name,
	}

	c.JSON(http.StatusOK, response)

}

func (h *GenreHandler) UpdateGenre(c *gin.Context) {

	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	var updateGenreDto dtos.UpdateGenreDto

	if err := c.ShouldBindJSON(&updateGenreDto); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	respone, err := h.service.UpdateGenre(id, &updateGenreDto)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, respone)

}

func (h *GenreHandler) DeleteGenre(c *gin.Context) {

	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	if err := h.service.DeleteGenre(id); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}

	c.JSON(http.StatusOK, gin.H{"message": "Genre deleted successfully"})

}

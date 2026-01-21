package routes

import (
	"gobackend/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	authorHandler *handlers.AuthorHandler,
	bookHandler *handlers.BookHandler,
	genreHandler *handlers.GenreHandler,
	loanHandler *handlers.LoanHandler,
	userHandler *handlers.UserHandler,
) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api")
	{
		authors := v1.Group("/authors")
		{
			authors.POST("/", authorHandler.CreateAuthor)
			authors.GET("/:id", authorHandler.GetAuthorById)
			authors.PUT("/:id", authorHandler.UpdateAuthor)
			authors.DELETE("/:id", authorHandler.DeleteAuthor)
		}

		books := v1.Group("/books")
		{
			books.POST("/", bookHandler.CreateBook)
			books.GET("/:id", bookHandler.GetBookById)
			books.PUT("/:id", bookHandler.UpdateBook)
			books.DELETE("/:id", bookHandler.DeleteBook)
		}

		genres := v1.Group("/genres")
		{
			genres.POST("/", genreHandler.CreateGenre)
			genres.GET("/:id", genreHandler.GetGenreById)
			genres.PUT("/:id", genreHandler.UpdateGenre)
			genres.DELETE("/:id", genreHandler.DeleteGenre)
		}

		loans := v1.Group("/loans")
		{
			loans.POST("/", loanHandler.CreateLoan)
			loans.GET("/:id", loanHandler.GetLoanById)
			loans.PUT("/:id", loanHandler.UpdateLoan)
			loans.DELETE("/:id", loanHandler.DeleteLoan)
		}

		users := v1.Group("/users")
		{
			users.POST("/", userHandler.CreateUser)
			users.GET("/:id", userHandler.GetById)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}
	}

	return r
}

package routes

import (
	"gobackend/api/handlers"
	"gobackend/api/middleware"
	"gobackend/domain/interfaces"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	authorHandler *handlers.AuthorHandler,
	bookHandler *handlers.BookHandler,
	genreHandler *handlers.GenreHandler,
	loanHandler *handlers.LoanHandler,
	userHandler *handlers.UserHandler,
	authHandler *handlers.AuthHandler,
	tokenService interfaces.TokenService,
) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/refresh-token", authHandler.RefreshToken)
		}

		v1.GET("/authors/:id", authorHandler.GetAuthorById)
		v1.GET("/books/:id", bookHandler.GetBookById)
		v1.GET("/genres/:id", genreHandler.GetGenreById)

		protected := v1.Group("/")
		protected.Use(middleware.AuthMiddleware(tokenService))
		{
			authors := protected.Group("/authors")
			{
				authors.POST("/", authorHandler.CreateAuthor)
				authors.PUT("/:id", authorHandler.UpdateAuthor)
				authors.DELETE("/:id", authorHandler.DeleteAuthor)
			}

			books := protected.Group("/books")
			{
				books.POST("/", bookHandler.CreateBook)
				books.PUT("/:id", bookHandler.UpdateBook)
				books.DELETE("/:id", bookHandler.DeleteBook)
			}

			genres := protected.Group("/genres")
			{
				genres.POST("/", genreHandler.CreateGenre)
				genres.PUT("/:id", genreHandler.UpdateGenre)
				genres.DELETE("/:id", genreHandler.DeleteGenre)
			}

			loans := protected.Group("/loans")
			{
				loans.POST("/", loanHandler.CreateLoan)
				loans.GET("/:id", loanHandler.GetLoanById)
				loans.PUT("/:id", loanHandler.UpdateLoan)
				loans.DELETE("/:id", loanHandler.DeleteLoan)
			}

			users := protected.Group("/users")
			{
				users.POST("/", userHandler.CreateUser)
				users.GET("/:id", userHandler.GetById)
				users.PUT("/:id", userHandler.UpdateUser)
				users.DELETE("/:id", userHandler.DeleteUser)
			}
		}
	}

	return r
}

package main

import (
	"gobackend/api/handlers"
	"gobackend/api/routes"
	dbconfig "gobackend/config"
	"gobackend/domain/models"
	"gobackend/infrastructure/repositories"
	"gobackend/services"
	"log"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println(".env dosyasi bulunamadi, env degiskenleri kontrol edin")
	}

	connection := dbconfig.NewConfig()

	db, err := gorm.Open(postgres.Open(connection.DBURL), &gorm.Config{})

	if err != nil {
		log.Fatal("db bağlanti hatasi: ", err)
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	if err := db.AutoMigrate(
		&models.Author{},
		&models.Book{},
		&models.Genre{},
		&models.Loan{},
		&models.User{},
	); err != nil {
		log.Fatal("Migration hatasi: ", err)
	}

	authorRepo := repositories.NewAuthorRepository(db)
	bookRepo := repositories.NewBookRepository(db)
	genreRepo := repositories.NewGenreRepository(db)
	loanRepo := repositories.NewLoanRepository(db)
	userRepo := repositories.NewUserRepository(db)

	authorService := services.NewAuthorService(authorRepo)
	bookService := services.NewBookService(bookRepo)
	genreService := services.NewGenreService(genreRepo)
	loanService := services.NewLoanService(loanRepo)
	userService := services.NewUserService(userRepo)

	authorHandler := handlers.NewAuthorHandler(authorService)
	bookHandler := handlers.NewBookHandler(bookService)
	genreHandler := handlers.NewGenreHandler(genreService)
	loanHandler := handlers.NewLoanHandler(loanService)
	userHandler := handlers.NewUserHandler(userService)

	r := routes.SetupRouter(
		authorHandler,
		bookHandler,
		genreHandler,
		loanHandler,
		userHandler,
	)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Sunucu başlatilamadi: ", err)
	}

}

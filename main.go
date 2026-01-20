package gobackend

import (
	"fmt"
	dbconfig "gobackend/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	connection := dbconfig.NewConfig()

	db, err := gorm.Open(postgres.Open(connection.DBURL), &gorm.Config{})

	if err != nil {
		log.Fatal("db bağlanti hatasi: ", err)
	}

	fmt.Println("bağlanti başarili:", db)

}

package database

import (
	"fmt"
	"log"
	"server/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	env := fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		config.Config.DB_TYPE,
		config.Config.DB_USERNAME,
		config.Config.DB_PASSWORD,
		config.Config.DB_HOST,
		config.Config.DB_PORT,
		config.Config.DB_NAME,
	)

	fmt.Println(env)

	db, err := gorm.Open(postgres.Open(env), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db

	config.ShowConfigDetails()
}

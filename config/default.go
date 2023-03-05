package config

import (
	"fmt"
	"server/utils"
)

type config struct {
	PORT        string
	SHOW_CONFIG string
	DB_TYPE     string
	DB_PORT     string
	DB_HOST     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
}

var Config config

func SetUpConfig() {
	Config = config{
		PORT:        "3000",
		SHOW_CONFIG: "true",
		DB_TYPE:     "postgres",
		DB_PORT:     "5432",
		DB_HOST:     "localhost",
		DB_USERNAME: "saurabh",
		DB_PASSWORD: "password",
		DB_NAME:     "go-test-1",
	}
}

func ShowConfigDetails() {
	if Config.SHOW_CONFIG == "true" {
		fmt.Println(utils.JSON_Stringify(Config))
		// fmt.Printf("%+v\n", Config)
	} else {
		fmt.Println("Config cannot be available!")
	}
}

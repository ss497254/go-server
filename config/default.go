package config

import (
	"fmt"
	"server/utils"
)

type config struct {
	PORT              string
	SHOW_CONFIG       string
	DB_TYPE           string
	DB_PORT           string
	DB_HOST           string
	DB_USERNAME       string
	DB_PASSWORD       string
	DB_NAME           string
	COOKIE_NAME       string
	COOKIE_SECRET     string
	COOKIE_EXPIRES_IN string
	COOKIE_SECURE     string
	COOKIE_SAMESITE   string
	COOKIE_DOMAIN     string
	COOKIE_MAXAGE     int
}

var Config config

func SetUpConfig() {
	Config = config{
		PORT:            "8080",
		SHOW_CONFIG:     "true",
		DB_TYPE:         "postgres",
		DB_PORT:         "5432",
		DB_HOST:         "localhost",
		DB_USERNAME:     "saurabh",
		DB_PASSWORD:     "password",
		DB_NAME:         "go-test-1",
		COOKIE_NAME:     "fin",
		COOKIE_SECRET:   ")Ff@#$RSAD(*&IFcxR32edfs",
		COOKIE_SAMESITE: "lax",
		COOKIE_DOMAIN:   "localhost",
		COOKIE_MAXAGE:   86400 * 30, // 30 days
		COOKIE_SECURE:   "false",
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

package main

import (
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func db() *gorm.DB {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get env
	MYSQL_USER := getEnv("MYSQL_USER", "root")
	MYSQL_PASSWORD := getEnv("MYSQL_PASSWORD", "")
	MYSQL_HOST := getEnv("MYSQL_HOST", "127.0.0.1")
	MYSQL_PORT := getEnv("MYSQL_PORT", "3306")
	MYSQL_DBNAME := getEnv("MYSQL_DBNAME", "dbname")

	var DB *gorm.DB
	conString := MYSQL_USER + ":" + MYSQL_PASSWORD + "@tcp(" + MYSQL_HOST + ":" + MYSQL_PORT + ")/" + MYSQL_DBNAME
	db, err := gorm.Open(mysql.Open(conString))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Contact{})
	DB = db
	return DB
}
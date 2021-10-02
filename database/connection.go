package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	// for local debugging
	// dsn = "u5z3elx9bcxq6mf8:niaslitks0xa71a2@tcp(uyu7j8yohcwo35j3.cbetxkdyhwsb.us-east-1.rds.amazonaws.com:3306)/iaws6lsi5kzfjyyo?charset=utf8mb4&&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to create db connection")
	}

	return db
}

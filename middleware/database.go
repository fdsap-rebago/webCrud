package middleware

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
	DBErr  error
)

func ConnectDB() {
	dsn := "host=localhost user=postgres password=1Qaz2wsx dbname=web_crud port=5432 sslmode=disable TimeZone=Asia/Manila"
	DBConn, DBErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if DBErr != nil {
		fmt.Println("DB Can't connect")
	}
}

func CreateTable(table interface{}) {

	DBErr = DBConn.AutoMigrate(table)
	if DBErr != nil {
		fmt.Println("Error creating table")
	}
}

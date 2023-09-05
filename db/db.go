package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func GormConnect() {
	// .envを取得、代入
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Fail to read .env file : %v", err)
	}

	DBMS := os.Getenv("DIALECT")
	USER := os.Getenv("USER_NAME")
	PASS := os.Getenv("PASSWORD")
	PROTOCOL := os.Getenv("PROTOCOL")
	DBNAME := os.Getenv("DB_NAME")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err = gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
}

package main

import (
	"fmt"
	"log"
	"metis/application/usecase"
	"metis/infrastructure/repositoryImpl"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("envファイルの読み込みに失敗しました。")
	}
	dbConf := os.Getenv("DB_CONF")
	db, err := gorm.Open(mysql.Open(dbConf), &gorm.Config{})
	if err != nil {
		fmt.Println("DB接続に失敗しました")
	}

	fileStr := os.Getenv("TARGET_FILES")
	files := strings.Split(fileStr, ",")

	lakshmiRepo := repositoryImpl.NewLakshmiRepositoryImpl(db)

	resultRepo := repositoryImpl.NewResultRepositoryImpl(files, 1000)

	lakshmiUsecase := usecase.NewLakshmiUsecase(lakshmiRepo, resultRepo)

	err = lakshmiUsecase.Update()
	if err != nil {
		log.Fatal(err)
	}

}

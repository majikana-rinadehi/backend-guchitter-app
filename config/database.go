package config

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() *gorm.DB {
	env := os.Getenv("GUCHITTER_ENV")

	if env == "production" {
		env = "production"
	} else {
		env = "development"
	}

	godotenv.Load(".env." + env)

	USER := os.Getenv("guchitter_USER")
	PASS := os.Getenv("guchitter_PASS")
	DBNAME := os.Getenv("guchitter_DBNAME")
	HOST := os.Getenv("guchitter_HOST")
	PORT := os.Getenv("guchitter_PORT")
	CONNECT := USER + ":" + PASS + "@tcp(" + HOST + ":" + PORT + ")/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{
		// level infoに設定することでSQLがログ出力される
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("connect failed")
	}

	return db

}

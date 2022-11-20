package config

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() *gorm.DB {

	CONNECT := GetDsn()

	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{
		// level infoに設定することでSQLがログ出力される
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("connect failed")
	}

	return db

}

func GetDsn() string {
	env := os.Getenv("GUCHITTER_ENV")

	if env == "production" {
		env = "production"
	} else {
		env = "development"
	}

	godotenv.Load(".env." + env)

	DBUSER := os.Getenv("guchitter_USER")
	DBPASS := os.Getenv("guchitter_PASS")
	DBNAME := os.Getenv("guchitter_DBNAME")
	HOST := os.Getenv("guchitter_HOST")
	DBPORT := os.Getenv("guchitter_PORT")

	// golang-migrate での migrate 実行時に、下記が設定されていないとエラーになる
	// multiStatements=true
	CONNECT := DBUSER + ":" + DBPASS + "@tcp(" + HOST + ":" + DBPORT + ")/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true"

	return CONNECT
}

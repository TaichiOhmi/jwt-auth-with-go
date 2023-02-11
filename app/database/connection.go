package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/TaichiOhmi/jwt-auth-with-goreact/app/config"
	"github.com/TaichiOhmi/jwt-auth-with-goreact/app/domain"
)

var DB *gorm.DB

func Connect() {
	conf, err := config.New()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Tokyo", conf.DB.Host, conf.DB.User, conf.DB.Password, conf.DB.Name, conf.DB.Port)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("could not connect to the database, %s", err))
	}
	DB = conn
	err = conn.AutoMigrate(&domain.User{})
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
}

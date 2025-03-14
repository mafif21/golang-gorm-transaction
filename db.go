package main

import (
	"fmt"
	"github.com/spf13/viper"
	"golang-transaction-experiment/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConfig struct {
	Host     string
	Username string
	Password string
	DBName   string
	Port     string
}

func (d *DBConfig) Read(viper *viper.Viper) {
	d.Host = viper.GetString("DB_HOST")
	d.Username = viper.GetString("DB_USER")
	d.Password = viper.GetString("DB_PASSWORD")
	d.DBName = viper.GetString("DB_NAME")
	d.Port = viper.GetString("DB_PORT")
}

func OpenConnection(db *DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FJakarta", db.Username, db.Password, db.Host, db.Port, db.DBName)

	openDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})

	if err != nil {
		return nil, err
	}

	fmt.Println(dsn)

	err = openDB.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{})
	if err != nil {
		return nil, err
	}

	return openDB, nil
}

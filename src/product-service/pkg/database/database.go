package database

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectDatabase() error{
	dsn := fmt.Sprintf("skgamebmhszt_root:1010970549abcABC@tcp(42.112.30.39:3306)/skgamebmhszt_manage_stock?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Debug("Failed to connect to database!", err)
		return err
	}
	DB = db

    log.Info("Connect database successfully")
	return nil
}

func CloseDatabase() error {
    sqlDB, err := DB.DB()
    if err != nil {
        log.Println("Failed to get database handle!", err)
        return err
    }

    err = sqlDB.Close()
    if err != nil {
        log.Println("Failed to close database connection!", err)
        return err
    }

    log.Println("Database connection closed!")
    return nil
}
package database

import (
	"fmt"

	"github.com/StephenJonesIT/Manage-Sale-Microservice/src/product-service/common"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectDatabase() error{
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        common.Config.MgDbUserName,
        common.Config.MgDbPassword,
        common.Config.MgAddrs,
        common.Config.MgDbName,
    )
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
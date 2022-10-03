package config

import (
	"fmt"
	"github.com/zakariawahyu/go-hacktiv8/assignment2/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func buildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "masukdb",
		DBName:   "golang_assignment2",
	}

	return &dbConfig
}

func dbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func DatabaseInit() *gorm.DB {
	db, err := gorm.Open(mysql.Open(dbURL(buildDBConfig())), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	defer db.AutoMigrate(
		&entity.Order{},
		&entity.Items{},
	)

	fmt.Println("Database connected")
	return db
}

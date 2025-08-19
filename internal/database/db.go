package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func Connect() {
	username := "shashank"
	password := "sql.Shash2"
	host := "127.0.0.1"
	port := "3306"
	dbname := "leads"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	var err error
	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to MySQL: %v", err)
	}
}

// Optional: close underlying DB
func Close() {
	sqlDB, _ := DBConn.DB()
	sqlDB.Close()
}

package seeds

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}
	defer sqlDB.Close()

	dbManager := NewDBManager(db)
	if err := dbManager.Seed(); err != nil {
		log.Fatalf("Seeder error: %v", err)
	}
}

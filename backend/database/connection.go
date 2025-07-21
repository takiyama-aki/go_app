// database/connection.go
// ------------------------------
// DB 接続やマイグレーションを実施するパッケージ。
// 起動時に Init を呼び出して GORM の *DB を初期化する。
// ------------------------------
package database

import (
	"fmt"
	"log"

	"github.com/takiyama-aki/go_app/config"
	"github.com/takiyama-aki/go_app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init は DB への接続とスキーマ自動マイグレーションを行う
func Init(cfg *config.Config) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 自動マイグレーション
	if err := db.AutoMigrate(
		&models.User{},
		&models.Trade{},
		&models.Account{},
	); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	DB = db
	log.Println("Database connection initialized and migrated")
}

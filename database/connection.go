// database/connection.go
package database

import (
    "fmt"
    "log"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/takiyama-aki/go_app/config"
    "github.com/takiyama-aki/go_app/models"
)

var DB *gorm.DB

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

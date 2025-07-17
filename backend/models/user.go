// models/user.go
package models

import "time"

type User struct {
    ID           uint      `gorm:"primaryKey"`
    Email        string    `gorm:"unique;not null"`
    PasswordHash string    `gorm:"not null"`
    CreatedAt    time.Time
}

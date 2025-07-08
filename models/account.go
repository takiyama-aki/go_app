// models/account.go
package models

import "time"

type Account struct {
    UserID    uint      `gorm:"primaryKey"`
    Balance   float64   `gorm:"not null"`
    UpdatedAt time.Time
}

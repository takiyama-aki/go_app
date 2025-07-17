// models/trade.go
package models

import "time"

type Trade struct {
    ID          uint      `gorm:"primaryKey"`
    UserID      uint      `gorm:"index;not null"`
    Date        time.Time `gorm:"not null"`
    SymbolName  string    `gorm:"not null"`
    SymbolCode  string    `gorm:"not null"`
    Price       float64   `gorm:"not null"`
    Quantity    int       `gorm:"not null;default:100"`
    Side        string    `gorm:"not null"` // "LONG" or "SHORT"
    ProfitLoss  float64
    ManualEntry bool      `gorm:"not null"`
    CreatedAt   time.Time
}

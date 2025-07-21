// models/account.go
package models

// 口座残高を表すモデル

import "time"

type Account struct {
	UserID    uint      `gorm:"primaryKey"` // ユーザーID
	Balance   float64   `gorm:"not null"`   // 残高
	UpdatedAt time.Time // 更新日時
}

// models/trade.go
package models

// 取引履歴を表すモデル

import "time"

type Trade struct {
	ID          uint      `gorm:"primaryKey"`           // 主キー
	UserID      uint      `gorm:"index;not null"`       // ユーザーID
	Date        time.Time `gorm:"not null"`             // 取引日
	SymbolName  string    `gorm:"not null"`             // 銘柄名
	SymbolCode  string    `gorm:"not null"`             // 銘柄コード
	Price       float64   `gorm:"not null"`             // 価格
	Quantity    int       `gorm:"not null;default:100"` // 数量
	Side        string    `gorm:"not null"`             // "LONG" or "SHORT"
	ProfitLoss  float64   // 損益
	ManualEntry bool      `gorm:"not null"` // 手入力かどうか
	CreatedAt   time.Time // 登録日時
}

// models/user.go
package models

// ユーザー情報を表すモデル

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey"`      // 主キー
	Email        string    `gorm:"unique;not null"` // メールアドレス
	PasswordHash string    `gorm:"not null"`        // ハッシュ化されたパスワード
	CreatedAt    time.Time // 登録日時
}

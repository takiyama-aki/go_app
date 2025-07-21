package helpers

// セッションから現在のユーザーIDを取得するユーティリティ

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CurrentUserID はセッションから user_id を取り出して返す
// 型の違いを考慮して uint へ変換する
func CurrentUserID(c *gin.Context) (uint, bool) {
	switch v := sessions.Default(c).Get("user_id").(type) {
	case uint:
		return v, true
	case int:
		return uint(v), true
	case float64: // JSON エンコード経由だと float64
		return uint(v), true
	default:
		return 0, false
	}
}

package middleware

// 認証状態を確認するためのミドルウェア

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// RequireLogin はセッションに user_id が存在するかを確認する
// 無い場合は 401 を返して処理を中断する
func RequireLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		sess := sessions.Default(c)
		if sess.Get("user_id") == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

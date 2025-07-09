package middleware

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
)

// RequireLogin は user_id セッションがなければ 401 を返す
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

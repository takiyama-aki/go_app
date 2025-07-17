package helpers

import (
    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
)

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

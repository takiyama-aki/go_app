package helpers

import "github.com/gin-gonic/gin"

// APIError 統一フォーマット
type APIError struct {
    Code    string `json:"code"`
    Message string `json:"message"`
}

// RespondError : どのハンドラでも呼び出せる
func RespondError(c *gin.Context, status int, code, msg string) {
    c.AbortWithStatusJSON(status, APIError{Code: code, Message: msg})
}

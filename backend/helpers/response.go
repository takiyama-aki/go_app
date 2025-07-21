package helpers

// API レスポンス関連の共通処理

import "github.com/gin-gonic/gin"

// APIError はエラーレスポンスのフォーマットを統一するための構造体
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// RespondError は任意のハンドラから呼び出し可能な汎用エラーレスポンス
func RespondError(c *gin.Context, status int, code, msg string) {
	c.AbortWithStatusJSON(status, APIError{Code: code, Message: msg})
}

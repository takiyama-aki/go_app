package handlers

// アカウント残高に関するハンドラ（サンプル実装）

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetBalance は残高を取得する（現状は固定値）
func GetBalance(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"balance": 0}) }

// UpdateBalance は残高を更新する（ダミー）
func UpdateBalance(c *gin.Context) { c.Status(http.StatusNoContent) }

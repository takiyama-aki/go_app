package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func GetBalance(c *gin.Context)    { c.JSON(http.StatusOK, gin.H{"balance": 0}) }
func UpdateBalance(c *gin.Context) { c.Status(http.StatusNoContent) }

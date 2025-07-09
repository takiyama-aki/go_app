package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func ListTrades(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"trades": []string{}})
}

func CreateTrade(c *gin.Context)    { c.Status(http.StatusCreated) }
func UpdateTrade(c *gin.Context)    { c.Status(http.StatusNoContent) }
func DeleteTrade(c *gin.Context)    { c.Status(http.StatusNoContent) }

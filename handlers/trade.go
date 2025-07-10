package handlers

import (
    "net/http"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"

    "github.com/takiyama-aki/go_app/database"
    "github.com/takiyama-aki/go_app/models"
    // "github.com/takiyama-aki/go_app/helpers"
)

// CreateTradeRequest は取引登録用のリクエストボディ構造体
type CreateTradeRequest struct {
    Date        time.Time `json:"date" binding:"required"`
    SymbolName  string    `json:"symbolName" binding:"required"`
    SymbolCode  string    `json:"symbolCode" binding:"required"`
    Price       float64   `json:"price" binding:"required"`
    Quantity    int       `json:"quantity" binding:"required"`
    Side        string    `json:"side" binding:"required,oneof=LONG SHORT"`
    ProfitLoss  float64   `json:"profitLoss"`
    ManualEntry bool      `json:"manualEntry"`
}

// UpdateTradeRequest は更新用に CreateTradeRequest と同じ
type UpdateTradeRequest = CreateTradeRequest

// ListTradesResponse は取引一覧取得 API のレスポンス
type ListTradesResponse struct {
    Trades []models.Trade `json:"trades"`
}

// ListTrades: GET /trades?month=YYYY-MM
func ListTrades(c *gin.Context) {
    monthStr := c.DefaultQuery("month", time.Now().Format("2006-01"))
    month, err := time.Parse("2006-01", monthStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid month format"})
        return
    }
    start := time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, time.Local)
    end := start.AddDate(0, 1, 0)

    sess := sessions.Default(c)
    uidRaw := sess.Get("user_id")
    uid, ok := uidRaw.(uint)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user id in session"})
        return
    }

    var trades []models.Trade
    if err := database.DB.Where("user_id = ? AND date >= ? AND date < ?", uid, start, end).
        Order("date asc").
        Find(&trades).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch trades"})
        return
    }

    c.JSON(http.StatusOK, ListTradesResponse{Trades: trades})
}

// CreateTrade: POST /trades
func CreateTrade(c *gin.Context) {
    var req CreateTradeRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() + " (required fields: date, symbolName, symbolCode, price, quantity, side)"})
        return
    }

    sess := sessions.Default(c)
    uidRaw := sess.Get("user_id")
    uid, ok := uidRaw.(uint)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user id in session"})
        return
    }

    trade := models.Trade{
        UserID:      uid,
        Date:        req.Date,
        SymbolName:  req.SymbolName,
        SymbolCode:  req.SymbolCode,
        Price:       req.Price,
        Quantity:    req.Quantity,
        Side:        req.Side,
        ProfitLoss:  req.ProfitLoss,
        ManualEntry: req.ManualEntry,
        CreatedAt:   time.Now(),
    }

    if err := database.DB.Create(&trade).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create trade"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"id": trade.ID})
}

// UpdateTrade: PUT /trades/:id
func UpdateTrade(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid trade id"})
        return
    }

    var req UpdateTradeRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()+" - invalid request body"})
        return
    }

    sess := sessions.Default(c)
    uidRaw := sess.Get("user_id")
    uid, ok := uidRaw.(uint)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user id in session"})
        return
    }

    var trade models.Trade
    if err := database.DB.Where("id = ? AND user_id = ?", id, uid).First(&trade).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "trade not found"})
        return
    }

    trade.Date = req.Date
    trade.SymbolName = req.SymbolName
    trade.SymbolCode = req.SymbolCode
    trade.Price = req.Price
    trade.Quantity = req.Quantity
    trade.Side = req.Side
    trade.ProfitLoss = req.ProfitLoss
    trade.ManualEntry = req.ManualEntry

    if err := database.DB.Save(&trade).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update trade"})
        return
    }

    c.Status(http.StatusNoContent)
}

// DeleteTrade: DELETE /trades/:id
func DeleteTrade(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid trade id"})
        return
    }

    sess := sessions.Default(c)
    uidRaw := sess.Get("user_id")
    uid, ok := uidRaw.(uint)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user id in session"})
        return
    }

    if err := database.DB.Where("id = ? AND user_id = ?", id, uid).Delete(&models.Trade{}).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete trade"})
        return
    }

    c.Status(http.StatusNoContent)
}

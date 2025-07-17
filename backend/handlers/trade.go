package handlers

import (
    "net/http"
    "strconv"
    "time"

    //"github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"

    "github.com/takiyama-aki/go_app/database"
    "github.com/takiyama-aki/go_app/helpers"
    "github.com/takiyama-aki/go_app/models"
)

// -------------------- DTO --------------------

type CreateTradeRequest struct {
    Date        time.Time `json:"date" binding:"required"`
    SymbolName  string    `json:"symbolName" binding:"required"`
    SymbolCode  string    `json:"symbolCode" binding:"required"`
    Price       float64   `json:"price" binding:"required,gt=0"`
    Quantity    int       `json:"quantity" binding:"required,min=1"`
    Side        string    `json:"side" binding:"required,oneof=LONG SHORT"`
    ProfitLoss  float64   `json:"profitLoss"`
    ManualEntry bool      `json:"manualEntry"`
}

type UpdateTradeRequest = CreateTradeRequest

type ListTradesResponse struct {
    Trades []models.Trade `json:"trades"`
}

// -------------------- Handlers --------------------

// ListTrades GET /trades?month=YYYY-MM
func ListTrades(c *gin.Context) {
    monthStr := c.DefaultQuery("month", time.Now().Format("2006-01"))
    month, err := time.Parse("2006-01", monthStr)
    if err != nil {
        helpers.RespondError(c, http.StatusBadRequest, "PARAM-001", "invalid month format")
        return
    }
    start := time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, time.Local)
    end := start.AddDate(0, 1, 0)

    uid, ok := helpers.CurrentUserID(c)
    if !ok {
        helpers.RespondError(c, http.StatusUnauthorized, "AUTH-001", "unauthorized")
        return
    }

    var trades []models.Trade
    if err := database.DB.Where("user_id = ? AND date >= ? AND date < ?", uid, start, end).
        Order("date asc").Find(&trades).Error; err != nil {
        helpers.RespondError(c, http.StatusInternalServerError, "DB-001", "failed to fetch trades")
        return
    }

    c.JSON(http.StatusOK, ListTradesResponse{Trades: trades})
}

// CreateTrade POST /trades
func CreateTrade(c *gin.Context) {
    var req CreateTradeRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        helpers.RespondError(c, http.StatusBadRequest, "PARAM-002", err.Error())
        return
    }

    uid, ok := helpers.CurrentUserID(c)
    if !ok {
        helpers.RespondError(c, http.StatusUnauthorized, "AUTH-001", "unauthorized")
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
        helpers.RespondError(c, http.StatusInternalServerError, "DB-002", "failed to create trade")
        return
    }

    c.JSON(http.StatusCreated, gin.H{"id": trade.ID})
}

// UpdateTrade PUT /trades/:id
func UpdateTrade(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 64)
    if err != nil {
        helpers.RespondError(c, http.StatusBadRequest, "PARAM-003", "invalid trade id")
        return
    }

    var req UpdateTradeRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        helpers.RespondError(c, http.StatusBadRequest, "PARAM-002", err.Error())
        return
    }

    uid, ok := helpers.CurrentUserID(c)
    if !ok {
        helpers.RespondError(c, http.StatusUnauthorized, "AUTH-001", "unauthorized")
        return
    }

    var trade models.Trade
    res := database.DB.Where("id = ? AND user_id = ?", id, uid).First(&trade)
    if res.RowsAffected == 0 {
        helpers.RespondError(c, http.StatusNotFound, "RES-404", "trade not found")
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
        helpers.RespondError(c, http.StatusInternalServerError, "DB-003", "failed to update trade")
        return
    }

    c.Status(http.StatusNoContent)
}

// DeleteTrade DELETE /trades/:id
func DeleteTrade(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 64)
    if err != nil {
        helpers.RespondError(c, http.StatusBadRequest, "PARAM-003", "invalid trade id")
        return
    }

    uid, ok := helpers.CurrentUserID(c)
    if !ok {
        helpers.RespondError(c, http.StatusUnauthorized, "AUTH-001", "unauthorized")
        return
    }

    res := database.DB.Where("id = ? AND user_id = ?", id, uid).Delete(&models.Trade{})
    if res.RowsAffected == 0 {
        helpers.RespondError(c, http.StatusNotFound, "RES-404", "trade not found")
        return
    }
    if res.Error != nil {
        helpers.RespondError(c, http.StatusInternalServerError, "DB-004", "failed to delete trade")
        return
    }

    c.Status(http.StatusNoContent)
}

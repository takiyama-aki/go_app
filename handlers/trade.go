// handlers/trade.go
package handlers

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"

    "github.com/takiyama-aki/go_app/database"
    "github.com/takiyama-aki/go_app/models"
)

type ListTradesResponse struct {
    Trades []models.Trade `json:"trades"`
}

// ListTrades : GET /trades?month=YYYY-MM
func ListTrades(c *gin.Context) {
    // 1) month パラメータ取得（無ければ今月）
    monthStr := c.DefaultQuery("month", time.Now().Format("2006-01"))
    // パースして月初・月末の time を作成
    month, err := time.Parse("2006-01", monthStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid month format"})
        return
    }
    start := time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, month.Location())
    end  := start.AddDate(0, 1, 0)

    // 2) セッションから user_id 取得
    sess := sessions.Default(c)
    uid := sess.Get("user_id").(uint)

    // 3) DB クエリ
    var trades []models.Trade
    if err := database.DB.
        Where("user_id = ? AND date >= ? AND date < ?", uid, start, end).
        Order("date asc").
        Find(&trades).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch trades"})
        return
    }

    // 4) レスポンス
    c.JSON(http.StatusOK, ListTradesResponse{Trades: trades})
}

// CreateTrade : POST /trades
func CreateTrade(c *gin.Context) {
  // TODO: バインド → DB登録
  c.Status(http.StatusCreated)
}

// UpdateTrade : PUT /trades/:id
func UpdateTrade(c *gin.Context) {
  // TODO: バインド → DB更新
  c.Status(http.StatusNoContent)
}

// DeleteTrade : DELETE /trades/:id
func DeleteTrade(c *gin.Context) {
  // TODO: ID取得 → DB削除
  c.Status(http.StatusNoContent)
}
// main.go
package main

import (
    "fmt"
    "log"

    "github.com/gin-gonic/gin"
    "goapp/config"
    "goapp/database"
    // "goApp/handlers" など
)

func main() {
    // 1. 設定読み込み
    cfg := config.Load()

    // 2. DB 初期化（接続＆マイグレーション）
    database.Init(cfg)

    // 3. Gin エンジン起動
    r := gin.Default()

    // セッションミドルウェア等の設定
    // ...

    // ルート定義
    // r.POST("/signup", handlers.SignUp)
    // r.POST("/login", handlers.Login)
    // api := r.Group("/")
    // api.Use(middleware.RequireLogin())
    // api.GET("/trades", handlers.ListTrades)
    // ...

    addr := fmt.Sprintf(":%s", cfg.AppPort)
    log.Printf("Starting server on %s…", addr)
    if err := r.Run(addr); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}

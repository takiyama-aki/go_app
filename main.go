// main.go
package main

import (
    "fmt"
    "log"

    "github.com/gin-gonic/gin"
    "github.com/takiyama-aki/go_app/config"
    "github.com/takiyama-aki/go_app/database"
    "github.com/takiyama-aki/go_app/handlers"
    "github.com/takiyama-aki/go_app/middleware"
)

func main() {
    // 1. 設定読み込み
    cfg := config.Load()

    // 2. DB 初期化（接続＆マイグレーション）
    database.Init(cfg)

    // 3. Gin エンジン起動
    r := gin.Default()

    // セッション用 Cookie ストア登録
    store := cookie.NewStore([]byte(cfg.SessionKey))
    r.Use(sessions.Sessions("goapp_session", store))

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

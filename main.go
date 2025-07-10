// main.go
package main

import (
    "fmt"
    "log"

    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
    "github.com/gin-contrib/cors"

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

    // CORS ─ フロント (Vite) からのリクエストを許可
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Content-Type"},
        AllowCredentials: true, // Cookie 認証なら必須
    }))

    // セッション用 Cookie ストア登録
    store := cookie.NewStore([]byte(cfg.SessionKey))

    // セッションのオプション設定
    store.Options(sessions.Options{
    Path:     "/",
    MaxAge:   60 * 60 * 24, // 1 日
    HttpOnly: true,
})
    r.Use(sessions.Sessions("goapp_session", store))

    // ルート定義
    // 認証不要ルート
    r.POST("/signup", handlers.SignUp)
    r.POST("/login",  handlers.Login)

    // 認証必須ルート
    auth := r.Group("/")
    auth.Use(middleware.RequireLogin())

    // Trade CRUD
    auth.GET(   "/trades",        handlers.ListTrades)
    auth.POST(  "/trades",        handlers.CreateTrade)
    auth.PUT(   "/trades/:id",    handlers.UpdateTrade)
    auth.DELETE("/trades/:id",    handlers.DeleteTrade)

    // Account 残高
    auth.GET(  "/account/balance", handlers.GetBalance)
    auth.PUT(  "/account/balance", handlers.UpdateBalance)

    // サーバー起動
    addr := fmt.Sprintf(":%s", cfg.AppPort)
    log.Printf("Starting server on %s…", addr)
    if err := r.Run(addr); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}

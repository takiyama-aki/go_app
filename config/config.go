package config

import (
    "log"
    "os"
    "strconv"

    "github.com/joho/godotenv"
)

// Config holds application-wide settings.
type Config struct {
    AppPort    string // アプリケーション起動ポート
    DBHost     string
    DBPort     int
    DBUser     string
    DBPassword string
    DBName     string
    SessionKey string // セッション用シークレット
}

// Load reads .env (if exists) and environment変数を Config に詰める
func Load() *Config {
    // .envがあれば読み込む（本番では .env を置かない運用にすることも多い）
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, relying on environment variables")
    }

    port, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
    if err != nil {
        log.Fatalf("Invalid DB_PORT: %v", err)
    }

    return &Config{
        AppPort:    getEnv("APP_PORT", "8080"),
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     port,
        DBUser:     getEnv("DB_USER", ""),
        DBPassword: getEnv("DB_PASSWORD", ""),
        DBName:     getEnv("DB_NAME", ""),
        SessionKey: getEnv("SESSION_KEY", "changeme"),
    }
}

// getEnv は環境変数を返し、未設定ならデフォルトを返す
func getEnv(key, defaultVal string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return defaultVal
}

package config

// アプリ全体の設定値を読み込むためのパッケージ。
// .env から取得した値を構造体にまとめて提供する。

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

// Load は .env と環境変数を読み取り Config 構造体を返す
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
		AppPort:    getEnv("APP_PORT", "8081"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     port,
		DBUser:     getEnv("DB_USER", ""),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", ""),
		SessionKey: getEnv("SESSION_KEY", "changeme"),
	}
}

// getEnv は環境変数を返し、未設定ならデフォルトを返す
// getEnv は環境変数を取得し、存在しなければデフォルト値を返す
func getEnv(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}

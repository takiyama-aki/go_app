package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"

	"github.com/takiyama-aki/go_app/config"
	"github.com/takiyama-aki/go_app/database"
	"github.com/takiyama-aki/go_app/models"
)

var oauthCfg *oauth2.Config

// InitOAuth は設定から OAuth2 の設定を初期化する
func InitOAuth(cfg *config.Config) {
	oauthCfg = &oauth2.Config{
		ClientID:     cfg.OAuthClientID,
		ClientSecret: cfg.OAuthClientSecret,
		RedirectURL:  cfg.OAuthRedirectURL,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

// OAuthLogin はプロバイダへリダイレクトする
func OAuthLogin(c *gin.Context) {
	if oauthCfg == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "oauth not configured"})
		return
	}
	url := oauthCfg.AuthCodeURL("state", oauth2.AccessTypeOnline)
	c.Redirect(http.StatusFound, url)
}

// googleUser は Google から取得するユーザー情報
type googleUser struct {
	Email string `json:"email"`
}

// OAuthCallback は認可コードからユーザー情報を取得しログインさせる
func OAuthCallback(c *gin.Context) {
	if oauthCfg == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "oauth not configured"})
		return
	}
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "code not provided"})
		return
	}

	ctx := context.Background()
	tok, err := oauthCfg.Exchange(ctx, code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token exchange failed"})
		return
	}

	client := oauthCfg.Client(ctx, tok)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to fetch user info"})
		return
	}
	defer resp.Body.Close()

	var guser googleUser
	if err := json.NewDecoder(resp.Body).Decode(&guser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user info"})
		return
	}

	var user models.User
	res := database.DB.Where("email = ?", guser.Email).First(&user)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			hash, _ := bcrypt.GenerateFromPassword([]byte("oauth"), bcrypt.DefaultCost)
			user = models.User{Email: guser.Email, PasswordHash: string(hash), CreatedAt: time.Now()}
			if err := database.DB.Create(&user).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch user"})
			return
		}
	}

	sess := sessions.Default(c)
	sess.Set("user_id", user.ID)
	if err := sess.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save session"})
		return
	}

	// ログイン後はフロントエンドへリダイレクト
	c.Redirect(http.StatusFound, "http://localhost:5173")
}

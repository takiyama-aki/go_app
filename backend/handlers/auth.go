// handlers/auth.go
package handlers

// 認証関連のハンドラをまとめたファイル。
// ユーザー登録・ログイン・ログアウトおよび自身の情報取得を扱う。

import (
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/takiyama-aki/go_app/database"
	"github.com/takiyama-aki/go_app/helpers"
	"github.com/takiyama-aki/go_app/models"
)

type SignUpRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// SignUp は新規ユーザーを作成しセッションに user_id を保存する
func SignUp(c *gin.Context) {
	var req SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user := models.User{Email: req.Email, PasswordHash: string(hash), CreatedAt: time.Now()}
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "email already in use"})
		return
	}

	sess := sessions.Default(c)
	sess.Set("user_id", user.ID)
	if err := sess.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save session"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": user.ID, "email": user.Email})
}

// Login は既存ユーザーの認証を行いセッションに user_id を設定する
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	sess := sessions.Default(c)
	sess.Set("user_id", user.ID)
	if err := sess.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "login successful"})
}

// Logout はセッション情報を破棄してログアウトさせる
func Logout(c *gin.Context) {
	sess := sessions.Default(c)
	sess.Clear()
	if err := sess.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to clear session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}

// GetMe は現在ログイン中のユーザー情報を返す
func GetMe(c *gin.Context) {
	uid, ok := helpers.CurrentUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var user models.User
	if err := database.DB.First(&user, uid).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": user.ID, "email": user.Email})
}

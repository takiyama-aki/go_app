// handlers/auth.go
package handlers

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
    "golang.org/x/crypto/bcrypt"

    "github.com/takiyama-aki/go_app/database"
    "github.com/takiyama-aki/go_app/models"
)

// SignUp : 新規ユーザー登録
func SignUp(c *gin.Context) {
    var req SignUpRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // パスワードハッシュ化
    hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "could not hash password"})
        return
    }

    user := models.User{
        Email:        req.Email,
        PasswordHash: string(hash),
        CreatedAt:    time.Now(),
    }
    if err := database.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusConflict, gin.H{"error": "email already in use"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"id": user.ID, "email": user.Email})
}

// Login : 認証 → セッションに user_id をセット
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

    // パスワードチェック
    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }

    // セッションに保存
    sess := sessions.Default(c)
    sess.Set("user_id", user.ID)
    if err := sess.Save(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save session"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "login successful"})
}

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

type SignUpRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

// SignUp はユーザー登録のエンドポイント（スタブ）
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
    c.JSON(http.StatusCreated, gin.H{"id": user.ID, "email": user.Email})
}

// Login はログインのエンドポイント（スタブ）
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
    sess.Save()
    c.JSON(http.StatusOK, gin.H{"message": "login successful"})
}

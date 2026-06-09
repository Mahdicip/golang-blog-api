package handler

import (
	"net/http"

	"blog-api/internal/database"
	model "blog-api/internal/models"
	"blog-api/pkg/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)
type RegisterInput struct {
    Name     string `json:"name" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

type LoginInput struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword(
		[]byte(input.Password),
		bcrypt.DefaultCost,
	)

	user := model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     "user",
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Email already exists",
		})
		return
	}

	token, _ := utils.GenerateToken(user.ID, user.Role)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"token":   token,
		"user":    user,
	})
}

func Login(c *gin.Context) {
    var input LoginInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user model.User
    if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token, _ := utils.GenerateToken(user.ID, user.Role)
    c.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}

func GetMe(c *gin.Context) {
    userID, _ := c.Get("userID")
    var user model.User
    database.DB.First(&user, userID)
    c.JSON(http.StatusOK, user)
}

func UpdateProfile(c *gin.Context) {
    userID, _ := c.Get("userID")
    var input struct {
        Name string `json:"name"`
    }
    c.ShouldBindJSON(&input)
    database.DB.Model(&model.User{}).Where("id = ?", userID).Update("name", input.Name)
    c.JSON(http.StatusOK, gin.H{"message": "Profile updated"})
}
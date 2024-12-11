package controllers

import (
	"net/http"
	"server/configs"
	"server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *gin.Context) {
	var users []models.User;

	if err := configs.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users);
}

func CreateUser(c *gin.Context) {
	var body models.CreateNewUser;

	// Validate body input
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()});

		return;
	}

	// Check Email exist
	var existingUser models.User;

	if err := configs.DB.Where("email = ?", body.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"});

		return;
	}

	// Password hased
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost);

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"});

		return;
	}

	// Create new user
	user := models.User{
		ID:       uuid.New(),
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hashedPassword),
	}

	// Save database
	if err := configs.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()});

		return;
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        user.ID,
		"name":      user.Name,
		"email":     user.Email,
		"createdAt": user.CreatedAt,
	});
}
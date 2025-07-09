package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/rizkyimaduddin24/techtest/internal/entity"
	"gitlab.com/rizkyimaduddin24/techtest/internal/usecase"
)

type AuthHandler struct {
	uc usecase.AuthUsecase
}

func AuthRoutes(r *gin.Engine, uc usecase.AuthUsecase) {
	h := &AuthHandler{uc}

	// Auth Group
	authGroup := r.Group("/auth")
	authGroup.POST("/register", h.Register)
	authGroup.POST("/login", h.Login)
}

func (h *AuthHandler) Register(c *gin.Context) {
	// Get Body
	var body entity.AuthRegisterBody
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// Usecase Register
	if err := h.uc.Register(body.Name, body.Email, body.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User Registered"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	// Get Body
	var body entity.AuthLoginBody
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// Usecase Login
	id, token, err := h.uc.Login(body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"token": token,
	})
}

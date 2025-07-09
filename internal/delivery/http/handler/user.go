package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	mw "gitlab.com/rizkyimaduddin24/techtest/internal/delivery/http/middleware"
	"gitlab.com/rizkyimaduddin24/techtest/internal/entity"
	"gitlab.com/rizkyimaduddin24/techtest/internal/usecase"
)

type UserHandler struct {
	uc usecase.UserUsecase
}

func UserRoutes(r *gin.Engine, uc usecase.UserUsecase) {
	h := &UserHandler{uc}

	// User Group
	userGroup := r.Group("/user")
	userGroup.GET("", mw.RequireAuth(""), h.GetAll)
	userGroup.GET("/:id", mw.RequireAuth(""), h.GetByID)
	userGroup.PUT("/:id", mw.RequireAuth(""), h.Update)
	userGroup.DELETE("/:id", mw.RequireAuth(""), h.Delete)
}

func (h *UserHandler) GetAll(c *gin.Context) {
	// Usecase Get All
	users, err := h.uc.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetByID(c *gin.Context) {
	// Get Params
	id := c.Param("id")

	// Usecase Login
	user, err := h.uc.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Update(c *gin.Context) {
	// Get Params
	id := c.Param("id")

	// Validate Role & ID
	tokenRole := c.GetString("x-user-role")
	if tokenRole != "admin" {
		tokenId := c.GetString("x-user-id")
		if id != tokenId {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}

	// Get Body
	var body entity.UserUpdateBody
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// Usecase Update
	if err := h.uc.Update(id, body.Name, body.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated"})
}

func (h *UserHandler) Delete(c *gin.Context) {
	// Get Params
	id := c.Param("id")

	// Validate Role & ID
	tokenRole := c.GetString("x-user-role")
	if tokenRole != "admin" {
		tokenId := c.GetString("x-user-id")
		if id != tokenId {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}

	// Usecase Delete
	if err := h.uc.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

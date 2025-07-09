package handler

import (
	"io"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	mw "gitlab.com/rizkyimaduddin24/techtest/internal/delivery/http/middleware"
	"gitlab.com/rizkyimaduddin24/techtest/internal/entity"
	"gitlab.com/rizkyimaduddin24/techtest/internal/usecase"
	"gitlab.com/rizkyimaduddin24/techtest/pkg"
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
	tokenId := c.GetString("x-user-id")
	tokenRole := c.GetString("x-user-role")
	if tokenRole != "admin" {
		if id != tokenId {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}

	// Get Body
	var body entity.UserUpdateBody
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// Get File
	var fileUrl string
	fileHeader, _ := c.FormFile("document")
	if fileHeader != nil {
		// Open File
		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "cannot open file"})
			return
		}
		defer file.Close()

		// Read File
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read file"})
			return
		}

		// Upload to Blackblaze
		safeName := filepath.Base(fileHeader.Filename)
		fileUrl, err = pkg.UploadToB2(fileBytes, safeName)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	// Usecase Update
	if err := h.uc.Update(id, body.Name, body.Email, fileUrl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated"})
}

func (h *UserHandler) Delete(c *gin.Context) {
	// Get Params
	id := c.Param("id")

	// Validate Role & ID
	tokenId := c.GetString("x-user-id")
	tokenRole := c.GetString("x-user-role")
	if tokenRole != "admin" {
		if id != tokenId {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}

	// Usecase Delete
	if err := h.uc.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

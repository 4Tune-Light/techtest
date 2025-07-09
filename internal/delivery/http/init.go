package deliverhttp

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/rizkyimaduddin24/techtest/internal/delivery/http/handler"
	"gitlab.com/rizkyimaduddin24/techtest/internal/repository"
	"gitlab.com/rizkyimaduddin24/techtest/internal/usecase"
	"gorm.io/gorm"
)

func Serve(db *gorm.DB) {
	r := gin.Default()

	// Load Repository
	userRepo := repository.NewUserRepository(db)

	// Load Usecase
	authUc := usecase.NewAuthUsecase(userRepo)
	userUc := usecase.NewUserUsecase(userRepo)

	// Load Routes
	handler.AuthRoutes(r, authUc)
	handler.UserRoutes(r, userUc)

	r.Run()
}

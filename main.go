package main

import (
	"auth/src/controller"
	"auth/src/repositories"
	"auth/src/server/config"
	"auth/src/server/middleware"
	"auth/src/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db           *gorm.DB                    = config.SetupDatabaseConnection()
	repo         repositories.UserRepository = repositories.NewUserRepository(db)
	tokenService service.Token               = service.NewTokenUc()
	authService  service.Auth                = service.NewAuthUC(repo)
	ac           controller.AuthController   = controller.NewAuthController(authService, tokenService)
)

func main() {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	authRoutes := r.Group("auth")
	{
		authRoutes.POST("/register", ac.Register)
		authRoutes.POST("/login", ac.Login)
		authRoutes.POST("/logout", ac.Logout, middleware.AuthorizeJWT(tokenService))
	}

	r.Run()
}

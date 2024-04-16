package auth

import (
	"github.com/CValier/PruebaGO/internal/infra/api/repositories/db"
	"github.com/CValier/PruebaGO/internal/pkg/service/auth"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	authRoutes := e.Group("/api/v1/auth")

	repo := db.NewAuthRepository()
	authSvc := auth.NewAuthService(repo)
	authHandler := newHandler(authSvc)

	authRoutes.POST("/login", authHandler.signInUser)
	authRoutes.POST("/signup", authHandler.registerUser)
}

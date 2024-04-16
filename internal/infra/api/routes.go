package api

import (
	"github.com/CValier/PruebaGO/internal/infra/api/auth"
	"github.com/gin-gonic/gin"
)

func registerRoutes(e *gin.Engine) {
	auth.RegisterRoutes(e)
}

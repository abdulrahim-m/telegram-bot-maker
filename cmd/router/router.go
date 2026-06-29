package router

import (
	"time"

	"github.com/abdulrahim-m/telegram-bot-maker/cmd/consts"
	"github.com/abdulrahim-m/telegram-bot-maker/internal/handlers/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() {
	r := gin.New()
	{ // ==== Config ====
		r.Use(gin.CustomRecovery(func(c *gin.Context, err any) {
			c.JSON(500, consts.InternalServerError)
			c.Abort()
		}), gin.Logger())
		r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: false,
			MaxAge:           12 * time.Hour,
		}))
		r.Use(middleware.ErrorHandlerMiddleware())
		r.GET("/metrics", gin.WrapH(promhttp.Handler()))
		r.GET("/test", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"status": 200}) })
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	// apiV1 := r.Group("/api/v1")
}

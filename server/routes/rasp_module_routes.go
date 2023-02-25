package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"server/controller"
	"server/middleware"
)

func InitRaspModuleRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	raspModuleController := controller.NewRaspModuleController()
	router := r.Group("/module")
	router.Use(authMiddleware.MiddlewareFunc())
	router.Use(middleware.CasbinMiddleware())
	{
		router.POST("/create", raspModuleController.CreateRaspModule)
		router.GET("/list", raspModuleController.GetRaspModules)
		router.DELETE("/delete/batch", raspModuleController.BatchDeleteModuleByIds)
	}
	return r
}

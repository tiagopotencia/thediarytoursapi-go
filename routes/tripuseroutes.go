package routes

import (
	"mol/business"

	"github.com/gin-gonic/gin"
)

func SetTripUserRoutes(r *gin.RouterGroup) {

	resource := r.Group("tripuser")

	resource.GET("/", business.GetAllTripUsers)
	resource.GET("/:id", business.GetTripUser)
	resource.POST("/", business.PostTripUser)
	resource.PUT("/:id", business.PutTripUser)
	resource.DELETE("/:id", business.DeleteTripUser)

}

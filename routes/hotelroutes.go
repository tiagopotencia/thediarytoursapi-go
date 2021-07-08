package routes

import (
	"mol/business"

	"github.com/gin-gonic/gin"
)

func SetHotelRoutes(r *gin.RouterGroup) {

	resource := r.Group("hotels")

	resource.GET("/", business.GetAllHotels)
	resource.GET("/:id", business.GetHotel)
	resource.POST("/", business.PostHotel)
	resource.PUT("/:id", business.PutHotel)
	resource.DELETE("/:id", business.DeleteHotel)

}

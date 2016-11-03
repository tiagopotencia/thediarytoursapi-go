package routes

import (
	"github.com/gin-gonic/gin"
	"git.heroku.com/thediarytoursapi-go/business"
)

func SetHotelRoutes(r *gin.RouterGroup) {

	resource := r.Group("hotels")

	resource.GET("/", business.GetAllHotels)
	resource.GET("/:id", business.GetHotel)
	resource.POST("/", business.PostHotel)
	resource.PUT("/:id", business.PutHotel)
	resource.DELETE("/:id", business.DeleteHotel)

}
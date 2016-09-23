package routes

import (
	"github.com/gin-gonic/gin"
	"git.heroku.com/thediarytoursapi-go/business"
)

func SetItineraryRoutes(r *gin.RouterGroup) {

	resource := r.Group("itinerary")

	resource.GET("/", business.GetAllItinerarys)
	resource.GET("/:id", business.GetItinerary)
	resource.POST("/", business.PostItinerary)
	resource.PUT("/:id", business.PutItinerary)
	resource.DELETE("/:id", business.DeleteItinerary)

}
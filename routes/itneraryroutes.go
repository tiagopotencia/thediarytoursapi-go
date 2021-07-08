package routes

import (
	"mol/business"

	"github.com/gin-gonic/gin"
)

func SetItineraryRoutes(r *gin.RouterGroup) {

	resource := r.Group("itinerary")

	resource.GET("/", business.GetAllItinerarys)
	//resource.GET("/:id", business.GetItinerary)
	resource.GET("/:dia", business.GetItineraryByDay)
	resource.POST("/", business.PostItinerary)
	resource.PUT("/:id", business.PutItinerary)
	resource.DELETE("/:id", business.DeleteItinerary)

}

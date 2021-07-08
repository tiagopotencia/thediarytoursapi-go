package routes

import (
	"mol/business"

	"github.com/gin-gonic/gin"
)

func SetTripRoutes(r *gin.RouterGroup) {

	resource := r.Group("trips")

	resource.GET("/", business.GetAllTrips)
	resource.GET("/:id", business.GetTrip)
	resource.POST("/", business.PostTrip)
	resource.PUT("/:id", business.PutTrip)
	resource.DELETE("/:id", business.DeleteTrip)

}

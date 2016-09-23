package routes

import (
	"github.com/gin-gonic/gin"
	"git.heroku.com/thediarytoursapi-go/business"
)

func SetTripRoutes(r *gin.RouterGroup)  {

	resource := r.Group("trips")

	resource.GET("/", business.GetAllTrips)
	resource.GET("/:id", business.GetTrip)
	resource.POST("/", business.PostTrip)
	resource.PUT("/:id", business.PutTrip)
	resource.DELETE("/:id", business.DeleteTrip)

}
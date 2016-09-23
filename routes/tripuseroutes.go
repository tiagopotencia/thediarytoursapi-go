package routes

import (
	"github.com/gin-gonic/gin"
	"git.heroku.com/thediarytoursapi-go/business"
)

func SetTripUserRoutes(r *gin.RouterGroup) {

	resource := r.Group("tripuser")

	resource.GET("/", business.GetAllTripUsers)
	resource.GET("/:id", business.GetTripUser)
	resource.POST("/", business.PostTripUser)
	resource.PUT("/:id", business.PutTripUser)
	resource.DELETE("/:id", business.DeleteTripUser)

}
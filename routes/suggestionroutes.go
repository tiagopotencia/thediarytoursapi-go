package routes

import (
	"github.com/gin-gonic/gin"
	"git.heroku.com/thediarytoursapi-go/business"
)

func SetSuggestionRoutes(r *gin.RouterGroup)  {

	resource := r.Group("suggestions")

	resource.GET("/", business.GetAllSuggestions)
	resource.GET("/:id", business.GetSuggestion)
	resource.POST("/", business.PostSuggestion)
	resource.PUT("/:id", business.PutSuggestion)
	resource.DELETE("/:id", business.DeleteSuggestion)

}
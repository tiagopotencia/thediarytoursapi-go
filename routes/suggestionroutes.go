package routes

import (
	"mol/business"

	"github.com/gin-gonic/gin"
)

func SetSuggestionRoutes(r *gin.RouterGroup) {

	resource := r.Group("suggestions")

	resource.GET("/", business.GetAllSuggestions)
	resource.GET("/:id", business.GetSuggestion)
	resource.POST("/", business.PostSuggestion)
	resource.PUT("/:id", business.PutSuggestion)
	resource.DELETE("/:id", business.DeleteSuggestion)

}

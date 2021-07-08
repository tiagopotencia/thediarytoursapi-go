package routes

import (
	"mol/business"

	"github.com/gin-gonic/gin"
)

func SetMessageRoutes(r *gin.RouterGroup) {

	resource := r.Group("message")

	resource.GET("/", business.GetAllMessages)
	resource.GET("/:id", business.GetMessage)
	resource.POST("/", business.PostMessage)
	resource.PUT("/:id", business.PutMessage)
	resource.DELETE("/:id", business.DeleteMessage)

}

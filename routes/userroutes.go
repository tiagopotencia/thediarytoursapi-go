package routes

import (
	"mol/business"

	"github.com/gin-gonic/gin"
)

func SetUserRoutes(r *gin.RouterGroup) {

	resource := r.Group("user")

	resource.GET("/", business.GetAllUsers)
	resource.GET("/:id", business.GetUser)
	resource.POST("/", business.PostUser)
	resource.PUT("/:id", business.PutUser)
	resource.DELETE("/:id", business.DeleteUser)

}

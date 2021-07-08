package routes

import (
	"mol/business"

	"github.com/gin-gonic/gin"
)

func SetPhoneInBrazilRoutes(r *gin.RouterGroup) {

	resource := r.Group("phoneinbrazil")

	resource.GET("/", business.GetAllPhoneInBrazil)
	resource.GET("/:id", business.GetPhoneInBrazil)
	resource.POST("/", business.PostPhoneInBrazil)
	resource.PUT("/:id", business.PutPhoneInBrazil)
	resource.DELETE("/:id", business.DeletePhoneInBrazil)

}

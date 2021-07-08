package routes

import (
	"mol/business"

	"github.com/gin-gonic/gin"
)

func SetTripUtilNumberRoutes(r *gin.RouterGroup) {

	resource := r.Group("triputilnumber")

	resource.GET("/", business.GetAllTripUtilNumbers)
	resource.GET("/:id", business.GetTripUtilNumber)
	resource.POST("/", business.PostTripUtilNumber)
	resource.PUT("/:id", business.PutTripUtilNumber)
	resource.DELETE("/:id", business.DeleteTripUtilNumber)

}

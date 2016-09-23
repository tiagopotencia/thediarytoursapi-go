package routes

import "github.com/gin-gonic/gin"

//SetRoutes set all endpoint's routes
func SetRoutes(r *gin.Engine){

	v1 := r.Group("v1")

	SetTripRoutes(v1)

}
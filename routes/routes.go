package routes

import "github.com/gin-gonic/gin"

//SetRoutes set all endpoint's routes
func SetRoutes(r *gin.Engine) {

	v1 := r.Group("v1")

	SetTripRoutes(v1)
	SetSuggestionRoutes(v1)
	SetPhoneInBrazilRoutes(v1)
	SetUserRoutes(v1)
	SetTripUserRoutes(v1)
	SetTripUtilNumberRoutes(v1)
	SetItineraryRoutes(v1)
	SetMessageRoutes(v1)

}

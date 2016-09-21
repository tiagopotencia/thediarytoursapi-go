package business

import (
	"git.heroku.com/thediarytoursapi-go/connection"
	"git.heroku.com/thediarytoursapi-go/bindata"
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"
	"strconv"
	"git.heroku.com/thediarytoursapi-go/apireturns"
)

type Trip struct {

ID int64 `db:"id" json:"id"`
Code string `db:"code" json:"code"`
About string `db:"about" json:"about"`
Name string `db:"name" json:"name"`

}

var jr apireturns.JSONResult

//GetAllTrips get all Trips from db
func GetAllTrips(c *gin.Context)  {
	db := connection.GetConnection()
	defer db.Close()
	query, err := bindata.Asset("queries\\trip\\getAllTrips.sql")

	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("Query file not found").Error())
		return
	}

	var trips []Trip

	db.Select(&trips, string(query))

	c.JSON(http.StatusOK, trips)
}

//GetTrip get a Trip from db based on ID
func GetTrip(c *gin.Context)  {
	db := connection.GetConnection()
	defer db.Close()
	query, err := bindata.Asset("queries\\trip\\getTrip.sql")

	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("Query file not found").Error())
		return
	}

	var trip Trip

	ID, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	db.Get(&trip, string(query), ID)

	c.JSON(http.StatusOK, trip)
}

//PostTrip insert a new Trip
func PostTrip(c *gin.Context)  {
	db := connection.GetConnection()
	defer db.Close()
	query, err := bindata.Asset("queries\\trip\\insertTrip.sql")
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("Query file not found").Error())
		return
	}

	var trip Trip

	err = c.Bind(&trip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	tx, err := db.Begin()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	tripStmt, err := tx.Prepare(string(query))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	result, err := tripStmt.Exec(trip.Code, trip.About, trip.Name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	rows, err := result.RowsAffected()

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		tx.Rollback()

	}

	if rows > 0{
		jr.Code = http.StatusOK
		jr.Message = "Trip created successfully"
		tx.Commit()
	} else {
		jr.Code = http.StatusInternalServerError
		jr.Message = "Some error ocourred during insert on db"
		tx.Rollback()
	}

	c.JSON(jr.Code, jr)


}
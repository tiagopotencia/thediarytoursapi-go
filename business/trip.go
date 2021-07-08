package business

import (
	"errors"
	"log"
	"mol/apireturns"
	"mol/connection"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/doug-martin/goqu.v3"
	_ "gopkg.in/doug-martin/goqu.v3/adapters/postgres"
)

type Trip struct {
	ID    int64  `db:"id" json:"id"`
	Code  string `db:"code" json:"code"`
	About string `db:"about" json:"about"`
	Name  string `db:"name" json:"name"`
}

var jr apireturns.JSONResult

//GetAllTrips get all Trips from db
func GetAllTrips(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()

	db := goqu.New("postgres", conn.DB)

	//query, err := bindata.Asset("queries\\trip\\getAllTrips.sql")

	query, _, err := db.From("trip").ToSql()

	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("Query file not found").Error())
		return
	}

	var trips []Trip

	log.Print(query)

	conn.Select(&trips, string(query))

	c.JSON(http.StatusOK, trips)
}

//GetTrip get a Trip from db based on ID
func GetTrip(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()

	db := goqu.New("postgres", conn.DB)

	query, _, err := db.From("trip").Where(goqu.Ex{"id": c.Param("id")}).ToSql()

	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("Query file not found").Error())
		return
	}

	var trip Trip

	conn.Get(&trip, string(query))

	c.JSON(http.StatusOK, trip)
}

//PostTrip insert a new Trip
func PostTrip(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()
	//query, err := bindata.Asset("queries\\trip\\insertTrip.sql")
	//INSERT INTO trip(code, about, name) VALUES ($1, $2, $3);

	db := goqu.New("postgres", conn.DB)

	var trip Trip

	err := c.Bind(&trip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	query := db.From("trip").Insert(goqu.Record{"code": trip.Code, "about": trip.About, "name": trip.Name})

	result, err := query.Exec()

	if err != nil {
		log.Print(err.Error())
	}

	rows, err := result.RowsAffected()

	if rows > 0 {
		jr.Code = http.StatusOK
		jr.Message = "Trip created successfully"
	} else {
		jr.Code = http.StatusInternalServerError
		jr.Message = "Some error ocourred during insert on db"
	}

	c.JSON(jr.Code, jr)

}

func PutTrip(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()
	//query, err := bindata.Asset("queries\\trip\\insertTrip.sql")
	//INSERT INTO trip(code, about, name) VALUES ($1, $2, $3);

	db := goqu.New("postgres", conn.DB)

	var trip Trip

	err := c.Bind(&trip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	query := db.From("trip").Where(goqu.Ex{"id": c.Param("id")})
	exec := query.Update(goqu.Record{"code": trip.Code, "about": trip.About, "name": trip.Name})

	result, err := exec.Exec()

	if err != nil {
		log.Print(err.Error())
	}

	rows, err := result.RowsAffected()

	if rows > 0 {
		jr.Code = http.StatusOK
		jr.Message = "Trip edited successfully"
	} else {
		jr.Code = http.StatusInternalServerError
		jr.Message = "Some error ocourred during insert on db"
	}

	c.JSON(jr.Code, jr)

}

func DeleteTrip(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()
	//query, err := bindata.Asset("queries\\trip\\insertTrip.sql")
	//INSERT INTO trip(code, about, name) VALUES ($1, $2, $3);

	db := goqu.New("postgres", conn.DB)

	var trip Trip

	err := c.Bind(&trip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	query := db.From("trip").Where(goqu.Ex{"id": c.Param("id")})
	exec := query.Delete()

	result, err := exec.Exec()

	if err != nil {
		log.Print(err.Error())
	}

	rows, err := result.RowsAffected()

	if rows > 0 {
		jr.Code = http.StatusOK
		jr.Message = "Trip deleted successfully"
	} else {
		jr.Code = http.StatusInternalServerError
		jr.Message = "Some error ocourred during insert on db"
	}

	c.JSON(jr.Code, jr)

}

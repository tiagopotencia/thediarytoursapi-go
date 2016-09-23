package business

import (
	"git.heroku.com/thediarytoursapi-go/connection"
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"
	"log"
	"gopkg.in/doug-martin/goqu.v3"
	_ "gopkg.in/doug-martin/goqu.v3/adapters/postgres"
)

//TODO: Renomear nome das coluas do banco
type TripUser struct {
	ID    int64 `db:"id" json:"id"`
	IdTrip int64 `db:"id_trip"`
	IdUser int64 `db:"id_user"`
}

func GetAllTripUsers(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()

	db := goqu.New("postgres", conn.DB)

	//query, err := bindata.Asset("queries\\trip\\getAllTripUsers.sql")

	query, _, err := db.From("trip_user").ToSql()

	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("Query file not found").Error())
		return
	}

	var tripUser []TripUser

	log.Print(query)

	conn.Select(&tripUser, string(query))

	c.JSON(http.StatusOK, tripUser)
}

//GetTripUser get a TripUser from db based on ID
func GetTripUser(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()

	db := goqu.New("postgres", conn.DB)

	query, _, err := db.From("trip_user").Where(goqu.Ex{"id":c.Param("id")}).ToSql()

	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("Query file not found").Error())
		return
	}

	var trip TripUser

	conn.Get(&trip, string(query))

	c.JSON(http.StatusOK, trip)
}

//PostTripUser insert a new TripUser
func PostTripUser(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()
	//query, err := bindata.Asset("queries\\trip\\insertTripUser.sql")
	//INSERT INTO trip(code, about, name) VALUES ($1, $2, $3);

	db := goqu.New("postgres", conn.DB)

	var trip TripUser

	err := c.Bind(&trip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	query := db.From("trip_user").Insert(goqu.Record{"id_trip": trip.IdTrip, "id_user":trip.IdUser})

	result, err := query.Exec()

	if err != nil {
		log.Print(err.Error())
	}

	rows, err := result.RowsAffected()

	if rows > 0 {
		jr.Code = http.StatusOK
		jr.Message = "TripUser created successfully"
	} else {
		jr.Code = http.StatusInternalServerError
		jr.Message = "Some error ocourred during insert on db"
	}

	c.JSON(jr.Code, jr)

}

func PutTripUser(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()
	//query, err := bindata.Asset("queries\\trip\\insertTripUser.sql")
	//INSERT INTO trip(code, about, name) VALUES ($1, $2, $3);

	db := goqu.New("postgres", conn.DB)

	var trip TripUser

	err := c.Bind(&trip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	query := db.From("trip_user").Where(goqu.Ex{"id":c.Param("id")})
	exec := query.Update(goqu.Record{"id_trip": trip.IdTrip, "id_user":trip.IdUser})

	result, err := exec.Exec()

	if err != nil {
		log.Print(err.Error())
	}

	rows, err := result.RowsAffected()

	if rows > 0 {
		jr.Code = http.StatusOK
		jr.Message = "TripUser edited successfully"
	} else {
		jr.Code = http.StatusInternalServerError
		jr.Message = "Some error ocourred during insert on db"
	}

	c.JSON(jr.Code, jr)

}

func DeleteTripUser(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()
	//query, err := bindata.Asset("queries\\trip\\insertTripUser.sql")
	//INSERT INTO trip(code, about, name) VALUES ($1, $2, $3);

	db := goqu.New("postgres", conn.DB)

	var trip TripUser

	err := c.Bind(&trip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	query := db.From("trip_user").Where(goqu.Ex{"id":c.Param("id")})
	exec := query.Delete()

	result, err := exec.Exec()

	if err != nil {
		log.Print(err.Error())
	}

	rows, err := result.RowsAffected()

	if rows > 0 {
		jr.Code = http.StatusOK
		jr.Message = "TripUser deleted successfully"
	} else {
		jr.Code = http.StatusInternalServerError
		jr.Message = "Some error ocourred during insert on db"
	}

	c.JSON(jr.Code, jr)

}
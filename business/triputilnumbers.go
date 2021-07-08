package business

import (
	"errors"
	"log"
	"mol/connection"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/doug-martin/goqu.v3"
	_ "gopkg.in/doug-martin/goqu.v3/adapters/postgres"
)

//TODO: Renomear nome das coluas do banco
type TripUtilNumbers struct {
	ID          int64  `db:"id"`
	Type        string `db:"type"`
	Phone       string `db:"phone"`
	Description string `db:"description"`
	IdTrip      int64  `db:"id_trip"`
}

func GetAllTripUtilNumbers(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()

	db := goqu.New("postgres", conn.DB)

	query, _, err := db.From("trip_util_numbers").ToSql()

	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("Query file not found").Error())
		return
	}

	var tripUtilNumbers []TripUtilNumbers

	log.Print(query)

	conn.Select(&tripUtilNumbers, string(query))

	c.JSON(http.StatusOK, tripUtilNumbers)
}

//GetTripUtilNumber get a TripUtilNumber from db based on ID
func GetTripUtilNumber(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()

	db := goqu.New("postgres", conn.DB)

	query, _, err := db.From("trip_util_numbers").Where(goqu.Ex{"id": c.Param("id")}).ToSql()

	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("Query file not found").Error())
		return
	}

	var tripUtilNumber TripUtilNumbers

	conn.Get(&tripUtilNumber, string(query))

	c.JSON(http.StatusOK, tripUtilNumber)
}

//PostTripUtilNumber insert a new TripUtilNumber
func PostTripUtilNumber(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()
	//query, err := bindata.Asset("queries\\tripUtilNumber\\insertTripUtilNumber.sql")
	//INSERT INTO tripUtilNumber(code, about, name) VALUES ($1, $2, $3);

	db := goqu.New("postgres", conn.DB)

	var tripUtilNumber TripUtilNumbers

	err := c.Bind(&tripUtilNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	query := db.From("trip_util_numbers").Insert(goqu.Record{"type": tripUtilNumber.Type, "phone": tripUtilNumber.Phone, "description": tripUtilNumber.Description, "id_trip": tripUtilNumber.IdTrip})

	result, err := query.Exec()

	if err != nil {
		log.Print(err.Error())
	}

	rows, err := result.RowsAffected()

	if rows > 0 {
		jr.Code = http.StatusOK
		jr.Message = "TripUtilNumber created successfully"
	} else {
		jr.Code = http.StatusInternalServerError
		jr.Message = "Some error ocourred during insert on db"
	}

	c.JSON(jr.Code, jr)

}

func PutTripUtilNumber(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()
	//query, err := bindata.Asset("queries\\tripUtilNumber\\insertTripUtilNumber.sql")
	//INSERT INTO tripUtilNumber(code, about, name) VALUES ($1, $2, $3);

	db := goqu.New("postgres", conn.DB)

	var tripUtilNumber TripUtilNumbers

	err := c.Bind(&tripUtilNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	query := db.From("trip_util_numbers").Where(goqu.Ex{"id": c.Param("id")})
	exec := query.Update(goqu.Record{"type": tripUtilNumber.Type, "phone": tripUtilNumber.Phone, "description": tripUtilNumber.Description, "id_trip": tripUtilNumber.IdTrip})

	result, err := exec.Exec()

	if err != nil {
		log.Print(err.Error())
	}

	rows, err := result.RowsAffected()

	if rows > 0 {
		jr.Code = http.StatusOK
		jr.Message = "TripUtilNumber edited successfully"
	} else {
		jr.Code = http.StatusInternalServerError
		jr.Message = "Some error ocourred during insert on db"
	}

	c.JSON(jr.Code, jr)

}

func DeleteTripUtilNumber(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()
	//query, err := bindata.Asset("queries\\tripUtilNumber\\insertTripUtilNumber.sql")
	//INSERT INTO tripUtilNumber(code, about, name) VALUES ($1, $2, $3);

	db := goqu.New("postgres", conn.DB)

	var tripUtilNumber TripUtilNumbers

	err := c.Bind(&tripUtilNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	query := db.From("trip_util_numbers").Where(goqu.Ex{"id": c.Param("id")})
	exec := query.Delete()

	result, err := exec.Exec()

	if err != nil {
		log.Print(err.Error())
	}

	rows, err := result.RowsAffected()

	if rows > 0 {
		jr.Code = http.StatusOK
		jr.Message = "TripUtilNumber deleted successfully"
	} else {
		jr.Code = http.StatusInternalServerError
		jr.Message = "Some error ocourred during insert on db"
	}

	c.JSON(jr.Code, jr)

}

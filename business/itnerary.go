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


//TODO: Renomear nome das colunas do banco
type Itinerary struct {
	ID          int64 `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	IdTrip      int64 `db:"id_trip"`
	Image      string `db:"image"`
}

//GetAllItinerarys get all Itinerarys from db
func GetAllItinerarys(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()

	db := goqu.New("postgres", conn.DB)

	//query, err := bindata.Asset("queries\\message\\getAllMessages.sql")

	query, _, err := db.From("itinerary").ToSql()

	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("Query file not found").Error())
		return
	}

	var itineraries []Itinerary

	log.Print(query)

	conn.Select(&itineraries, string(query))

	c.JSON(http.StatusOK, itineraries)
}

//GetItinerary get a Itinerary from db based on ID
func GetItinerary(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()

	db := goqu.New("postgres", conn.DB)

	query, _, err := db.From("itinerary").Where(goqu.Ex{"id":c.Param("id")}).ToSql()

	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("Query file not found").Error())
		return
	}

	var itinerary Itinerary

	conn.Get(&itinerary, string(query))

	c.JSON(http.StatusOK, itinerary)
}

//PostItinerary insert a new Itinerary
func PostItinerary(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()
	//query, err := bindata.Asset("queries\\itinerary\\insertItinerary.sql")
	//INSERT INTO itinerary(code, about, name) VALUES ($1, $2, $3);

	db := goqu.New("postgres", conn.DB)

	var itinerary Itinerary

	err := c.Bind(&itinerary)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	query := db.From("itinerary").Insert(goqu.Record{"title": itinerary.Title, "description":itinerary.Description, "id_trip":itinerary.IdTrip, "image":itinerary.Image})

	result, err := query.Exec()

	if err != nil {
		log.Print(err.Error())
	}

	rows, err := result.RowsAffected()

	if rows > 0 {
		jr.Code = http.StatusOK
		jr.Message = "Itinerary created successfully"
	} else {
		jr.Code = http.StatusInternalServerError
		jr.Message = "Some error ocourred during insert on db"
	}

	c.JSON(jr.Code, jr)

}

func PutItinerary(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()
	//query, err := bindata.Asset("queries\\itinerary\\insertItinerary.sql")
	//INSERT INTO itinerary(code, about, name) VALUES ($1, $2, $3);

	db := goqu.New("postgres", conn.DB)

	var itinerary Itinerary

	err := c.Bind(&itinerary)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	query := db.From("itinerary").Where(goqu.Ex{"id":c.Param("id")})
	exec := query.Update(goqu.Record{"title": itinerary.Title, "description":itinerary.Description, "id_trip":itinerary.IdTrip})

	result, err := exec.Exec()

	if err != nil {
		log.Print(err.Error())
	}

	rows, err := result.RowsAffected()

	if rows > 0 {
		jr.Code = http.StatusOK
		jr.Message = "Itinerary edited successfully"
	} else {
		jr.Code = http.StatusInternalServerError
		jr.Message = "Some error ocourred during insert on db"
	}

	c.JSON(jr.Code, jr)

}

func DeleteItinerary(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()
	//query, err := bindata.Asset("queries\\itinerary\\insertItinerary.sql")
	//INSERT INTO itinerary(code, about, name) VALUES ($1, $2, $3);

	db := goqu.New("postgres", conn.DB)

	var itinerary Itinerary

	err := c.Bind(&itinerary)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	query := db.From("itinerary").Where(goqu.Ex{"id":c.Param("id")})
	exec := query.Delete()

	result, err := exec.Exec()

	if err != nil {
		log.Print(err.Error())
	}

	rows, err := result.RowsAffected()

	if rows > 0 {
		jr.Code = http.StatusOK
		jr.Message = "Itinerary deleted successfully"
	} else {
		jr.Code = http.StatusInternalServerError
		jr.Message = "Some error ocourred during insert on db"
	}

	c.JSON(jr.Code, jr)

}
package business

import (
	"errors"
	"fmt"
	"log"
	"mol/connection"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/doug-martin/goqu.v3"
	_ "gopkg.in/doug-martin/goqu.v3/adapters/postgres"
)

//TODO: Renomear nome das colunas do banco
type Hotel struct {
	ID              int    `db:"id" json:"id"`
	Image           string `db:"image" json:"image"`
	Name            string `db:"name" json:"name"`
	Description     string `db:"description" json:"description"`
	Endereco        string `db:"endereco" json:"endereco"`
	Telefone        string `db:"telefone" json:"telefone"`
	FullDescription string `db:"fullDescription" json:"fullDescription"`
	StarNumbers     int64  `db:"starNumbers" json:"starNumbers"`
}

//GetAllHotels get all Hotels from db
func GetAllHotels(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()

	db := goqu.New("postgres", conn.DB)

	//query, err := bindata.Asset("queries\\message\\getAllMessages.sql")

	query, _, err := db.From("hotels").ToSql()

	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("Query file not found").Error())
		return
	}

	var hotels []Hotel

	log.Print(query)

	conn.Select(&hotels, string(query))

	c.JSON(http.StatusOK, hotels)
}

//GetHotel get a Hotel from db based on ID
func GetHotel(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()

	db := goqu.New("postgres", conn.DB)

	query, _, err := db.From("hotels").Where(goqu.Ex{"id": c.Param("id")}).ToSql()

	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("Query file not found").Error())
		return
	}

	var hotel Hotel

	conn.Get(&hotel, string(query))

	c.JSON(http.StatusOK, hotel)
}

//PostHotel insert a new Hotel
func PostHotel(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()
	//query, err := bindata.Asset("queries\\hotel\\insertHotel.sql")
	//INSERT INTO hotel(code, about, name) VALUES ($1, $2, $3);

	db := goqu.New("postgres", conn.DB)

	var hotel Hotel

	err := c.Bind(&hotel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	query := db.From("hotels").Insert(goqu.Record{"image": hotel.Image, "name": hotel.Name, "description": hotel.Description, "endereco": hotel.Endereco, "telefone": hotel.Telefone, "fullDescription": hotel.FullDescription, "starNumbers": hotel.StarNumbers})

	fmt.Println(query)

	result, err := query.Exec()

	if err != nil {
		log.Print(err.Error())
	}

	rows, err := result.RowsAffected()

	if rows > 0 {
		jr.Code = http.StatusOK
		jr.Message = "Hotel created successfully"
	} else {
		jr.Code = http.StatusInternalServerError
		jr.Message = "Some error ocourred during insert on db"
	}

	c.JSON(jr.Code, jr)

}

func PutHotel(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()
	//query, err := bindata.Asset("queries\\hotel\\insertHotel.sql")
	//INSERT INTO hotel(code, about, name) VALUES ($1, $2, $3);

	db := goqu.New("postgres", conn.DB)

	var hotel Hotel

	err := c.Bind(&hotel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	query := db.From("hotels").Where(goqu.Ex{"id": c.Param("id")})
	exec := query.Update(goqu.Record{"image": hotel.Image, "name": hotel.Name, "description": hotel.Description, "endereco": hotel.Endereco, "telefone": hotel.Telefone, "fullDescription": hotel.FullDescription, "starNumbers": hotel.StarNumbers})

	result, err := exec.Exec()

	if err != nil {
		log.Print(err.Error())
	}

	rows, err := result.RowsAffected()

	if rows > 0 {
		jr.Code = http.StatusOK
		jr.Message = "Hotel edited successfully"
	} else {
		jr.Code = http.StatusInternalServerError
		jr.Message = "Some error ocourred during insert on db"
	}

	c.JSON(jr.Code, jr)

}

func DeleteHotel(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()
	//query, err := bindata.Asset("queries\\hotel\\insertHotel.sql")
	//INSERT INTO hotel(code, about, name) VALUES ($1, $2, $3);

	db := goqu.New("postgres", conn.DB)

	var hotel Hotel

	err := c.Bind(&hotel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	query := db.From("hotels").Where(goqu.Ex{"id": c.Param("id")})
	exec := query.Delete()

	result, err := exec.Exec()

	if err != nil {
		log.Print(err.Error())
	}

	rows, err := result.RowsAffected()

	if rows > 0 {
		jr.Code = http.StatusOK
		jr.Message = "Hotel deleted successfully"
	} else {
		jr.Code = http.StatusInternalServerError
		jr.Message = "Some error ocourred during insert on db"
	}

	c.JSON(jr.Code, jr)

}

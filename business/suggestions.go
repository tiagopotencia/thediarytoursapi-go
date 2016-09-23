package business

import (
	"github.com/gin-gonic/gin"
	"errors"
	"log"
	"git.heroku.com/thediarytoursapi-go/connection"
	"gopkg.in/doug-martin/goqu.v3"
	_ "gopkg.in/doug-martin/goqu.v3/adapters/postgres"
	"net/http"
)

//TODO: Renomear nome das coluas do banco
type Suggestions struct {
	ID int64 `db:"id"`
	Type string `db:"type"`
	Name string `db:"name"`
	Description string `db:"description"`
	IdTrip int64 `db:"id_trip"`
}

//GetAllSuggestions get all Suggestions from db
func GetAllSuggestions(c *gin.Context)  {
	conn := connection.GetConnection()
	defer conn.Close()

	db := goqu.New("postgres", conn.DB)

	query, _, err := db.From("suggestions").ToSql()

	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("Query file not found").Error())
		return
	}

	var suggestions []Suggestions

	log.Print(query)

	conn.Select(&suggestions, string(query))

	c.JSON(http.StatusOK, suggestions)
}

//GetSuggestion get a Suggestion from db based on ID
func GetSuggestion(c *gin.Context)  {
	conn := connection.GetConnection()
	defer conn.Close()

	db := goqu.New("postgres", conn.DB)

	query, _, err := db.From("suggestions").Where(goqu.Ex{"id":c.Param("id")}).ToSql()

	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("Query file not found").Error())
		return
	}

	var suggestion Suggestions

	conn.Get(&suggestion, string(query))

	c.JSON(http.StatusOK, suggestion)
}

//PostSuggestion insert a new Suggestion
func PostSuggestion(c *gin.Context)  {
	conn := connection.GetConnection()
	defer conn.Close()
	//query, err := bindata.Asset("queries\\suggestion\\insertSuggestion.sql")
	//INSERT INTO suggestion(code, about, name) VALUES ($1, $2, $3);

	db := goqu.New("postgres", conn.DB)

	var suggestion Suggestions

	err := c.Bind(&suggestion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	query := db.From("suggestions").Insert(goqu.Record{"type": suggestion.Type, "name":suggestion.Name, "description":suggestion.Description, "id_trip":suggestion.IdTrip})

	result, err := query.Exec()

	if err != nil{
		log.Print(err.Error())
	}

	rows, err := result.RowsAffected()

	if rows > 0{
		jr.Code = http.StatusOK
		jr.Message = "Suggestion created successfully"
	} else {
		jr.Code = http.StatusInternalServerError
		jr.Message = "Some error ocourred during insert on db"
	}

	c.JSON(jr.Code, jr)


}

func PutSuggestion(c *gin.Context)  {
	conn := connection.GetConnection()
	defer conn.Close()
	//query, err := bindata.Asset("queries\\suggestion\\insertSuggestion.sql")
	//INSERT INTO suggestion(code, about, name) VALUES ($1, $2, $3);

	db := goqu.New("postgres", conn.DB)

	var suggestion Suggestions

	err := c.Bind(&suggestion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	query := db.From("suggestions").Where(goqu.Ex{"id":c.Param("id")})
	exec := query.Update(goqu.Record{"type": suggestion.Type, "name":suggestion.Name, "description":suggestion.Description, "id_trip":suggestion.IdTrip})

	result, err := exec.Exec()

	if err != nil{
		log.Print(err.Error())
	}

	rows, err := result.RowsAffected()

	if rows > 0{
		jr.Code = http.StatusOK
		jr.Message = "Suggestion edited successfully"
	} else {
		jr.Code = http.StatusInternalServerError
		jr.Message = "Some error ocourred during insert on db"
	}

	c.JSON(jr.Code, jr)


}

func DeleteSuggestion(c *gin.Context)  {
	conn := connection.GetConnection()
	defer conn.Close()
	//query, err := bindata.Asset("queries\\suggestion\\insertSuggestion.sql")
	//INSERT INTO suggestion(code, about, name) VALUES ($1, $2, $3);

	db := goqu.New("postgres", conn.DB)

	var suggestion Suggestions

	err := c.Bind(&suggestion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	query := db.From("suggestions").Where(goqu.Ex{"id":c.Param("id")})
	exec := query.Delete()

	result, err := exec.Exec()

	if err != nil{
		log.Print(err.Error())
	}

	rows, err := result.RowsAffected()

	if rows > 0{
		jr.Code = http.StatusOK
		jr.Message = "Suggestion deleted successfully"
	} else {
		jr.Code = http.StatusInternalServerError
		jr.Message = "Some error ocourred during insert on db"
	}

	c.JSON(jr.Code, jr)


}
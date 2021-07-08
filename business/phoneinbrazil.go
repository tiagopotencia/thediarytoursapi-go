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
type PhoneInBrazil struct {
	ID         int64  `db:"id"`
	ParentName string `db:"parent_name"`
	Phone      string `db:"phone"`
	IdUser     int64  `db:"id_user"`
}

//GetAllPhoneInBrazil get all PhoneInBrazil from db
func GetAllPhoneInBrazil(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()

	db := goqu.New("postgres", conn.DB)

	query, _, err := db.From("phone_in_brazil").ToSql()

	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("Query file not found").Error())
		return
	}

	var suggestions []PhoneInBrazil

	log.Print(query)

	conn.Select(&suggestions, string(query))

	c.JSON(http.StatusOK, suggestions)
}

//GetPhoneInBrazil get a PhoneInBrazil from db based on ID
func GetPhoneInBrazil(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()

	db := goqu.New("postgres", conn.DB)

	query, _, err := db.From("phone_in_brazil").Where(goqu.Ex{"id": c.Param("id")}).ToSql()

	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("Query file not found").Error())
		return
	}

	var phoneInBrazil PhoneInBrazil

	conn.Get(&phoneInBrazil, string(query))

	c.JSON(http.StatusOK, phoneInBrazil)
}

//PostPhoneInBrazil insert a new PhoneInBrazil
func PostPhoneInBrazil(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()
	//query, err := bindata.Asset("queries\\phoneInBrazil\\insertPhoneInBrazil.sql")
	//INSERT INTO phoneInBrazil(code, about, name) VALUES ($1, $2, $3);

	db := goqu.New("postgres", conn.DB)

	var phoneInBrazil PhoneInBrazil

	err := c.Bind(&phoneInBrazil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	query := db.From("phone_in_brazil").Insert(goqu.Record{"parent_name": phoneInBrazil.ParentName, "phone": phoneInBrazil.Phone, "id_user": phoneInBrazil.IdUser})

	result, err := query.Exec()

	if err != nil {
		log.Print(err.Error())
	}

	rows, err := result.RowsAffected()

	if rows > 0 {
		jr.Code = http.StatusOK
		jr.Message = "PhoneInBrazil created successfully"
	} else {
		jr.Code = http.StatusInternalServerError
		jr.Message = "Some error ocourred during insert on db"
	}

	c.JSON(jr.Code, jr)

}

func PutPhoneInBrazil(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()
	//query, err := bindata.Asset("queries\\phoneInBrazil\\insertPhoneInBrazil.sql")
	//INSERT INTO phoneInBrazil(code, about, name) VALUES ($1, $2, $3);

	db := goqu.New("postgres", conn.DB)

	var phoneInBrazil PhoneInBrazil

	err := c.Bind(&phoneInBrazil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	query := db.From("phone_in_brazil").Where(goqu.Ex{"id": c.Param("id")})
	exec := query.Update(goqu.Record{"parent_name": phoneInBrazil.ParentName, "phone": phoneInBrazil.Phone, "id_user": phoneInBrazil.IdUser})

	result, err := exec.Exec()

	if err != nil {
		log.Print(err.Error())
	}

	rows, err := result.RowsAffected()

	if rows > 0 {
		jr.Code = http.StatusOK
		jr.Message = "PhoneInBrazil edited successfully"
	} else {
		jr.Code = http.StatusInternalServerError
		jr.Message = "Some error ocourred during insert on db"
	}

	c.JSON(jr.Code, jr)

}

func DeletePhoneInBrazil(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()
	//query, err := bindata.Asset("queries\\phoneInBrazil\\insertPhoneInBrazil.sql")
	//INSERT INTO phoneInBrazil(code, about, name) VALUES ($1, $2, $3);

	db := goqu.New("postgres", conn.DB)

	var phoneInBrazil PhoneInBrazil

	err := c.Bind(&phoneInBrazil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	query := db.From("phone_in_brazil").Where(goqu.Ex{"id": c.Param("id")})
	exec := query.Delete()

	result, err := exec.Exec()

	if err != nil {
		log.Print(err.Error())
	}

	rows, err := result.RowsAffected()

	if rows > 0 {
		jr.Code = http.StatusOK
		jr.Message = "PhoneInBrazil deleted successfully"
	} else {
		jr.Code = http.StatusInternalServerError
		jr.Message = "Some error ocourred during insert on db"
	}

	c.JSON(jr.Code, jr)

}

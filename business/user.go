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
type User struct {
	ID                int64  `db:"id" json:"id"`
	Name              string `db:"name" json:"name"`
	Email             string `db:"email" json:"email"`
	Phone             string `db:"phone" json:"phone"`
	Address           string `db:"address" json:"address"`
	AddressNumber     string `db:"address_number" json:"addressNumber"`
	AddressComplement string `db:"address_complement" json:"addressComplement"`
	ZipCode           string `db:"zip_code" json:"zipCode"`
	IsLeader          bool   `db:"is_leader" json:"isLeader"`
}

func GetAllUsers(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()

	db := goqu.New("postgres", conn.DB)

	query, _, err := db.From("user").ToSql()

	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("Query file not found").Error())
		return
	}

	var users []User

	log.Print(query)

	conn.Select(&users, string(query))

	c.JSON(http.StatusOK, users)
}

//GetUser get a User from db based on ID
func GetUser(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()

	db := goqu.New("postgres", conn.DB)

	query, _, err := db.From("user").Where(goqu.Ex{"id": c.Param("id")}).ToSql()

	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("Query file not found").Error())
		return
	}

	var user User

	conn.Get(&user, string(query))

	c.JSON(http.StatusOK, user)
}

//PostUser insert a new User
func PostUser(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()
	//query, err := bindata.Asset("queries\\user\\insertUser.sql")
	//INSERT INTO user(code, about, name) VALUES ($1, $2, $3);

	db := goqu.New("postgres", conn.DB)

	var user User

	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	//ID                int64 `db:"id"`
	//Name              string `db:"name"`
	//Email             string `db:"email"`
	//Phone             string `db:"phone"`
	//Address           string `db:"address"`
	//AddressNumber     string `db:"address_number"`
	//AddressComplement string `db:"address_complement"`
	//ZipCode           string `db:"zip_code"`
	//IsLeader          bool `db:"is_leader"`

	query := db.From("user").Insert(goqu.Record{"name": user.Name, "email": user.Email, "phone": user.Phone, "address": user.Address,
		"address_number": user.AddressNumber, "address_complement": user.AddressComplement, "zip_code": user.ZipCode, "is_leader": user.IsLeader})

	result, err := query.Exec()

	if err != nil {
		log.Print(err.Error())
	}

	rows, err := result.RowsAffected()

	if rows > 0 {
		jr.Code = http.StatusOK
		jr.Message = "User created successfully"
	} else {
		jr.Code = http.StatusInternalServerError
		jr.Message = "Some error ocourred during insert on db"
	}

	c.JSON(jr.Code, jr)

}

func PutUser(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()
	//query, err := bindata.Asset("queries\\user\\insertUser.sql")
	//INSERT INTO user(code, about, name) VALUES ($1, $2, $3);

	db := goqu.New("postgres", conn.DB)

	var user User

	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	query := db.From("user").Where(goqu.Ex{"id": c.Param("id")})
	exec := query.Update(goqu.Record{"name": user.Name, "email": user.Email, "phone": user.Phone, "address": user.Address,
		"address_number": user.AddressNumber, "address_complement": user.AddressComplement, "zip_code": user.ZipCode, "is_leader": user.IsLeader})

	result, err := exec.Exec()

	if err != nil {
		log.Print(err.Error())
	}

	rows, err := result.RowsAffected()

	if rows > 0 {
		jr.Code = http.StatusOK
		jr.Message = "User edited successfully"
	} else {
		jr.Code = http.StatusInternalServerError
		jr.Message = "Some error ocourred during insert on db"
	}

	c.JSON(jr.Code, jr)

}

func DeleteUser(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()
	//query, err := bindata.Asset("queries\\user\\insertUser.sql")
	//INSERT INTO user(code, about, name) VALUES ($1, $2, $3);

	db := goqu.New("postgres", conn.DB)

	var user User

	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	query := db.From("user").Where(goqu.Ex{"id": c.Param("id")})
	exec := query.Delete()

	result, err := exec.Exec()

	if err != nil {
		log.Print(err.Error())
	}

	rows, err := result.RowsAffected()

	if rows > 0 {
		jr.Code = http.StatusOK
		jr.Message = "User deleted successfully"
	} else {
		jr.Code = http.StatusInternalServerError
		jr.Message = "Some error ocourred during insert on db"
	}

	c.JSON(jr.Code, jr)

}

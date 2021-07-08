package business

import (
	"errors"
	"log"
	"mol/connection"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/doug-martin/goqu.v3"
	_ "gopkg.in/doug-martin/goqu.v3/adapters/postgres"
)

//TODO: Renomear nome das coluas do banco
type Message struct {
	IsRead     bool       `db:"is_read"`
	ID         int64      `db:"id"`
	Message    string     `db:"message"`
	SentDate   *time.Time `db:"sent_date"`
	IdUserFrom int64      `db:"id_user_from"`
	IdUserTo   int64      `db:"id_user_to"`
}

func GetAllMessages(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()

	db := goqu.New("postgres", conn.DB)

	//query, err := bindata.Asset("queries\\message\\getAllMessages.sql")

	query, _, err := db.From("message").ToSql()

	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("Query file not found").Error())
		return
	}

	var messages []Message

	log.Print(query)

	conn.Select(&messages, string(query))

	c.JSON(http.StatusOK, messages)
}

//GetMessage get a Message from db based on ID
func GetMessage(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()

	db := goqu.New("postgres", conn.DB)

	query, _, err := db.From("message").Where(goqu.Ex{"id": c.Param("id")}).ToSql()

	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("Query file not found").Error())
		return
	}

	var message Message

	conn.Get(&message, string(query))

	c.JSON(http.StatusOK, message)
}

//PostMessage insert a new Message
func PostMessage(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()
	//query, err := bindata.Asset("queries\\message\\insertMessage.sql")
	//INSERT INTO message(code, about, name) VALUES ($1, $2, $3);

	db := goqu.New("postgres", conn.DB)

	var message Message

	err := c.Bind(&message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	query := db.From("message").Insert(goqu.Record{"is_read": message.IsRead, "message": message.Message, "sent_date": message.SentDate, "id_user_from": message.IdUserFrom, "id_user_to": message.IdUserTo})

	result, err := query.Exec()

	if err != nil {
		log.Print(err.Error())
	}

	rows, err := result.RowsAffected()

	if rows > 0 {
		jr.Code = http.StatusOK
		jr.Message = "Message created successfully"
	} else {
		jr.Code = http.StatusInternalServerError
		jr.Message = "Some error ocourred during insert on db"
	}

	c.JSON(jr.Code, jr)

}

func PutMessage(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()
	//query, err := bindata.Asset("queries\\message\\insertMessage.sql")
	//INSERT INTO message(code, about, name) VALUES ($1, $2, $3);

	db := goqu.New("postgres", conn.DB)

	var message Message

	err := c.Bind(&message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	query := db.From("message").Where(goqu.Ex{"id": c.Param("id")})
	exec := query.Update(goqu.Record{"is_read": message.IsRead, "message": message.Message, "sent_date": message.SentDate, "id_user_from": message.IdUserFrom, "id_user_to": message.IdUserTo})

	result, err := exec.Exec()

	if err != nil {
		log.Print(err.Error())
	}

	rows, err := result.RowsAffected()

	if rows > 0 {
		jr.Code = http.StatusOK
		jr.Message = "Message edited successfully"
	} else {
		jr.Code = http.StatusInternalServerError
		jr.Message = "Some error ocourred during insert on db"
	}

	c.JSON(jr.Code, jr)

}

func DeleteMessage(c *gin.Context) {
	conn := connection.GetConnection()
	defer conn.Close()
	//query, err := bindata.Asset("queries\\message\\insertMessage.sql")
	//INSERT INTO message(code, about, name) VALUES ($1, $2, $3);

	db := goqu.New("postgres", conn.DB)

	var message Message

	err := c.Bind(&message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	query := db.From("message").Where(goqu.Ex{"id": c.Param("id")})
	exec := query.Delete()

	result, err := exec.Exec()

	if err != nil {
		log.Print(err.Error())
	}

	rows, err := result.RowsAffected()

	if rows > 0 {
		jr.Code = http.StatusOK
		jr.Message = "Message deleted successfully"
	} else {
		jr.Code = http.StatusInternalServerError
		jr.Message = "Some error ocourred during insert on db"
	}

	c.JSON(jr.Code, jr)

}

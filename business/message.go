package business

import "time"


//TODO: Renomear nome das coluas do banco
type Message struct {
	ID         int64 `db:"Id"`
	IsRead     bool `db:"Is_Read"`
	Message    string `db:"Message"`
	SentDate   *time.Time `db:"Sent_Date"`
	IdUserFrom int64 `db:"Id_User_From"`
	IdUserTo   int64 `db:"Id_User_To"`
}
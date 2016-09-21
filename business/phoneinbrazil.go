package business

//TODO: Renomear nome das coluas do banco
type PhoneInBrazil struct {

	ParentName string `db:"Parent_Name"`
	Phone string `db:"Phone"`
	IdUser int64 `db:"Id_User"`

}
package business

//TODO: Renomear nome das coluas do banco
type User struct {

	ID int64 `db:"Id"`
	Name string `db:"Name"`
	Email string `db:"Email"`
	Phone string `db:"Phone"`
	Address string `db:"Address"`
	AddressNumber string `db:"Address_Number"`
	AddressComplement string `db:"Address_Complement"`
	ZipCode string `db:"Zip_Code"`
	IsLeader bool `db:"Is_Leader"`

}
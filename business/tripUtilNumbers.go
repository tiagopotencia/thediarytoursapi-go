package business

//TODO: Renomear nome das coluas do banco
type TripUtilNumbers struct {

	ID int64 `db:"Id"`
	Type string `db:"Type"`
	Phone string `db:"Phone"`
	Description string `db:"Description"`
	IdTrip int64 `db:"Id_Trip"`

}

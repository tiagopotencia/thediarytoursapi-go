package business

//TODO: Renomear nome das coluas do banco
type TripUtilNumbers struct {
	ID          int64  `db:"id"`
	Type        string `db:"type"`
	Phone       string `db:"phone"`
	Description string `db:"description"`
	IdTrip      int64 `db:"id_trip"`
}

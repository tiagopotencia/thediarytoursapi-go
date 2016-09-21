package business

//TODO: Renomear nome das coluas do banco
type TripUser struct {

	IdTrip int64 `db:"Id_Trip"`
	IdUser int64 `db:"Id_User"`
}
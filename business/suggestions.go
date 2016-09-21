package business

//TODO: Renomear nome das coluas do banco
type Suggestions struct {
	ID int64 `db:"id"`
	Type string `db:"type"`
	Name string `db:"name"`
	Description string `db:"description"`
	IdTrip int64 `db:"id_trip"`
}
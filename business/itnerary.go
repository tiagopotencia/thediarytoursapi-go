package business

//TODO: Renomear nome das coluas do banco
type Itnerary struct {
	ID          int64 `db:"Id"`
	Title       string `db:"Title"`
	Description string `db:"Description"`
	IDTrip      int64 `db:"Id_Trip"`
	Name        string `db:"Name"`
}
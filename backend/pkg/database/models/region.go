package models

type Region struct {
	Id     uint   `db:"id"`
	Nombre string `db:"nombre"`
	Numero uint   `db:"numero"`
}

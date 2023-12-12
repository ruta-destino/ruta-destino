package models

type Empresa struct {
	Id     uint   `db:"id"`
	Nombre string `db:"nombre"`
}

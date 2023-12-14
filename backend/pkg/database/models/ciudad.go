package models

type Ciudad struct {
	Id              uint   `db:"id"`
	Nombre          string `db:"nombre"`
	IdProvincia     uint   `db:"id_provincia"`
	NombreProvincia string `db:"nombre_provincia"`
}

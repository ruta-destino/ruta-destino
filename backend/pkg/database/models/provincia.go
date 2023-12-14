package models

type Provincia struct {
	Id           uint   `db:"id"`
	Nombre       string `db:"nombre"`
	IdRegion     uint   `db:"id_region"`
	NombreRegion string `db:"nombre_region"`
}

package models

type Terminal struct {
	Id           uint    `db:"id"`
	Nombre       string  `db:"nombre"`
	Longitud     float64 `db:"longitud"`
	Latitud      float64 `db:"latitud"`
	Direccion    string  `db:"direccion"`
	IdCiudad     uint    `db:"id_ciudad"`
	NombreCiudad string  `db:"nombre_ciudad"`
}

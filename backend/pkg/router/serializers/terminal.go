package serializers

type Terminal struct {
	Id        uint    `json:"id"`
	Nombre    string  `json:"nombre"`
	Longitud  float64 `json:"longitud"`
	Latitud   float64 `json:"latitud"`
	Direccion string  `json:"direccion"`
	IdCiudad  uint    `json:"id_ciudad"`
}

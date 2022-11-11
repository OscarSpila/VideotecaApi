package dtos

type NuevoAlquilerDTO struct {
	Importe float64 `json:"importe"`
	Abonado bool    `json:"abonado"`
	SocioID int     `json:"socioID"`
}

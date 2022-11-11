package dtos

type NuevaDevolucionDTO struct {
	Nombre            string `json:"firstName"`
	Apellido          string `json:"lastName"`
	FechaNacimiento   string `json:"birthday"`
	CorreoElectronico string `json:"email"`
	NombreDocumento   string `json:"documentName"`
}

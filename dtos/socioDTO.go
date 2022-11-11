package dtos

type NuevoSocioDTO struct {
	Nombre            string `json:"firstName"`
	Apellido          string `json:"lastName"`
	FechaNacimiento   string `json:"birthday"`
	CorreoElectronico string `json:"email"`
	NombreDocumento   string `json:"documentName"`
	NumeroDocumento   string `json:"documentNumber"`
}

type SelectSocioDTO struct {
	ID                int    `json:"id"`
	Nombre            string `json:"firstName"`
	Apellido          string `json:"lastName"`
	FechaNacimiento   string `json:"birthday"`
	CorreoElectronico string `json:"email"`
	TipoDocumento     string `json:"ducumentType"`
	NombreDocumento   string `json:"documentName"`
	NumeroDocumento   string `json:"documentNumber"`
}

type ModificarSocioDTO struct {
	Nombre            string `json:"firstName"`
	Apellido          string `json:"lastName"`
	FechaNacimiento   string `json:"birthday"`
	CorreoElectronico string `json:"email"`
	NombreDocumento   string `json:"documentName"`
	NumeroDocumento   string `json:"documentNumber"`
}

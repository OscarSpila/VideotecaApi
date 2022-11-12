package dtos

type NuevaPeliculaDTO struct {
	Nombre       string `json:"nombre"`
	Idioma       string `json:"idioma"`
	Productora   string `json:"productora"`
	Actores      string `json:"actores"`
	PaisDeOrigen string `json:"paisDeOrigen"`
	GeneroID     int    `json:"generoTypeID"`
}

type SelectPeliculaDTO struct {
	Nombre       string `json:"nombre"`
	Idioma       string `json:"idioma"`
	Productora   string `json:"productora"`
	Actores      string `json:"actores"`
	PaisDeOrigen string `json:"paisDeOrigen"`
	GeneroNombre string `json:"genero"`
}

type ModificarPeliculaDTO struct {
	Nombre       string `json:"nombre"`
	Idioma       string `json:"idioma"`
	Productora   string `json:"productora"`
	Actores      string `json:"actores"`
	PaisDeOrigen string `json:"paisDeOrigen"`
	GeneroNombre string `json:"genero"`
}

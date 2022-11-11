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
	ID           int    `json:"id"`
	Nombre       string `json:"nombre"`
	Idioma       string `json:"idioma"`
	Productora   string `json:"productora"`
	Actores      string `json:"actores"`
	PaisDeOrigen string `json:"paisDeOrigen"`
	GeneroID     int    `json:"generoTypeID"`
	Genero       string `json:"genero"`
}

type ModificarPeliculaDTO struct {
	Nombre       string `json:"nombre"`
	Idioma       string `json:"idioma"`
	Productora   string `json:"productora"`
	Actores      string `json:"actores"`
	PaisDeOrigen string `json:"paisDeOrigen"`
	GeneroID     int    `json:"generoTypeID"`
	Genero       string `json:"genero"`
}
type DatosPeliculaFiltradoDTO struct {
	Nombre       string `json:"nombre"`
	Idioma       string `json:"idioma"`
	Productora   string `json:"productora"`
	GeneroID     int    `json:"generoTypeID"`
	GeneroNombre string `json:"generoNombre"`
}

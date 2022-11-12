package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"videotecaapi/dtos"
	"videotecaapi/models"
	"videotecaapi/repositories"
	"videotecaapi/utils"
)

type PeliculaController struct{}

func (controller PeliculaController) Get(context *gin.Context) {
	ID := context.Param("peliculaID")

	id, err := strconv.Atoi(ID) // se convierte un string a int
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.Error("ID is missing."))
		return
	}

	entityRep := new(repositories.PeliculaRepository)
	entity := entityRep.Get(id)

	// Paso el registro devuelto a su correspondiente DTO
	entityPeliculaDTO := new(dtos.SelectPeliculaDTO)

	if entity.ID == 0 {
		context.JSON(http.StatusNotFound, "")
	} else {

		entityPeliculaDTO.Nombre = entity.Nombre
		entityPeliculaDTO.Actores = entity.Actores
		entityPeliculaDTO.PaisDeOrigen = entity.PaisDeOrigen
		entityPeliculaDTO.Productora = entity.Productora
		entityPeliculaDTO.GeneroNombre = entity.Genero.Nombre
		entityPeliculaDTO.Idioma = entity.Idioma

		// context.JSON(http.StatusOK, entity)

		context.JSON(http.StatusOK, entityPeliculaDTO)
	}
}

func (controller PeliculaController) Find(context *gin.Context) {

	entityRep := new(repositories.PeliculaRepository)
	entities := entityRep.Find(context)
	context.JSON(http.StatusOK, entities)
}

func (controller PeliculaController) Create(context *gin.Context) {

	entity := new(models.Pelicula)
	entityDTO := new(dtos.NuevaPeliculaDTO)

	if err := context.BindJSON(&entityDTO); err != nil {
		context.JSON(http.StatusBadRequest, utils.Error(err.Error()))
		return
	}

	entityGenero := getGeneroType(entityDTO.GeneroID)

	if entityGenero.ID == 0 {
		context.JSON(http.StatusBadRequest, "Genero no encontrado")
		return
	}

	entity.Nombre = entityDTO.Nombre
	entity.Actores = entityDTO.Actores
	entity.GeneroID = entityDTO.GeneroID
	entity.Idioma = entityDTO.Idioma
	entity.PaisDeOrigen = entityDTO.PaisDeOrigen
	entity.Productora = entityDTO.Productora

	rep := new(repositories.PeliculaRepository)
	id := rep.Insert(*entity)
	if id == 0 {
		context.JSON(http.StatusBadRequest, "Ya existe esta pelicula")
		return
	}
	context.JSON(http.StatusCreated, id)

}

func (controller PeliculaController) Update(context *gin.Context) {

	entityDTO := new(dtos.ModificarPeliculaDTO)
	entityDB := new(models.Pelicula)

	// Se convierte el Json al objeto DTO
	if err := context.BindJSON(&entityDTO); err != nil {
		context.JSON(http.StatusBadRequest, utils.Error(err.Error())) // NO MOSTRAR ERRORES DIRECTOS del SISTEMA. (NO SE HACE!!!!)
		return
	}

	ID := context.Param("peliculaID")

	id, err := strconv.Atoi(ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.Error("ID is missing."))
		return
	}

	rep := new(repositories.PeliculaRepository)
	repo := new(repositories.GeneroRepository)

	entityDB = rep.Get(id)
	entityGenero := repo.GetByGeneroName(entityDTO.GeneroNombre)

	if entityDB == nil {
		context.JSON(http.StatusNotFound, "")
		return
	}

	entityDB.Idioma = entityDTO.Idioma
	entityDB.Productora = entityDTO.Productora
	entityDB.Actores = entityDTO.Actores
	entityDB.PaisDeOrigen = entityDTO.PaisDeOrigen
	entityDB.GeneroID = int(entityGenero.ID)
	entityDB.Nombre = entityDTO.Nombre

	rowAffected := rep.Update(id, *entityDB)

	if rowAffected == 0 {
		context.JSON(http.StatusNotFound, "")
	} else {
		context.JSON(http.StatusOK, rowAffected)
	}
}

func (controller PeliculaController) Delete(context *gin.Context) {
	ID := context.Param("peliculaID")

	id, err := strconv.Atoi(ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.Error("ID is missing."))
		return
	}

	entityRep := new(repositories.PeliculaRepository)
	rowAffected := entityRep.Delete(id)

	if rowAffected == 0 {
		context.JSON(http.StatusNotFound, "")
	} else {
		context.JSON(http.StatusOK, "")
	}
}

func getGeneroType(GeneroID int) (entityGenero models.Genero) {

	repGenero := new(repositories.GeneroRepository)

	entityGenero = *repGenero.Get(GeneroID)

	return entityGenero
}

func getGeneroName(genero string) (entityGenero models.Genero) {

	// Busco el nombre del Documento en la tabla Tipo de Documento, esto con dos fines:

	repGenero := new(repositories.GeneroRepository)

	entityGenero = *repGenero.GetByGeneroName(genero)

	return entityGenero
}

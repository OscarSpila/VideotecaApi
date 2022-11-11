package controllers

import (
	"net/http"
	"strconv"
	"strings"

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

	if entity.ID == 0 {
		context.JSON(http.StatusNotFound, "")
	} else {
		context.JSON(http.StatusOK, entity)
	}
}

func (controller PeliculaController) Find(context *gin.Context) {

	entityRep := new(repositories.PeliculaRepository)
	entities := entityRep.Find(context)
	context.JSON(http.StatusOK, entities)
}

func (controller PeliculaController) Create(context *gin.Context) {

	entityDTO := new(dtos.NuevaPeliculaDTO)
	entity := new(models.Pelicula)

	if err := context.BindJSON(&entityDTO); err != nil {
		context.JSON(http.StatusBadRequest, utils.Error(err.Error()))
		return
	}

	//listError := checkeoNuevosInputs(*entityDTO)
	//if len(listError) != 0 {
	//	context.JSON(http.StatusBadRequest, utils.Error(listError))
	//	return
	//}

	// Se busca el ID del genero de la pelicula en la tabla Genero, esto con dos fines:
	// 1 * Verificar que existe el Genero.
	// 2 * Se debe obtener el ID del genero para guardarlo en la tabla Socio
	entityGenero := getGeneroID(entityDTO.GeneroID)

	if entityGenero.ID == 0 {
		context.JSON(http.StatusBadRequest, "Genero no encontrado")
		return
	}

	entity.Nombre = entityDTO.Nombre
	entity.Idioma = entityDTO.Idioma
	entity.Productora = entityDTO.Productora
	entity.Actores = entityDTO.Actores
	entity.PaisDeOrigen = entityDTO.PaisDeOrigen
	entity.Genero = entityGenero

	rep := new(repositories.PeliculaRepository)

	id, err := rep.Insert(*entity)

	if err != nil {
		context.JSON(http.StatusInternalServerError, "Ya existe una pelicula con ese nombre!")
		//context.JSON(http.StatusInternalServerError, err.Error())
	} else {
		context.JSON(http.StatusCreated, id)
	}
}

func (controller PeliculaController) Update(context *gin.Context) {

	entity := new(models.Pelicula)

	ID := context.Param("peliculaID")

	id, err := strconv.Atoi(ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.Error("ID is missing."))
		return
	}

	if err := context.BindJSON(&entity); err != nil {
		context.JSON(http.StatusBadRequest, utils.Error(err.Error())) // NO MOSTRAR ERRORES DIRECTOS del SISTEMA. (NO SE HACE!!!!)
		return
	}

	rep := new(repositories.PeliculaRepository)
	rowAffected := rep.Update(id, *entity)

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

func checkeoNuevosInputs(entityDTO dtos.NuevaPeliculaDTO) (listError string) {
	listError = ""

	// Verifico que haya indicado el Nombre de la pelicula.
	if len(strings.TrimSpace(entityDTO.Nombre)) == 0 {
		listError = "Debe indicar el nombre de la pelicula.\r\n"
	}

	// Verifico que haya indicado el idioma.
	if len(strings.TrimSpace(entityDTO.Idioma)) == 0 {
		listError += "Debe indicar el idioma de la película.\r\n"
	}

	// Verifico que haya indicado la productora.
	if len(strings.TrimSpace(entityDTO.Productora)) == 0 {
		listError += "Debe indicar la productora de la pelicula.\r\n"
	}

	// Verifico que haya indicado los actores.
	if len(strings.TrimSpace(entityDTO.Actores)) == 0 {
		listError += "Debe indicar los Actores de la película.\r\n"
	}
	// Verifico que haya indicado el Pais de origen.
	if len(strings.TrimSpace(entityDTO.PaisDeOrigen)) == 0 {
		listError += "Debe indicar el Pais de Origen de la película.\r\n"
	}

	// Verificar si la fecha de alquiler es valida.
	// Verificar el tipo de genero.

	return listError
}

func getGeneroID(generoID int) (entityGenero models.Genero) {

	// Busco el ID del genero en la tabla , esto con dos fines:

	repGenero := new(repositories.GeneroRepository)

	entityGenero = *repGenero.Get(generoID)

	return entityGenero
}

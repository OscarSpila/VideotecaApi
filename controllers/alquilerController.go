package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"videotecaapi/models"
	"videotecaapi/repositories"
	"videotecaapi/utils"
)

type AlquilerController struct{}

func (controller AlquilerController) Get(context *gin.Context) {
	ID := context.Param("alquilerID")

	id, err := strconv.Atoi(ID) // se convierte un string a int
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.Error("ID is missing."))
		return
	}

	entityRep := new(repositories.AlquilerRepository)
	entity := entityRep.Get(id)

	if entity.ID == 0 {
		context.JSON(http.StatusNotFound, "")
	} else {
		context.JSON(http.StatusOK, entity)
	}
}

func (controller AlquilerController) Find(context *gin.Context) {

	entityRep := new(repositories.AlquilerRepository)
	entities := entityRep.Find(context)
	context.JSON(http.StatusOK, entities)
}

func (controller AlquilerController) Create(context *gin.Context) {

	entity := new(models.Alquiler)

	if err := context.BindJSON(&entity); err != nil {
		context.JSON(http.StatusBadRequest, utils.Error(err.Error()))
		return
	}

	rep := new(repositories.AlquilerRepository)
	id := rep.Insert(*entity)

	context.JSON(http.StatusCreated, id)
}

func (controller AlquilerController) Update(context *gin.Context) {

	entity := new(models.Alquiler)

	ID := context.Param("alquilerID")

	id, err := strconv.Atoi(ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.Error("ID is missing."))
		return
	}

	if err := context.BindJSON(&entity); err != nil {
		context.JSON(http.StatusBadRequest, utils.Error(err.Error())) // NO MOSTRAR ERRORES DIRECTOS del SISTEMA. (NO SE HACE!!!!)
		return
	}

	rep := new(repositories.AlquilerRepository)
	rowAffected := rep.Update(id, *entity)

	if rowAffected == 0 {
		context.JSON(http.StatusNotFound, "")
	} else {
		context.JSON(http.StatusOK, rowAffected)
	}
}

func (controller AlquilerController) Delete(context *gin.Context) {
	ID := context.Param("alquilerID")

	id, err := strconv.Atoi(ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.Error("ID is missing."))
		return
	}

	entityRep := new(repositories.AlquilerRepository)
	rowAffected := entityRep.Delete(id)

	if rowAffected == 0 {
		context.JSON(http.StatusNotFound, "")
	} else {
		context.JSON(http.StatusOK, "")
	}
}

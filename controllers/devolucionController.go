package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"videotecaapi/models"
	"videotecaapi/repositories"
	"videotecaapi/utils"
)

type DevolucionController struct{}

func (controller DevolucionController) Get(context *gin.Context) {
	ID := context.Param("devolucionID")

	id, err := strconv.Atoi(ID) // se convierte un string a int
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.Error("ID is missing."))
		return
	}

	entityRep := new(repositories.DevolucionRepository)
	entity := entityRep.Get(id)

	if entity.ID == 0 {
		context.JSON(http.StatusNotFound, "")
	} else {
		context.JSON(http.StatusOK, entity)
	}
}

func (controller DevolucionController) Create(context *gin.Context) {

	entity := new(models.Devolucion)

	if err := context.BindJSON(&entity); err != nil {
		context.JSON(http.StatusBadRequest, utils.Error(err.Error()))
		return
	}

	rep := new(repositories.DevolucionRepository)
	id := rep.Insert(*entity)

	context.JSON(http.StatusCreated, id)
}

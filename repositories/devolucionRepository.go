package repositories

import (
	"videotecaapi/db"
	"videotecaapi/models"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
)

// ModelRepository ...
type DevolucionRepository struct{}

func (rep DevolucionRepository) Find(context *gin.Context) paginate.Page {
	db := db.DBConn
	pg := paginate.New()

	model := db.Joins("Devolucion").Model(&models.Devolucion{})

	return pg.With(model).Request(context.Request).Response(&[]models.Devolucion{})
}

func (rep DevolucionRepository) Get(id int) *models.Devolucion {

	entity := new(models.Devolucion)

	db := db.DBConn
	db.First(&entity, id)

	return entity
}

func (rep DevolucionRepository) Insert(entity models.Devolucion) uint {
	db := db.DBConn

	db.Create(&entity)

	return entity.ID
}

func (rep DevolucionRepository) Delete(ID int) int {

	entityToDelete := new(models.Devolucion)

	db := db.DBConn
	db.First(&entityToDelete, ID)

	result := db.Delete(&entityToDelete)

	return int(result.RowsAffected)
}

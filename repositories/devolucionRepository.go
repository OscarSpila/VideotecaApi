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

	model := db.Preload("Socio.TipoDocumento").Preload("Peliculas.Genero").Model(&models.Devolucion{})

	return pg.With(model).Request(context.Request).Response(&[]models.Devolucion{})
}

func (rep DevolucionRepository) Get(id int) *models.Devolucion {

	entity := new(models.Devolucion)

	db := db.DBConn
	db.Preload("Socio.TipoDocumento").Preload("Peliculas.Genero").First(&entity, id)

	return entity
}

func (rep DevolucionRepository) Insert(entity models.Devolucion) uint {
	db := db.DBConn

	db.Create(&entity)

	return entity.ID
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type Socio struct {
	gorm.Model                  // -> el gorm.Model implementa el ID, CreatedAt, UpdatedAt, DeletedAt
	Nombre            string    `json:"firstName"`
	Apellido          string    `json:"lastName"`
	FechaNacimiento   time.Time `json:"birthday"`
	CorreoElectronico string    `json:"email"`
	TipoDocumentoID   int       `gorm:"uniqueIndex:idx_unicosocio" json:"documentTypeID"`
	TipoDocumento     TipoDocumento
	NumeroDoc         string `gorm:"uniqueIndex:idx_unicosocio" json:"documentNumber"`
}

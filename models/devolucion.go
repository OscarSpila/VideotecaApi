package models

import (
	"time"

	"gorm.io/gorm"
)

type Devolucion struct {
	gorm.Model
	FechaDevolucion time.Time
	Socio           Socio
	SocioID         int
}

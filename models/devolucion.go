package models

import (
	"time"

	"gorm.io/gorm"
)

type Devolucion struct {
	gorm.Model              // -> el gorm.Model implementa el ID, CreatedAt, UpdatedAt, DeletedAt
	Id                      int
	socio                   Socio
	FechaDevolucion         time.Time
	listaPeliculaAlquileres []*PeliculaAlquiler `gorm:"many2many:alquiler_pelicula;"`

	// NOTA:
	// `gorm:"many2many:alquiler_pelicula;"` AUTOMATICAMENTE el ORM Gorm crea la tabla alquier_pelicula
	// Para evitar una relacion de muchos a muchos y así respetar el modelo de datos Relacional
}

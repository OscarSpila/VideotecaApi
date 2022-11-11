package models

import (
	"time"

	"gorm.io/gorm"
)

type Alquiler struct {
	gorm.Model          // -> el gorm.Model implementa el ID, CreatedAt, UpdatedAt, DeletedAt
	Importe     float64 `json:"importe"`
	FechaCierre time.Time
	Abonado     bool `json:"abonado"`
	SocioID     int  `json:"socioID"`
	Socio       Socio
	Peliculas   []*Pelicula `gorm:"many2many:alquiler_pelicula;"`

	// NOTA:
	// `gorm:"many2many:alquiler_pelicula;"` AUTOMATICAMENTE el ORM Gorm crea la tabla alquier_pelicula
	// Para evitar una relacion de muchos a muchos y así respetar el modelo de datos Relacional
}

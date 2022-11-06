package models

import (
	"gorm.io/gorm"
)

type PeliculaAlquiler struct {
	gorm.Model // -> el gorm.Model implementa el ID, CreatedAt, UpdatedAt, DeletedAt
	pelicula   Pelicula
	alquiler   Alquiler
}

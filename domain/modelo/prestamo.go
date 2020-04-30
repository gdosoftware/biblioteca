package modelo

import "time"

type Prestamo struct {
	Socio    Socio
	Libro    Libro
	Prestado time.Time
	Devuelto time.Time
}

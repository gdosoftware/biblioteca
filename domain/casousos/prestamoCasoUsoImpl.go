package casousos

import (
	"time"

	"github.com/gdosoftware/biblioteca/domain/interfaces"
	"github.com/gdosoftware/biblioteca/domain/modelo"
)

type PrestamoCasoUsoImpl struct {
	repoSocio    *interfaces.ISocioRepository
	repoLibro    *interfaces.ILibroRepository
	repoPrestamo *interfaces.IPrestamoRepository
}

func (r *PrestamoCasoUsoImpl) CreatePrestamo(socioId string, libroId string) (*modelo.Prestamo, error) {
	socio, _ := r.repoSocio.Retrieve(socioId)
	libro, _ := r.repoLibro.Retrieve(libroId)
	prestamo := modelo.Prestamo{Libro: libro, Socio: socio, Prestado: time.Now()}
	return r.repoPrestamo.Create(&prestamo)
}

func (r *PrestamoCasoUsoImpl) UpdatePrestamo(prestamoId string) (*modelo.Prestamo, error) {
	prestamo, _ := r.repoPrestamo.Retrieve(prestamoId)
	// prestamo.Devuelto:=time.Now()
	return r.repoPrestamo.Update(prestamoId, prestamo)
}

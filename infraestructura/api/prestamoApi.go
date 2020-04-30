package api

import (
	"github.com/gdosoftware/biblioteca/domain/interfaces"
	"github.com/gdosoftware/biblioteca/domain/modelo"
)

type PrestamoApi struct {
	caso interfaces.IPrestamoCasoUso
}

func (p *PrestamoApi) prestarLibro(socioId string, libroId string) (modelo.Prestamo, error) {
	return p.caso.CreatePrestamo(socioId, libroId)
}

func (p *PrestamoApi) devolverLibro(prestamoId string) (modelo.Prestamo, error) {
	return p.caso.DevolucionPrestamo(prestamoId)
}


package interfaces

import "github.com/gdosoftware/biblioteca/domain/modelo"


type  IPrestamoCasoUso interface {
    CreatePrestamo(socioId string, libroId string) (modelo.Prestamo,error)
    UpdatePrestamo(prestamoId string) (modelo.Prestamo, error)
}

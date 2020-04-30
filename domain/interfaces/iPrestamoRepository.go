package interfaces

import "github.com/gdosoftware/biblioteca/domain/modelo"

type IPrestamoRepository interface {
	Create(prestamo *modelo.Prestamo) (*modelo.Prestamo, error)
	Update(id string, prestamo *modelo.Prestamo) (*modelo.Prestamo, error)
	Retrieve(id string) (*modelo.Prestamo, error)
	Delete(id string) error
	FindAll() ([]modelo.Prestamo, error)
}

package interfaces

import "github.com/gdosoftware/biblioteca/domain/modelo"

type ILibroRepository interface {
	Create(libro *modelo.Libro) (*modelo.Libro, error)
	Update(id string, libro *modelo.Libro) (*modelo.Libro, error)
	Retrieve(id string) (modelo.Libro, error)
	Delete(id string) error
	FindAll() ([]modelo.Libro, error)
}

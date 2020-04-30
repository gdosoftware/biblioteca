package interfaces

import "github.com/gdosoftware/biblioteca/domain/modelo"

type ILibroCasoUso interface {
	CreateLibro(libro *modelo.Libro) (*modelo.Libro, error)
	UpdateLibro(id string, libro *modelo.Libro) (*modelo.Libro, error)
	RetrieveLibro(id string) (modelo.Libro, error)
	DeleteLibro(id string) error
	FindAllLibro() ([]modelo.Libro, error)
}

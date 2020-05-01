package casousos

import (
	"github.com/gdosoftware/biblioteca/domain/interfaces"
	"github.com/gdosoftware/biblioteca/domain/modelo"
)

type LibroCasoUsoImpl struct {
	repo interfaces.ILibroRepository
}

func (r *LibroCasoUsoImpl) CreateLibro(libro *modelo.Libro) (*modelo.Libro, error) {
	return r.repo.Create(libro)
}

func (r *LibroCasoUsoImpl) RetrieveLibro(id string) (modelo.Libro, error) {
	return r.repo.Retrieve(id)
}

func (r *LibroCasoUsoImpl) UpdateLibro(id string, libro *modelo.Libro) (*modelo.Libro, error) {
	return r.repo.Update(id, libro)
}

func (r *LibroCasoUsoImpl) DeleteLibro(id string) error {
	return r.repo.Delete(id)
}

func (r *LibroCasoUsoImpl) FindAllLibro() ([]modelo.Libro, error) {
	return r.repo.FindAll()
}

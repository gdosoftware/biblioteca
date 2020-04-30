package casousos

import (
	"github.com/gdosoftware/biblioteca/domain/interfaces"
	"github.com/gdosoftware/biblioteca/domain/modelo"
)

type LibroCasoUsoImpl struct {
	repo interfaces.ILibroRepository
}

func (r *LibroCasoUsoImpl) createLibro(libro *modelo.Libro) (*modelo.Libro, error) {
	return r.repo.Create(libro)
}

func (r *LibroCasoUsoImpl) retrieveLibro(id string) (modelo.Libro, error) {
	return r.repo.Retrieve(id)
}

func (r *LibroCasoUsoImpl) updateLibro(id string, libro *modelo.Libro) (*modelo.Libro, error) {
	return r.repo.Update(id, libro)
}

func (r *LibroCasoUsoImpl) deleteLibro(id string) error {
	return r.repo.Delete(id)
}

func (r *LibroCasoUsoImpl) findAllLibro() ([]modelo.Libro, error) {
	return r.repo.FindAll()
}

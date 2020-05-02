package casousos

import (
	"github.com/gdosoftware/biblioteca/domain/interfaces"
	"github.com/gdosoftware/biblioteca/domain/modelo"
)

type LibroCasoUsoImpl struct {
	Repo interfaces.ILibroRepository
}

func CreateLibroCasoUsoImpl(repo interfaces.ILibroRepository) *LibroCasoUsoImpl{
	return &LibroCasoUsoImpl{Repo:repo}
}

func (r *LibroCasoUsoImpl) CreateLibro(libro *modelo.Libro) (*modelo.Libro, error) {
	return r.Repo.Create(libro)
}

func (r *LibroCasoUsoImpl) RetrieveLibro(id string) (modelo.Libro, error) {
	return r.Repo.Retrieve(id)
}

func (r *LibroCasoUsoImpl) UpdateLibro(id string, libro *modelo.Libro) (*modelo.Libro, error) {
	return r.Repo.Update(id, libro)
}

func (r *LibroCasoUsoImpl) DeleteLibro(id string) error {
	return r.Repo.Delete(id)
}

func (r *LibroCasoUsoImpl) FindAllLibro() ([]modelo.Libro, error) {
	return r.Repo.FindAll()
}

/*
var libro LibroCasoUsoImpl

libro = LibroCasoUsoImpl {Repo : LibroRepositoryImpl}
libro.CreateLIbro
	 .RetrieveLibro
	 .UpdateLibro
	 .DelteLibro
	 .FindAllLibro

func(l IlibroCasouso)

*/

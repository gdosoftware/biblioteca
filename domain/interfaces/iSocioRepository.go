package interfaces

import "github.com/gdosoftware/biblioteca/domain/modelo"

type ISocioRepository interface {
	Create(socio *modelo.Socio) (*modelo.Socio, error)
	Update(id string, socio *modelo.Socio) (*modelo.Socio, error)
	Retrieve(id string) (modelo.Socio, error)
	Delete(id string) error
	FindAll() ([]modelo.Socio, error)
}

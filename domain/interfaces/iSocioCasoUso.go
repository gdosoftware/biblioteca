package interfaces

import "github.com/gdosoftware/biblioteca/domain/modelo"

type ISocioCasoUso interface {
	CreateSocio(socio *modelo.Socio) (*modelo.Socio, error)
	RetrieveSocio(id string) (modelo.Socio, error)
	UpdateSocio(id string, socio *modelo.Socio) (*modelo.Socio, error)
	DeleteSocio(id string) error
	FindAllSocio() ([]modelo.Socio, error)
}

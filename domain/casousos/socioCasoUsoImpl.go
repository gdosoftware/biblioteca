package casousos

import (
	"github.com/gdosoftware/biblioteca/domain/interfaces"
	"github.com/gdosoftware/biblioteca/domain/modelo"
)

type SocioCasoUsoImpl struct {
	Repo interfaces.ISocioRepository
}

func CreateSocioCasoUsoImpl(repo interfaces.ISocioRepository) *SocioCasoUsoImpl{
	return &SocioCasoUsoImpl{Repo:repo}
}

func (r *SocioCasoUsoImpl) CreateSocio(socio *modelo.Socio) (*modelo.Socio, error) {
	return r.Repo.Create(socio)
}

func (r *SocioCasoUsoImpl) RetrieveSocio(id string) (modelo.Socio, error) {
	return r.Repo.Retrieve(id)
}

func (r *SocioCasoUsoImpl) UpdateSocio(id string, socio *modelo.Socio) (*modelo.Socio, error) {
	return r.Repo.Update(id, socio)
}

func (r *SocioCasoUsoImpl) DeleteSocio(id string) error {
	return r.Repo.Delete(id)
}

func (r *SocioCasoUsoImpl) FindAllSocio() ([]modelo.Socio, error) {
	return r.Repo.FindAll()
}

package casousos

import (
	"github.com/gdosoftware/biblioteca/domain/interfaces"
	"github.com/gdosoftware/biblioteca/domain/modelo"
)

type SocioCasoUsoImpl struct {
	repo *interfaces.ISocioRepository
}

func (r *SocioCasoUsoImpl) createSocio(socio *modelo.Socio) (*modelo.Socio, error) {
	return r.repo.Create(socio)
}

func (r *SocioCasoUsoImpl) retriveSocio(id string) (modelo.Socio, error) {
	return r.repo.Retrieve(id)
}

func (r *SocioCasoUsoImpl) updateSocio(id string, socio *modelo.Socio) (*modelo.Socio, error) {
	return r.repo.Update(id, socio)
}

func (r *SocioCasoUsoImpl) deleteSocio(id string) error {
	return r.repo.Delete(id)
}

func (r *SocioCasoUsoImpl) findAllLibro() ([]modelo.Socio, error) {
	return r.repo.FindAll()
}

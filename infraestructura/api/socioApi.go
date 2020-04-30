package api

import (
	"github.com/gdosoftware/biblioteca/domain/interfaces"
	"github.com/gdosoftware/biblioteca/domain/modelo"
)

type SocioApi struct {
	caso interfaces.ISocioCasoUso
}

func (l *SocioApi) altaSocio(socio *modelo.Socio) (*modelo.Socio, error) {
	return l.caso.CreateSocio(socio)
}

func (l *SocioApi) modificacionSocio(id string, socio *modelo.Socio) (*modelo.Socio, error) {
	return l.caso.UpdateSocio(id, socio)
}

func (l *SocioApi) borrarSocio(id string) error {
	return l.caso.DeleteSocio(id)
}

func (l *SocioApi) recuperarSocio(id string) (modelo.Socio, error) {
	return l.caso.RetrieveSocio(id)
}

func (l *SocioApi) recuperarTodosLosSocios() ([]modelo.Socio, error) {
	return l.caso.FindAllSocio()
}

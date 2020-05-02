package module

import (
	"github.com/gdosoftware/biblioteca/infraestructura/api"
	"github.com/gdosoftware/biblioteca/infraestructura/controllers"
	"github.com/gdosoftware/biblioteca/infraestructura"
	"github.com/gdosoftware/biblioteca/domain/interfaces"
	"github.com/gdosoftware/biblioteca/domain/casousos"
)

func MakeControllers(
	iLibroRepository interfaces.ILibroRepository,
	iSocioRepository interfaces.ISocioRepository,
	
) []server.Controller {
	
	libroController := controllers.CreateLibroController(createLibroApi(iLibroRepository))
	socioController := controllers.CreateSocioController(createSocioApi(iSocioRepository))

	return []server.Controller{libroController, socioController}
}

func createLibroApi(repo interfaces.ILibroRepository) *api.LibroApi{
   return api.CreateLibroApi(&casousos.LibroCasoUsoImpl{Repo:repo})
}

func createSocioApi(repo interfaces.ISocioRepository) *api.SocioApi{
	return api.CreateSocioApi(&casousos.SocioCasoUsoImpl{Repo:repo})
 }


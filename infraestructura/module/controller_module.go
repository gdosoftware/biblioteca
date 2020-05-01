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
	//iSocioRepository *ISocioRepository,
	//iPrestamoRepository *IPrestamoRepository
	
) []server.Controller {
	
	libroController := controllers.CreateLibroController(createLibroApi(iLibroRepository))
	//SocioController := controllers.CreateSocioController(createSocioApi(iSocioRepository))
	//prestamoController := controllers.CreatePrestamoController(createPrestamoApi(iPrestamoRepository))

	return []server.Controller{libroController}
}

func createLibroApi(repo interfaces.ILibroRepository) *api.LibroApi{
   return api.CreateLibroApi(&casousos.LibroCasoUsoImpl{Repo:repo})
}

/*func createSocioApi(repo *ISocioRepository) *SocioApi{
	return api.createSocioApi(&SocioCasoUsoImpl{repo:repo})
 }

 func createPrestamoApi(repo *IPrestamoRepository) *PrestamoApi{
	return api.createPrestamoApi(&PrestamoCasoUsoImpl{repo:repo})
 }*/


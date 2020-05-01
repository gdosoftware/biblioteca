package module

import (
	"github.com/gdosoftware/biblioteca/domain/casousos"
	"github.com/gdosoftware/biblioteca/domain/interfaces"
	"github.com/gdosoftware/biblioteca/infraestructura/api"
	"github.com/gdosoftware/biblioteca/infraestructura/controllers"
	"github.com/gdosoftware/biblioteca/infraestructura"

)

func MakeControllers(
	iLibroRepository *ILibroRepository,
	iSocioRepository *ISocioRepository,
	iPrestamoRepository *IPrestamoRepository
	
) []server.Controller {
	
	libroController := controllers.CreateLibroController(createLibroApi(iLibroRepository))
	SocioController := controllers.CreateSocioController(createSocioApi(iSocioRepository))
	prestamoController := controllers.CreatePrestamoController(createPrestamoApi(iPrestamoRepository))

	return []server.Controller{libroController, SocioController, prestamoController}
}

func createLibroApi(repo *ILibroRepository) *LibroApi{
   return api.createLibroApi(&LibroCasoUsoImpl{repo:repo})
}

func createSocioApi(repo *ISocioRepository) *SocioApi{
	return api.createSocioApi(&SocioCasoUsoImpl{repo:repo})
 }

 func createPrestamoApi(repo *IPrestamoRepository) *PrestamoApi{
	return api.createPrestamoApi(&PrestamoCasoUsoImpl{repo:repo})
 }


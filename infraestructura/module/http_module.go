package module

import (
	"github.com/gdosoftware/biblioteca/infraestructura/http"
	"github.com/gdosoftware/biblioteca/infraestructura/controllers"
	"github.com/gdosoftware/biblioteca/infraestructura"
	"github.com/gdosoftware/biblioteca/domain/interfaces"
	"github.com/gdosoftware/biblioteca/domain/casousos"
)

func MakeControllers(
	iLibroRepository interfaces.ILibroRepository,
	iSocioRepository interfaces.ISocioRepository,
	
) []server.Controller {
	
	libroController := controllers.CreateLibroController(createLibroHttp(iLibroRepository))
	socioController := controllers.CreateSocioController(createSocioHttp(iSocioRepository))

	return []server.Controller{libroController, socioController}
}

func createLibroHttp(repo interfaces.ILibroRepository) *http.LibroHttp{
   return http.CreateLibroHttp(casousos.CreateLibroCasoUsoImpl(repo))
}

func createSocioHttp(repo interfaces.ISocioRepository) *http.SocioHttp{
	return http.CreateSocioHttp(casousos.CreateSocioCasoUsoImpl(repo))
 }


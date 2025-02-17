package routes

import (
	"indodev/constans"
	"indodev/services"
	"indodev/services/divisionService"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// RoutesApi
func RoutesApi(e echo.Echo, usecaseSvc services.UsecaseService) {

	public := e.Group("")
	public.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: constans.TRUE_VALUE,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	divSvc := divisionService.NewDivisionService(usecaseSvc)

	divisionGroup := public.Group("/division")
	divisionGroup.POST("/add", divSvc.InsertDivision)
	divisionGroup.POST("/edit", divSvc.EditDivision)
	divisionGroup.POST("/delete", divSvc.DeleteDivision)
	divisionGroup.POST("/list", divSvc.ListDivision)
	divisionGroup.POST("/list-hirarki", divSvc.ListDivisionHiearchy)

}

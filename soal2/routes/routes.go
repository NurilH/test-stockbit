package routes

import (
	"net/http"
	"test_stockbit/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	e.GET("/movie/:imdbID", controllers.GetMovieDetailControllers)
	e.GET("/search/:searchword/:pagination", controllers.GetMovieControllers)

	return e
}

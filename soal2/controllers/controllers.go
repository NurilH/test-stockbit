package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"test_stockbit/models"
	response "test_stockbit/responses"

	"github.com/labstack/echo/v4"
)

func GetMovieControllers(c echo.Context) error {
	Search_Word := c.Param("searchword")
	Pagination := c.Param("pagination")
	api_url := "http://www.omdbapi.com/?apikey=faf7e5bb&s=" + Search_Word + "&page=" + Pagination

	res, _ := http.Get(api_url)

	responseData, _ := ioutil.ReadAll(res.Body)

	responseObject := models.Response{}
	json.Unmarshal(responseData, &responseObject)

	if responseObject.Movie == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", responseObject))
}

func GetMovieDetailControllers(c echo.Context) error {
	id := c.Param("imdbID")
	api_url := "http://www.omdbapi.com/?apikey=faf7e5bb&i=" + id

	res, _ := http.Get(api_url)

	responseData, _ := ioutil.ReadAll(res.Body)

	responseObject := models.MovieDetail{}
	json.Unmarshal(responseData, &responseObject)

	if responseObject.ImdbID == "" {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", responseObject))
}

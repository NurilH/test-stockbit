package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"test_stockbit/models"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEcho() *echo.Echo {
	e := echo.New()
	return e
}

type ResponseData struct {
	Message string
	Data    models.MovieDetail
}
type ResponseSearch struct {
	Message string
	Data    models.Response
}

func TestGetMovieDetailControllerSuccess(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "Success Operation",
		path:       "/movie/:imdbID",
		expectCode: http.StatusOK,
	}

	e := InitEcho()
	// InsertMockToDb()

	req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCases.path)
	context.SetParamNames("imdbID")
	context.SetParamValues("tt0869152")

	if assert.NoError(t, GetMovieDetailControllers(context)) {
		body := res.Body.String()
		var responses ResponseData
		err := json.Unmarshal([]byte(body), &responses)

		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	}

}
func TestGetMovieDetailControllerFail(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "Data Not Found",
		path:       "/movie/:imdbID",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()

	req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCases.path)
	context.SetParamNames("id_user")
	context.SetParamValues("1")

	if assert.NoError(t, GetMovieDetailControllers(context)) {
		body := res.Body.String()
		var responses ResponseData
		err := json.Unmarshal([]byte(body), &responses)

		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	}

}

func TestGetMovieControllerSuccess(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "Success Operation",
		path:       "/search/:searchword/1",
		expectCode: http.StatusOK,
	}

	e := InitEcho()

	req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCases.path)
	context.SetParamNames("searchword")
	context.SetParamValues("spiderman")

	if assert.NoError(t, GetMovieControllers(context)) {
		body := res.Body.String()
		var responses ResponseSearch
		err := json.Unmarshal([]byte(body), &responses)

		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	}

}

func TestGetMovieControllerFail(t *testing.T) {
	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{
		name:       "Data Not Found",
		path:       "/movie/:imdbID",
		expectCode: http.StatusBadRequest,
	}

	e := InitEcho()

	req := httptest.NewRequest(http.MethodGet, testCases.path, nil)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath(testCases.path)
	context.SetParamNames("id_user")
	context.SetParamValues("1")

	if assert.NoError(t, GetMovieControllers(context)) {
		body := res.Body.String()
		var responses ResponseSearch
		err := json.Unmarshal([]byte(body), &responses)

		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCases.expectCode, res.Code)
		assert.Equal(t, testCases.name, responses.Message)
	}

}

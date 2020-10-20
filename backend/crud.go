package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type CrudModel interface {
	Instance() interface{}
	DeleteRequestInstance() interface{}
	DeleteRequestKey(interface{}) interface{}
	Slice() interface{}
}

func ReadHandler(model CrudModel) func(c echo.Context) error {
	return func(c echo.Context) error {
		slice := model.Slice()
		if ListTable(&slice) != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, &slice)
	}
}

func CreateHandler(model CrudModel, logger echo.Logger) func(c echo.Context) error {
	return func(c echo.Context) error {
		instance := model.Instance()
		if err := c.Bind(&instance); err != nil {
			logger.Info(err)
			return c.NoContent(http.StatusBadRequest)
		}
		if err := c.Validate(instance); err != nil {
			logger.Info(err)
			return c.NoContent(http.StatusBadRequest)
		}
		return ParseSQLErrorToResponse(CreateRow(&instance), c)
	}
}

func DeleteHandler(model CrudModel, logger echo.Logger) func(c echo.Context) error {
	return func(c echo.Context) error {
		dinstance := model.DeleteRequestInstance()
		if err := c.Bind(&dinstance); err != nil {
			logger.Info(err)
			return c.NoContent(http.StatusBadRequest)
		}
		if err := c.Validate(dinstance); err != nil {
			logger.Info(err)
			return c.NoContent(http.StatusBadRequest)
		}
		instance := model.Instance()
		keys := model.DeleteRequestKey(dinstance)
		return ParseSQLErrorToResponse(DeleteRowByKeys(&instance, &keys), c)
	}
}

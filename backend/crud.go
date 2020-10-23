package main

import (
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"
)

/*
CrudModel is interface for CRUD generic handler functions.
When struct implement this interface, it can be interact database through generic handlers can treat.

- InstancePointer : Return empty struct instance's pointer.
- DeleteInstancePointer : Return empty delete request binding struct instance's pointer.
- DeleteKey : Return primary keys from given delete request struct instance's pointer.
- SlicePointer : Return empty struct slice instance's pointer.
*/
type CrudModel interface {
	InstancePointer() interface{}
	DeleteInstancePointer() interface{}
	DeleteKey(interface{}) interface{}
	SlicePointer() interface{}
}

func ReadHandler(model CrudModel) func(c echo.Context) error {
	return func(c echo.Context) error {
		slice := model.SlicePointer()
		c.Logger().Debug(reflect.TypeOf(slice).String())
		if ListTable(slice) != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, slice)
	}
}

func ReadAssocHandler(model CrudModel, assoc []string) func(c echo.Context) error {
	return func(c echo.Context) error {
		slice := model.SlicePointer()
		c.Logger().Debug(reflect.TypeOf(slice).String())
		if ListPreloadTable(slice, assoc) != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, slice)
	}
}

func CreateHandler(model CrudModel) func(c echo.Context) error {
	return func(c echo.Context) error {
		instance := model.InstancePointer()
		if err := c.Bind(instance); err != nil {
			c.Logger().Info(err)
			return c.NoContent(http.StatusBadRequest)
		}
		if err := c.Validate(instance); err != nil {
			c.Logger().Info(err)
			return c.NoContent(http.StatusBadRequest)
		}
		return ParseSQLErrorToResponse(CreateRow(instance), c)
	}
}

func DeleteHandler(model CrudModel) func(c echo.Context) error {
	return func(c echo.Context) error {
		dinstance := model.DeleteInstancePointer()
		if err := c.Bind(dinstance); err != nil {
			c.Logger().Info(err)
			return c.NoContent(http.StatusBadRequest)
		}
		if err := c.Validate(dinstance); err != nil {
			c.Logger().Info(err)
			return c.NoContent(http.StatusBadRequest)
		}
		instance := model.InstancePointer()
		keys := model.DeleteKey(dinstance)
		return ParseSQLErrorToResponse(DeleteRowByKeys(instance, keys), c)
	}
}

func UpdateHandler(model CrudModel) func(c echo.Context) error {
	return func(c echo.Context) error {
		instance := model.InstancePointer()
		if err := c.Bind(instance); err != nil {
			c.Logger().Info(err)
			return c.NoContent(http.StatusBadRequest)
		}
		if err := c.Validate(instance); err != nil {
			c.Logger().Info(err)
			return c.NoContent(http.StatusBadRequest)
		}
		return ParseSQLErrorToResponse(UpdateRow(instance), c)
	}
}

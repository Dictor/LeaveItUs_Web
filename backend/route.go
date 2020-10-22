package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// AttachRoutes attach route handlers to echo instance
func AttachRoutes(e *echo.Echo) {
	// endpoint "tag" has CRD actions
	e.GET("/api/tag", ReadHandler(Tag{}))
	e.POST("/api/tag", CreateHandler(Tag{}))
	e.DELETE("/api/tag", DeleteHandler(Tag{}))

	// endpoint "person" has CRUD actions
	e.GET("/api/person", ReadHandler(Person{}))
	e.POST("/api/person", CreateHandler(Person{}))
	e.DELETE("/api/person", DeleteHandler(Person{}))
	e.PUT("/api/person", UpdateHandler(Person{}))

	// endpoint for locker hardware
	e.GET("/api/timestamp", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]int64{"timestamp": time.Now().Unix()})
	})
}

func readLockerTag(c echo.Context) error {
	uid := c.Param("uid")
	v := validator.New()
	if err := v.Var(uid, "required,printascii"); err != nil {
		c.Logger().Info(err)
		return c.NoContent(http.StatusBadRequest)
	}

	locker := Locker{}
	if err := SelectTable(&locker, uid); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.NoContent(http.StatusNotFound)
		} else {
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	return c.JSON(http.StatusOK, locker.Tags)
}

func createLockerTagRecord(c echo.Context) error {

}

func createLockerDoorEvent(c echo.Context) error {

}

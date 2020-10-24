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

	// endpoint "locker" has CRUD actions
	e.GET("/api/locker", ReadAssocHandler(Locker{}, []string{"Tags"}))
	e.POST("/api/locker", CreateHandler(Locker{}))
	e.DELETE("/api/locker", DeleteHandler(Locker{}))
	e.PUT("/api/locker", UpdateHandler(Locker{}))

	// endpoint for locker hardware
	e.GET("/api/timestamp", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]int64{"timestamp": time.Now().Unix()})
	})
	e.GET("/api/locker/:uid/tag", readLockerTag)
	e.POST("/api/locker/:uid/tag", createLockerTagRecord)
	e.POST("/api/locker/:uid/door", createLockerDoorEvent)
}

func readLockerTag(c echo.Context) error {
	uid := c.Param("uid")
	v := validator.New()
	if err := v.Var(uid, "required,printascii"); err != nil {
		c.Logger().Info(err)
		return c.NoContent(http.StatusBadRequest)
	}

	locker := Locker{}
	if err := SelectPreloadTable(&locker, []string{"Tags"}, "UID = ?", uid); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.NoContent(http.StatusNotFound)
		} else {
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	return c.JSON(http.StatusOK, locker.Tags)
}

func createLockerTagRecord(c echo.Context) error {
	uid := c.Param("uid")
	v := validator.New()
	if err := v.Var(uid, "required,printascii"); err != nil {
		c.Logger().Info(err)
		return c.NoContent(http.StatusBadRequest)
	}

	req := LockerRecordCreateRequest{}
	if err := c.Bind(&req); err != nil {
		c.Logger().Info(err)
		return c.NoContent(http.StatusBadRequest)
	}
	if err := c.Validate(&req); err != nil {
		c.Logger().Info(err)
		return c.NoContent(http.StatusBadRequest)
	}

	row := LockerRecord{
		TagUIDs:   req.TagUIDs,
		Weight:    req.Weight,
		LockerUID: uid,
	}
	return ParseSQLErrorToResponse(CreateRow(row), c)
}

func createLockerDoorEvent(c echo.Context) error {
	uid := c.Param("uid")
	v := validator.New()
	if err := v.Var(uid, "required,printascii"); err != nil {
		c.Logger().Info(err)
		return c.NoContent(http.StatusBadRequest)
	}

	req := LockerDoorEventCreateRequest{}
	if err := c.Bind(&req); err != nil {
		c.Logger().Info(err)
		return c.NoContent(http.StatusBadRequest)
	}
	if err := c.Validate(&req); err != nil {
		c.Logger().Info(err)
		return c.NoContent(http.StatusBadRequest)
	}

	row := LockerDoorEvent{
		ClosedTime: req.ClosedTime,
		Duration:   req.Duration,
		LockerUID:  uid,
	}
	return ParseSQLErrorToResponse(CreateRow(row), c)
}

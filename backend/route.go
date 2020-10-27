package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
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

	e.GET("/api/locker/door", readDoorEventWithLimit)
	e.GET("/api/locker/tag", ReadHandler(LockerRecord{}))

	// endpoint for locker hardware
	e.GET("/api/timestamp", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]int64{"timestamp": time.Now().Unix()})
	})
	e.GET("/api/locker/:uid/tag", readAllLockerTag)
	e.GET("/api/locker/:uid/door", readLockerDoorEvent)
	e.POST("/api/locker/:uid/tag", createLockerTagRecord)
	e.POST("/api/locker/:uid/door", createLockerDoorEvent)
}

func readAllLockerTag(c echo.Context) error {
	if c.Param("uid") != "latest" {
		return readLockerTagRecord(c)
	}

	lockers := []Locker{}
	if err := ListTable(&lockers); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	tx := GetDatabase()
	lrs := map[string]*LockerRecord{}
	for _, l := range lockers {
		lr := LockerRecord{}
		tx.Where("locker_uid = ?", l.UID).Order("created_at desc").Limit(1).Find(&lr)
		lrs[l.UID] = &lr
	}
	return c.JSON(http.StatusOK, lrs)
}

func readDoorEventWithLimit(c echo.Context) error {
	slimit := c.QueryParam("limit")
	limit, err := strconv.Atoi(slimit)
	de := []LockerDoorEvent{}

	if err != nil {
		ListTable(&de)
	} else {

		timelimit := int(time.Now().Unix()) - limit
		c.Logger().Debugf("timelimit=%d", timelimit)
		SelectTable(&de, "closed_time > ?", timelimit)
	}

	return c.JSON(http.StatusOK, de)
}

func readLockerTagRecord(c echo.Context) error {
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

func readLockerDoorEvent(c echo.Context) error {
	uid := c.Param("uid")
	v := validator.New()
	if err := v.Var(uid, "required,printascii"); err != nil {
		c.Logger().Info(err)
		return c.NoContent(http.StatusBadRequest)
	}

	de := []LockerDoorEvent{}
	if err := SelectTable(&de, "locker_uid = ?", uid); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.NoContent(http.StatusNotFound)
		} else {
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	return c.JSON(http.StatusOK, de)
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

	uidser, err := json.Marshal(req.TagUIDs)
	if err != nil {
		c.Logger().Info(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	row := LockerRecord{
		TagUIDs:   string(uidser),
		Weight:    req.Weight,
		LockerUID: uid,
	}
	return ParseSQLErrorToResponse(CreateRow(&row), c)
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
	return ParseSQLErrorToResponse(CreateRow(&row), c)
}

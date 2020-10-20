package main

import "github.com/labstack/echo/v4"

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
}

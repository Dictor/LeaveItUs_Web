package main

import (
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

//VueRouterStatic is static file serving handler for vue-router
func VueRouterStatic(root string) func(c echo.Context) error {
	return func(c echo.Context) error {
		u := c.Request().URL
		path := filepath.Clean(u.EscapedPath())
		var servePath string

		if fileExists(root + path) {
			servePath = root + path
		} else {
			servePath = root + "/index.html"
		}
		c.Logger().Debugf("Serving static resource. request=%s cleaned_path=%s file_path=%s", u, path, servePath)
		return c.File(servePath)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

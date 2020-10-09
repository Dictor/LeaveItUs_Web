package main

import (
	elogrus "github.com/dictor/echologrus"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	elogrus.Attach(e)

	e.Static("/", "static")

	e.Logger.Fatal(e.Start(":80"))
}

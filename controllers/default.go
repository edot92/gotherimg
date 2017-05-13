package controllers

import (
	"github.com/labstack/echo"
)

//create user

func RenderIndex(c echo.Context) error {

	// var err error
	return c.Render(200, "views/index.html", "")
}

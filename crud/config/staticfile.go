package config

import (
	"github.com/labstack/echo"
)

func SetUpStaticFile(e *echo.Echo) {

	e.Static("/lib", "./static/lib")
}

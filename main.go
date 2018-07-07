package main

import (
	"github.com/gloryof/go-crud-practice/crud/config"
	"github.com/gloryof/go-crud-practice/crud/config/router"
	"github.com/gloryof/go-crud-practice/crud/externals"
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	e.Renderer = config.CreateTemplate()
	ctx := externals.CreateContext()
	defer ctx.Close()

	ug := e.Group("/user")
	router.User(ctx, ug)

	e.Start(":8000")
}

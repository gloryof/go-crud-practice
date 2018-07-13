package main

import (
	"fmt"

	"github.com/gloryof/go-crud-practice/crud/config"
	"github.com/gloryof/go-crud-practice/crud/config/router"
	"github.com/gloryof/go-crud-practice/crud/externals"
	"github.com/labstack/echo"
)

func main() {

	e := createEcho()

	ctx := createExternalsContext()
	defer ctx.Close()

	route(e, ctx)

	start(e)
}

func createEcho() *echo.Echo {

	e := echo.New()
	config.SetUpTemplate(e)
	config.SetUpStaticFile(e)

	return e
}

func createExternalsContext() externals.Context {

	ctx, err := externals.CreateContext()

	if err != nil {

		fmt.Printf("Error!: %s", err)
	}

	return ctx
}

func route(e *echo.Echo, ctx externals.Context) {

	ug := e.Group("/user")
	router.User(ctx, ug)

}

func start(e *echo.Echo) {

	e.Start(":8000")
}

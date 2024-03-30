package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
)

type Templates struct {
	templates *template.Template
}

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
}

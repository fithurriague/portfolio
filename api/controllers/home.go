package controllers

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Home godoc
//
//	@Summary		Home page
//	@Description	Builds and returns the home page file
//	@Tags			controller
//	@Produce		html
//
//	@Success		200	{string}	string		"Returns the home html page"
//	@Failure		500	{string}	HTTPError	"Internal server error"
//
//	@Router			/ [get]
func Home(templates map[string]*template.Template) func(c echo.Context) error {
	return func(c echo.Context) error {
		// This works because a map works with a pointer under the hood
		templates["home.html"] = template.Must(template.ParseFiles("dist/views/home.html", "dist/layouts/base.html"))

		return c.Render(http.StatusOK, "home.html", map[string]any{
			"name": "Ithurriague Francisco",
			"rol":  "Backend Developer | Golang",
		})
	}
}

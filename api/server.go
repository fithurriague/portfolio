package api

import (
	"errors"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/fithurriague/portfolio/api/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"golang.org/x/time/rate"

	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/fithurriague/portfolio/docs"
)

type TemplateRegistry struct {
	templates map[string]*template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data any, c echo.Context) error {
	tmpl, ok := t.templates[name]

	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}

	return tmpl.ExecuteTemplate(w, "base.html", data)
}

type Server struct {
	Router *echo.Echo
	Logger zerolog.Logger
}

func NewServer() (server *Server, err error) {
	server = &Server{
		Logger: zerolog.New(os.Stdout).With().Timestamp().Logger(),
	}

	server.Router = echo.New()
	server.Router.Use(middleware.CORS())
	server.Router.Use(middleware.BodyLimit("2M"))
	server.Router.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(1000))))
	server.Router.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			server.Logger.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))

	err = server.loadRoutes()
	if err != nil {
		return nil, err
	}

	return server, nil
}

func (server *Server) loadRoutes() (err error) {
	server.Router.GET("/swagger/*", echoSwagger.WrapHandler)

	server.Router.Static("/css", "public/assets/css")

	server.Router.GET("/source/:filename", func(c echo.Context) error {
		fileName := c.Param("filename")

		data, err := os.ReadFile(fileName)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "File not found")
		}

		c.Response().Header().Set("Content-Type", "text/plain")
		return c.String(http.StatusOK, string(data))
	})

	templates := make(map[string]*template.Template)

	server.Router.GET("/", controllers.Home(templates))

	server.Router.Renderer = &TemplateRegistry{
		templates: templates,
	}

	return nil
}

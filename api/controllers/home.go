package controllers

import (
	"html/template"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	URL           string
	ImgURL        string
	Name          string
	Company       string
	IsProprietary bool
	Role          string
	Description   string
	Stack         map[string]string
	StartDate     string
	EndDate       string
}

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
		templates["home.html"] = template.Must(template.ParseFiles("dist/components/article.html", "dist/components/navbar.html", "dist/views/home.html", "dist/layouts/base.html"))

		return c.Render(http.StatusOK, "home.html", map[string]any{
			"name": "Ithurriague Francisco",
			"role": "Golang Backend Developer",
			"projects": []Project{
				{
					URL:           "/project/portfolio",
					ImgURL:        "",
					Name:          "Portfolio",
					Company:       "-",
					IsProprietary: false,
					Role:          "Owner",
					Description:   "My personal portfolio. I wanted to showcast my expertise with Go, and I think websites should be fast and be clear rather than be heavy on animations",
					Stack: map[string]string{
						"#golang":     "https://go.dev/",
						"#echo":       "https://echo.labstack.com/",
						"#docker":     "https://www.docker.com/",
						"#javascript": "https://developer.mozilla.org/es/docs/Web/JavaScript",
						"#vitejs":     "https://vite.dev/",
						"#css":        "https://developer.mozilla.org/es/docs/Web/CSS",
						"#htmx":       "https://htmx.org/",
					},
					StartDate: "03/24/25",
					EndDate:   time.Now().Format("01/02/06"),
				},

				{
					URL:           "/project/streaming",
					ImgURL:        "/dist/assets/images/streaming.png",
					Name:          "Streaming Platform",
					Company:       "ThirstyOasis",
					IsProprietary: true,
					Role:          "Golang Backend Developer",
					Description:   "One to one streaming and broadcast streaming. Here I am in charge of maintaining the platform and doing the whole architecture and desing of the new features as well as being the main contributor to the backend. I also collaborate on a daily basis with my team to meet the weekly deadlines.",
					Stack: map[string]string{
						"#golang":   "https://go.dev/",
						"#echo":     "https://echo.labstack.com/",
						"#docker":   "https://www.docker.com/",
						"#redis":    "https://redis.io/es/",
						"#rabbitmq": "https://www.rabbitmq.com/",
						"#mongodb":  "https://www.mongodb.com/",
					},
					StartDate: "10/25/24",
					EndDate:   time.Now().Format("01/02/06"),
				},

				{
					URL:           "/project/transcriptor",
					ImgURL:        "/dist/assets/images/transcriptor.png",
					Name:          "AI Audio Transcriptor",
					Company:       "Kakeads",
					IsProprietary: true,
					Role:          "Fullstack developer",
					Description:   "I developed a whole transcription application which allows to enqueue audio files and transcribe them to text files using Google's AI APIs.",
					Stack: map[string]string{
						"#golang":       "https://go.dev/",
						"#echo":         "https://echo.labstack.com/",
						"#docker":       "https://www.docker.com/",
						"#google-cloud": "https://cloud.google.com/?hl=es",
						"#nginx":        "https://nginx.org/",
						"#flutter":      "https://flutter.dev/",
					},
					StartDate: "11/30/23",
					EndDate:   "12/16/23",
				},
			},
		})
	}
}

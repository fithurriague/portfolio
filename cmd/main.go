package main

import (
	"log"

	"github.com/fithurriague/portfolio/api"
	"github.com/joho/godotenv"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	Personal portfolio
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Ithurriague Francisco
//	@contact.url	https://ithurriague.francisco.com/support
//	@contact.email	francisco@ithurriague.com

//	@host	localhost:8000

//	@securitydefinitions.apikey	Authentication
//	@in							header
//	@name						Bearer

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("loading env variables: %s", err.Error())
	}

	server, err := api.NewServer()
	if err != nil {
		log.Fatalf("error setting up the server: %s", err.Error())
	}

	err = server.Router.Start(":8000")
	if err != nil {
		log.Fatalf("couldn't start server: %s", err.Error())
	}
}

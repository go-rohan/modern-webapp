package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rshukla9/modern-web-app/pkg/config"
	"github.com/rshukla9/modern-web-app/pkg/handlers"
	"github.com/rshukla9/modern-web-app/pkg/render"
)

var portNumber = ":8080"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Booting application on port %s", portNumber)

	_ = http.ListenAndServe(portNumber, nil)

}

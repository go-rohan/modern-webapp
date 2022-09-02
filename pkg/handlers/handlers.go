package handlers

import (
	"net/http"

	"github.com/go-rohan/modern-webapp/pkg/config"
	"github.com/go-rohan/modern-webapp/pkg/models"
	"github.com/go-rohan/modern-webapp/pkg/render"
)

// holds the data sent from handlers and send it to the template

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)

	// m.App.Session

	stringMap["test"] = "Hello again"
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}

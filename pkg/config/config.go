package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

//global app config

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	AppLog        *log.Logger
	AppVersion    string
	InProd        bool
	Session       *scs.SessionManager
}

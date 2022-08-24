package config

import (
	"html/template"
	"log"
)

//global app config

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	AppLog        *log.Logger
	AppVersion    string
}

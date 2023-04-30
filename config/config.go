package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
	"github.com/bertoxic/aarc/models"
)

type AppConfig struct {
	UserCache     bool
	MailChan chan models.MailData
	TemplateCache map[string]*template.Template
	InProduction  bool
	Sessions *scs.SessionManager
}      
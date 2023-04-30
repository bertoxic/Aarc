package models

import "github.com/bertoxic/aarc/forms"

type TemplateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	Floatmap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Flash           string
	Warning         string
	Error           string
	Form            *forms.Form
	IsAuthenticated int
}
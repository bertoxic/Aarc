package render

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/bertoxic/aarc/config"
	"github.com/bertoxic/aarc/models"
)

var functions = template.FuncMap {
    
}

var app *config.AppConfig
//var pathToTemplate = "C:/Users/HP/Desktop/Bert/templates"
var pathToTemplate = "templates"

func NewRenderer(a *config.AppConfig) {
	app = a
}
func AddDefaultData (td *models.TemplateData, r *http.Request)*models.TemplateData{
    //td.CSRFToken = nosurf.Token(r)
	td.Flash = app.Sessions.PopString(r.Context(), "flash")
	td.Error = app.Sessions.PopString(r.Context(), "error")
	td.Warning = app.Sessions.PopString(r.Context(), "warning")
	// if app.Sessions.Exists(r.Context(),"user_id") {
	// td.IsAuthenticated = 1
	// }

	return td
}





func Template (w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) error {
    var tc map[string]*template.Template

    if app.UserCache {
        tc = app.TemplateCache
    } else {
        tc, _ = CreateTemplateCache()
    }
    t, ok := tc[tmpl]
    if !ok {
        return errors.New("could not get template from cache")
    }
    buf := new(bytes.Buffer)
	//td = AddDefaultData(td, r)
    //log.Println("zzzzzz page template",tmpl)
    
	t.Execute(buf, td)
	_, err := buf.WriteTo(w)
	if err != nil {
		return err

}
return nil
}





func CreateTemplateCache() (map[string]*template.Template, error) {
    myCache := map[string]*template.Template{}

    pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.html", pathToTemplate )) 
    if err != nil{
        	return myCache, err

	}
       for  _, page := range pages {
            name := filepath.Base(page)
            ts , err := template.New(name).Funcs(functions).ParseFiles(page)
           log.Println("page found "+name + fmt.Sprintf("%d",len(pages)))
           if err != nil {
            return myCache, err
           }
        matches, err := filepath.Glob(fmt.Sprintf("%s/*layout.html", pathToTemplate))
        if err != nil {
            return myCache, err
        }
        if len(matches) > 0 {
            ts, err = ts.ParseGlob(fmt.Sprintf("%s/*layout.html",pathToTemplate))
            if err != nil {
				return myCache, err
			}
        }
         myCache[name] = ts
       }
      
    return  myCache , nil
}
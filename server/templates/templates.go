package templates

import (
   "log"
   "net/http"
   "html/template"
)

type Login struct {
   BAlertUser bool
   AlertMsg   string
}

type Register struct {
   BAlertUser bool
   AlertMsg   string
}

type Dashboard struct {
   BAlertUser bool
   AlertMsg   string
}

var template = template.Must(template.ParseFiles(
   "./server/templates/templateFiles/login.tmpl",
   "./server/templates/templateFiles/register.tmpl",
   "./server/templates/templateFiles/restricted.tmpl"
))

func RenderTemplate(res http.ResponseWriter, tmpl string, p interface{}) {
   err := templates.ExecuteTemplate(res, tmpl + ".tmpl", p)
   if err != nil {
      log.Printf("You've got template err %v", err)
      http.Error(res, err.Error(), http.StatusInternalServerError)
   }
}
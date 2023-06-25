package main

import (
	"html/template"
	"os"

	log "github.com/sirupsen/logrus"
)

var raw = `
/* generated by cmd/genswag/main.go */
package main

//	@title			Jepp API Documentation
//	@description	This is a simple api to access jeopardy data.
//	@version		1.0
//  @host			{{.Host}}
//	@BasePath		/api
//	@schemes		{{.Scheme}}
//	@contact.name	shreve
//	@license.name	MIT
//	@license.url	https://github.com/ecshreve/jepp/blob/main/LICENSE
`

func main() {
	tmpl := template.Must(template.New("swag").Parse(raw))

	f, err := os.Create("./cmd/server/swag.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	tmplArgs := struct {
		Host   string
		Scheme string
	}{
		Host:   os.Getenv("API_HOST"),
		Scheme: "https",
	}

	if os.Getenv("JEPP_LOCAL_DEV") == "true" {
		tmplArgs.Scheme += " http"
	}

	if err := tmpl.Execute(f, tmplArgs); err != nil {
		log.Fatal(err)
	}
}
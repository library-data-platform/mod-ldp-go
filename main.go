package main

import (
	"github.com/folio-org/mod-ldp/app"
	"github.com/folio-org/mod-ldp/app/config"
)

// Using this project as a template
// http://www.golangprograms.com/advance-programs/golang-restful-api-using-grom-and-gorilla-mux.html

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":8001")
}

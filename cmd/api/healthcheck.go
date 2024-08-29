package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Status: OK")
	fmt.Println(w, "environment: %s\n", app.config.env)
	fmt.Println(w, "version: %s\n", version)
}
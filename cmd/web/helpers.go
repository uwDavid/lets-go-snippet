package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// serverError helper writes error + stack trace to errorLog using debug.Stack()
// then sends 500 Internal Server Error to user
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

	app.errorLog.Output(2, trace) //need to set frame depth to 2

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// send status code + description to user
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// wrap around clientError, send 404 to user
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

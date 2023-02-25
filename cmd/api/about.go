package main

import "net/http"

func (app *application) aboutHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"company":  "Go Developer",
		"version":  "0.0.1",
		"golang":   "go1.20",
		"power-by": "Open Source Software",
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

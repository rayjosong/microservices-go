package main

import (
	"log-service/data"
	"net/http"

	"github.com/tsawler/toolbox"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	var tools toolbox.Tools

	// read json into var
	var requestPayload JSONPayload
	_ = tools.ReadJSON(w, r, &requestPayload)

	// insert data
	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		tools.ErrorJSON(w, err)
		return
	}

	resp := toolbox.JSONResponse{
		Error: false,
		Message: "logged",
	}

	tools.WriteJSON(w, http.StatusAccepted, resp)
}

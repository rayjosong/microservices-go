package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/tsawler/toolbox"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var tools toolbox.Tools
	var requestPayload struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	// read JSON into requestPayLoad struct
	err := tools.ReadJSON(w, r, &requestPayload)
	if err != nil {
			tools.ErrorJSON(w, err, http.StatusBadRequest)
	}

	// validate the user against the database
	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		tools.ErrorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		tools.ErrorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	payload := toolbox.JSONResponse{
		Error: false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data: user,
	}

	tools.WriteJSON(w, http.StatusAccepted, payload)
}

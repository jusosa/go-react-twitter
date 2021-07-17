package routers

import (
	"encoding/json"
	"github.com/jusosa/go-react-twitter/bd"
	"github.com/jusosa/go-react-twitter/models"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error verify the payload "+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "E-mail is required", 400)
		return
	}
	if len(user.Password) < 6 {
		http.Error(w, "Minimum password length is 6 ", 400)
		return
	}

	_, founded, _ := bd.UserExists(user.Email)
	if founded {
		http.Error(w, "E-Mail already exists", 400)
		return
	}

	_, status, err := bd.CreateUser(user)
	if err != nil {
		http.Error(w, "Error creating user "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Error creating user (wrong status) "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

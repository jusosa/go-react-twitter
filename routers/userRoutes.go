package routers

import (
	"encoding/json"
	"github.com/jusosa/go-react-twitter/bd"
	"github.com/jusosa/go-react-twitter/jwt"
	"github.com/jusosa/go-react-twitter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
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

	user.ID = primitive.NewObjectID()
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

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid Credentials "+err.Error(), 401)
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

	userLogged, logged := bd.TryLogin(user.Email, user.Password)

	if !logged {
		http.Error(w, "Invalid Credentials ", 401)
		return
	}

	jwtKey, err := jwt.GenerateJWT(userLogged)
	if err != nil {
		http.Error(w, "Token error: "+err.Error(), 401)
		return
	}

	var resp = models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "Token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}

func ViewProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Missed id param", http.StatusBadRequest)
		return
	}

	profile, err := bd.ViewProfile(ID)
	if err != nil {
		http.Error(w, "Error retrieving user info ["+err.Error()+"]", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)
}

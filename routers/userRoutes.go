package routers

import (
	"encoding/json"
	"github.com/jusosa/go-react-twitter/bd"
	"github.com/jusosa/go-react-twitter/jwt"
	"github.com/jusosa/go-react-twitter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
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

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var status bool

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Update error: "+err.Error(), http.StatusBadRequest)
		return
	}

	status, err = bd.UpdateUser(user, UserId)
	if err != nil {
		http.Error(w, "Update Error. Try again: "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "Update Error. Try again:", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatars")
	var extension = strings.Split(handler.Filename, ".")[1]
	var fileName string = "uploads/avatars/" + UserId + "." + extension

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error on upload: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error saving the file: "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Avatar = UserId + "." + extension
	status, err = bd.UpdateUser(user, UserId)
	if err != nil || !status {
		http.Error(w, "Error saving the avatar: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func UploadBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banners")
	var extension = strings.Split(handler.Filename, ".")[1]
	var fileName string = "uploads/banners/" + UserId + "." + extension

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error on upload: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error saving the file: "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Banner = UserId + "." + extension
	status, err = bd.UpdateUser(user, UserId)
	if err != nil || !status {
		http.Error(w, "Error saving the banner: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func GetAvatar(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var openFile *os.File
	if len(ID) < 1 {
		http.Error(w, "Id param is required", http.StatusBadRequest)
		return
	}

	profile, err := bd.ViewProfile(ID)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	openFile, err = os.Open("uploads/avatars/"+profile.Avatar)
	if err != nil {
		http.Error(w, "Source not found", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, openFile)
	if err != nil {
		http.Error(w, "Could not load the source", http.StatusBadRequest)
		return
	}
}

func GetBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var openFile *os.File
	if len(ID) < 1 {
		http.Error(w, "Id param is required", http.StatusBadRequest)
		return
	}

	profile, err := bd.ViewProfile(ID)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	openFile, err = os.Open("uploads/banners/"+profile.Banner)
	if err != nil {
		http.Error(w, "Source not found", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, openFile)
	if err != nil {
		http.Error(w, "Could not load the source", http.StatusBadRequest)
		return
	}
}

func GetUsersAll(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	search := r.URL.Query().Get("search")

	if len(typeUser) < 1 {
		http.Error(w, "Missed type param", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Missed page param", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil{
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}


	users, status := bd.FindAllUsers(UserId, int64(page), search, typeUser)
	if !status{
		http.Error(w, "No Data Found", http.StatusNotFound)
		return
	}

	w.Header().Set("content-type","application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(users)
}

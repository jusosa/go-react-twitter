package handlers

import (
	"github.com/gorilla/mux"
	"github.com/jusosa/go-react-twitter/middlew"
	"github.com/jusosa/go-react-twitter/routers"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func Handle() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.DBCheck(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.DBCheck(routers.Login)).Methods("POST")
	router.HandleFunc("/profile", middlew.DBCheck(middlew.Validate(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/updateProfile", middlew.DBCheck(middlew.Validate(routers.UpdateUser))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.DBCheck(middlew.Validate(routers.CreateTweet))).Methods("POST")
	router.HandleFunc("/tweet", middlew.DBCheck(middlew.Validate(routers.GetTweets))).Methods("GET")
	router.HandleFunc("/tweet", middlew.DBCheck(middlew.Validate(routers.DeleteTweet))).Methods("DELETE")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8081"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

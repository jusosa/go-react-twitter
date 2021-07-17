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

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8081"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
package routers

import (
	"encoding/json"
	"github.com/jusosa/go-react-twitter/bd"
	"github.com/jusosa/go-react-twitter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
	"time"
)

func CreateTweet(w http.ResponseWriter, r *http.Request) {
	var message models.TweetBody
	var status bool
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Incoming data error", http.StatusBadRequest)
		return
	}

	tweet := models.Tweet{
		ID:           primitive.NewObjectID(),
		UserID:       UserId,
		Message:      message.Message,
		CreationDate: time.Now(),
	}

	_, status, err = bd.CreateTweet(tweet)
	if err != nil {
		http.Error(w, "Error creating tweet: "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "IError creating tweet", http.StatusBadRequest)
		return
	}
}

func GetTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Missed id param", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Missed page param", http.StatusBadRequest)
		return
	}

	pageNumber, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil{
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}


	tweets, status := bd.FindTweets(ID, int64(pageNumber))
	if !status{
		http.Error(w, "No Data Found", http.StatusNotFound)
		return
	}

	w.Header().Set("content-type","application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(tweets)
}

func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Missed id param", http.StatusBadRequest)
		return
	}

	err := bd.DeleteTweet(ID, UserId)
	if err != nil{
		http.Error(w, "Unable to delete tweet", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type","application/json")
	w.WriteHeader(http.StatusOK)
}
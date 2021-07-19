package routers

import (
	"encoding/json"
	"github.com/jusosa/go-react-twitter/bd"
	"github.com/jusosa/go-react-twitter/models"
	"net/http"
	"strconv"
)

func CreateRelation (w http.ResponseWriter, r *http.Request){
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Missed id param", http.StatusBadRequest)
		return
	}

	var relation models.Relation
	relation.UserId = UserId
	relation.FollowingUser = ID

	_, status, err := bd.CreateRelation(relation)

	if err != nil {
		http.Error(w, "Could not crete the relation: "+err.Error(), http.StatusExpectationFailed)
		return
	}

	if !status {
		http.Error(w, "Could not crete the relation: ", http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func DeleteRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Missed id param", http.StatusBadRequest)
		return
	}

	var relation models.Relation
	relation.UserId = UserId
	relation.FollowingUser = ID

	err := bd.DeleteRelation(relation)
	if err != nil{
		http.Error(w, "Unable to delete relation", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type","application/json")
	w.WriteHeader(http.StatusOK)
}

func FindRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Missed id param", http.StatusBadRequest)
		return
	}

	var relation models.Relation
	relation.UserId = UserId
	relation.FollowingUser = ID

	status, err := bd.FindRelation(relation)
	if err != nil{
		http.Error(w, "Relation not found: "+err.Error(), http.StatusNotFound)
		return
	}

	if !status{
		http.Error(w, "Relation not found", http.StatusNotFound)
		return
	}

	w.Header().Set("content-type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.RelationResponse{
		Status: status,
	})
}

func ReadTweets(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Missed page param", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil{
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}

	tweets, status := bd.FindTweetsByUser(UserId, page)
	if !status{
		http.Error(w, "No Data Found", http.StatusNotFound)
		return
	}

	w.Header().Set("content-type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tweets)
}
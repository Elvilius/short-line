package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"github.com/Elvilius/short-line/db"
	"github.com/gorilla/mux"
)

var repository db.Db

func main() {
	repository = db.Connect(os.Getenv("PSQL_URL"))
	fmt.Println("Server listening!")
	http.ListenAndServe(":5656", initRoutes())
}

func CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	type UrlDto struct {
		Url string `json:"url"`
	}
	var urlDto UrlDto

	err := json.NewDecoder(r.Body).Decode(&urlDto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	shortUrl := urlDto.Url
	json.NewEncoder(w).Encode(map[string]string{"data": shortUrl})
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	repository.GetUrlByFullAddres("asdasdasd")
	http.Redirect(w, r, "https://github.com/Elvilius", http.StatusMovedPermanently)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to this life-changing API.")
}

func initRoutes() *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/", Index)
	route.HandleFunc("/{key}", Redirect)
	route.HandleFunc("/create", CreateShortUrl).Methods("POST")
	return route
}

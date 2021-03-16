package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Elvilius/short-line/db"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strconv"
)

var repository db.Db

func main() {
	repository = db.Connect(os.Getenv("PSQL_URL"))
	defer repository.Conn.Close(context.Background())
	fmt.Println("Server listening!")
	http.ListenAndServe(":5656", initRoutes())
}

func CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	var id int

	type UrlDto struct {
		Url string `json:"url"`
	}
	var urlDto UrlDto

	err := json.NewDecoder(r.Body).Decode(&urlDto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	existUrl, err := repository.GetUrlByFullAddres(urlDto.Url)

	if err != nil {
		id = repository.CreateUrl(urlDto.Url)
	} else {
		id = existUrl.Id
	}

	shortUrl := os.Getenv("HOST_URL") + "/" + strconv.Itoa(id)	
	json.NewEncoder(w).Encode(map[string]string{"data": shortUrl})

}

func Redirect(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["key"]
	url := repository.GetUrlById(id)
	http.Redirect(w, r, url.Full_address_name, http.StatusMovedPermanently)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to this life-changing API.")
}

func initRoutes() *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/", Index)
	route.HandleFunc("/create", CreateShortUrl).Methods("POST")
	route.HandleFunc("/{key}", Redirect)
	return route
}

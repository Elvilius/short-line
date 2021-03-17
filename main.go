package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Elvilius/short-line/repository"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strconv"
)

type Handler struct {
	repository *db.UrlRepository
}

type UrlResponse struct {
	Data string `json:"data"`
}

type UrlRequest struct {
	Url string `json:"url"`
}

func main() {
	repository := db.Connect(os.Getenv("PSQL_URL"))
	handler := Handler{repository: &repository}
	
	defer repository.Conn.Close(context.Background())

	route := mux.NewRouter()
	route.HandleFunc("/", IndexHandler).Methods("GET")
	route.HandleFunc("/", handler.CreateShortUrlHandler).Methods("POST")
	route.HandleFunc("/{key}", handler.RedirectHandler)

	fmt.Println("Server listening!")
	http.ListenAndServe(":5656", route)
}

func (h *Handler) CreateShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	var id int
	 urlRequest:= UrlRequest{}

	if err := json.NewDecoder(r.Body).Decode(&urlRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	existUrl, _ := h.repository.GetUrlByFullAddres(urlRequest.Url)

	if existUrl == nil {
		id, _ = h.repository.CreateUrl(urlRequest.Url)

	} else {
		id = existUrl.Id
	}
	shortUrl := createShortUrl(id)
	json.NewEncoder(w).Encode(UrlResponse{Data: shortUrl})

}

func (h *Handler) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["key"]
	url, err := h.repository.GetUrlById(id)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	http.Redirect(w, r, url.Url, http.StatusMovedPermanently)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to this life-changing API.")
}

func createShortUrl(id int) string {
	return os.Getenv("HOST_URL") + "/" + strconv.Itoa(id)
}

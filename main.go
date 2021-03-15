package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"net/http"
	"os"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("PSQL_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var id int
	err = conn.QueryRow(context.Background(), "INSERT INTO urls (id, full_address_name, short_key) VALUES ($1, $2, $3) RETURNING id", "1", "aaa", "eee").Scan(&id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	route := mux.NewRouter()

	route.HandleFunc("/", Index)
	route.HandleFunc("/{key}", Redirect)
	route.HandleFunc("/create", CreateShortUrl).Methods("POST")

	fmt.Println("Server listening!")
	http.ListenAndServe(":5656", route)
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
	http.Redirect(w, r, "https://github.com/Elvilius", http.StatusMovedPermanently)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to this life-changing API.")
}

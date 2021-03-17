package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

type UrlRepository struct {
	Conn *pgx.Conn
}

type UrlModel struct {
	Id              int
	Url string
}

func (db *UrlRepository) CreateUrl(url string) (int, error) {
	var id int
	err := db.Conn.QueryRow(context.Background(), "INSERT INTO urls (url, short_key) VALUES ($1, $2) RETURNING id", url, "null").Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (db *UrlRepository) GetUrlByFullAddres(address string) (*UrlModel, error) {
	var url UrlModel
	url = UrlModel{}
	query := `SELECT id, url FROM urls WHERE full_address_name = $1;`
	row := db.Conn.QueryRow(context.Background(), query, address)
	err := row.Scan(&url.Id, &url.Url)
	if err != nil {
		return nil, err
	}
	return &url, err
}

func (db *UrlRepository) GetUrlById(id string) (UrlModel, error) {
	url := UrlModel{}
	query := `SELECT id, url FROM urls WHERE id = $1;`
	row := db.Conn.QueryRow(context.Background(), query, id)
	err := row.Scan(&url.Id, &url.Url)
	if err != nil {
		return url, err
	}
	return url, nil
}

func Connect(psqUrl string) UrlRepository {
	db := UrlRepository{}
	conn, err := pgx.Connect(context.Background(), psqUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	db.Conn = conn
	return db
}

package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

type Db struct {
	Conn *pgx.Conn
}

type UrlModel struct {
	Id                int
	Full_address_name string
}

func (db *Db) CreateUrl(url string) int {
	var id int
	err := db.Conn.QueryRow(context.Background(), "INSERT INTO urls (full_address_name, short_key) VALUES ($1, $2) RETURNING id", url, "null").Scan(&id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	return id
}

func (db *Db) GetUrlByFullAddres(address string) (UrlModel, error) {
	url := UrlModel{}
	query := `SELECT id, full_address_name FROM urls WHERE full_address_name = $1;`
	row := db.Conn.QueryRow(context.Background(), query, address)
	err := row.Scan(&url.Id, &url.Full_address_name)
	if err != nil {
		return url, err
	}
	return url, nil
}

func (db *Db) GetUrlById(id string) UrlModel {
	url := UrlModel{}
	query := `SELECT id, full_address_name FROM urls WHERE id = $1;`
	row := db.Conn.QueryRow(context.Background(), query, id)
	err := row.Scan(&url.Id, &url.Full_address_name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	return url
}

func Connect(psqUrl string) Db {
	db := Db{}
	conn, err := pgx.Connect(context.Background(), psqUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	db.Conn = conn
	return db
}

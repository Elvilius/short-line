package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

type Db struct {
	conn *pgx.Conn
}

type UrlModel struct {
	Id                int
	Full_address_name string
}

func (db *Db) CreateUrl(url string) int {
	var id int
	err := db.conn.QueryRow(context.Background(), "INSERT INTO urls (id, full_address_name, short_key) VALUES ($1, $2, $3) RETURNING id", "1", url, "eee").Scan(&id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	return id
}

func (db *Db) GetUrlByFullAddres(address string) UrlModel {
	url := UrlModel{}
	query := `SELECT * FROM urls WHERE full_address_name = $1;`
	row := db.conn.QueryRow(context.Background(), query, address)
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
	defer conn.Close(context.Background())

	db.conn = conn
	return db
}

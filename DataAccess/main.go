package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const pathDB string = "./identifier.sqlite"

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float64
}

func main() {
	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)
}

func PrepareSqlite() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", pathDB)
	if err != nil {
		panic(err)
	}

	return db, nil
}

// albumsByArtist queries for albums that habe the specified artista name.
func albumsByArtist(name string) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []Album

	db, err := PrepareSqlite()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT id, title, artist, price FROM album WHERE artist= ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	var alb Album
	for rows.Next() {
		err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
		if err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

// Continue https://go.dev/doc/tutorial/database-access

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const pathDB string = "./identifier.sqlite"

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float64
}

func checkErrorFree(err error) bool {
	if err != nil {
		panic(err)
	}
	return true
}

func main() {
	albums, err := albumsByArtist("John Coltrane")
	checkErrorFree(err)
	fmt.Printf("Albums found: %v\n", albums)

	alb, err := albumByID(2)
	checkErrorFree(err)
	fmt.Printf("Album found: %v\n", alb)

	albId, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	checkErrorFree(err)
	fmt.Printf("ID of added album: %v\n", albId)

	if delAlbum(albId) {
		fmt.Printf("Album Id: %v deleted", albId)
	}
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

// albumByID queries for the album with the specified ID.
func albumByID(id int64) (Album, error) {
	// An Album to hols data from the returned row.
	var alb Album

	db, err := PrepareSqlite()
	checkErrorFree(err)

	row := db.QueryRow("SELECT id, title, artist, price FROM album WHERE id = ?", id)
	checkErrorFree(err)
	err = row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
	if err == sql.ErrNoRows {
		return alb, fmt.Errorf("albumByID %d: no such album", id)
	} else if err != nil {
		return alb, fmt.Errorf("albumByID %d: %v", id, err)
	}

	return alb, nil

}

// addAlbum adds specified album to the database.
// returning the album ID of the new entry
func addAlbum(alb Album) (int64, error) {
	db, err := PrepareSqlite()
	checkErrorFree(err)

	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	checkErrorFree(err)

	id, err := result.LastInsertId()
	checkErrorFree(err)

	return id, nil
}

// delAlbum delete specified album to the database
// returning the object deleted.
func delAlbum(id int64) bool {
	db, err := PrepareSqlite()
	checkErrorFree(err)

	result, err := db.Exec("DELETE FROM album where id = ?", id)
	checkErrorFree(err)

	rows, err := result.RowsAffected()
	checkErrorFree(err)

	return rows > 0

}

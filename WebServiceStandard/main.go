package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// album represents data about records album
type Album struct {
	ID     string
	Title  string
	Artist string
	Price  float64
}

// albums slice to seed record album data.
var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World 👋!")
	})

	mux.HandleFunc("/albums", getAlbums)
	mux.HandleFunc("/album/:id/", getAlbumByID)
	mux.HandleFunc("/album", postAlbum)

	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}

// renderJSON renders 'v' as JSON and writes it as a response into w.
func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(w http.ResponseWriter, r *http.Request) {
	renderJSON(w, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbum(w http.ResponseWriter, r *http.Request) {

	newAlbum := new(Album)

	err := json.NewDecoder(r.Body).Decode(&newAlbum)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Add the new album to the slice
	albums = append(albums, *newAlbum)
	renderJSON(w, albums)

	// curl http://localhost:3000/albums --include --header "Content-Type: application/json" --request "POST" --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			renderJSON(w, a)
		}
	}

	renderJSON(w, "Album Not Found.")
}

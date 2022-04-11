package main

import (
	"github.com/gofiber/fiber/v2"
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
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/albums", getAlbums)
	app.Get("/album/:id", getAlbumByID)
	app.Post("/album", postAlbum)

	app.Listen(":3000")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *fiber.Ctx) error {
	return c.JSON(albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbum(c *fiber.Ctx) error {
	newAlbum := new(Album)

	// Call BindJSON to bind the receibed JSON to newAlbum
	if err := c.BodyParser(newAlbum); err != nil {
		return err
	}

	// Add the new album to the slice
	albums = append(albums, *newAlbum)
	return c.JSON(newAlbum)

	// curl http://localhost:3000/albums --include --header "Content-Type: application/json" --request "POST" --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *fiber.Ctx) error {
	id := c.Params("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			return c.JSON(a)
		}
	}
	return fiber.ErrNotFound
}

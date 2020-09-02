package main

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(c *fiber.Ctx) {
		c.JSON("[{1:'Hello, World 👋!'}]")
	})

	type Tut struct {
		Title       string   `json:"title"`
		Description []string `json:"description"`
		Published   bool     `json:"published"`
		ID          string   `json:"id"`
	}

	tuts := []Tut{
		{
			Title:       "ravi",
			Description: []string{"art", "coding", "music", "travel"},
			ID:          "9420b12e-5220-4443-81c4-b482753551ad",
			Published:   true,
		}, {
			Title:       "rqvi",
			Description: []string{"art", "coding", "music", "travel"},
			ID:          "9420b12e-5q443-81cqqq82753551ad",
			Published:   true,
		},
	}

	tuts = append(tuts, []Tut{
		{
			Title:       "ravi",
			Description: []string{"art", "coding", "music", "travel"},
			ID:          "9420b12e-5220-4443-81c4-b482753551ad",
			Published:   true,
		}}...)

	app.Get("/api/tutorials", func(c *fiber.Ctx) {
		c.JSON(tuts)
	})

	type Staff struct {
		Title     string `json:"title"`
		Published bool   `json:"published"`
		ID        string `json:"id"`
	}

	app.Listen(3000)
}

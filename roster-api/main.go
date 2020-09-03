package main

import (
	"roster-api/middleware"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	_ "github.com/lib/pq"
)

/* type employee struct {
	ID     string `json: "Id"`
	Name   string `json: "name"`
	Salary string `json: "salary"`
	Age    string `json: "age"`
} */

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/api/rosters", middleware.GetAllRoster)
	app.Post("/api/insertrosters", middleware.CreateRoster)

	app.Listen(3000)
}

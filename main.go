package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"math/rand"
	_ "math/rand"
	"strconv"
	"strings"
)

var port = ":3000"

func main() {

	app := fiber.New(fiber.Config{
		AppName:      "Go Fiber Api",
		ServerHeader: "Go Fiber",
	})
	api := app.Group("/api", apiHandler) // /api
	v1 := api.Group("/v1", handler)      // /api/v1
	v1.Get("/list", func(c *fiber.Ctx) error {
		return c.SendString("api v1 list")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/json", func(c *fiber.Ctx) error {

		// variable from params
		message := c.Query("name", "World")

		message = message + " 👋!"

		// get message from query string again
		message2 := c.Query("name")

		log.Println(message2)

		// return a sample JSON response
		return c.JSON(fiber.Map{
			"message": message,
		})
	})

	app.Get("test/:val?", func(c *fiber.Ctx) error {

		val := c.Params("val", "test1")

		// return a sample JSON response
		return c.JSON(fiber.Map{
			"val": val,
		})
	})

	// Match request starting with /api
	app.Use("/api", customHeaderMiddleware)

	app.Get("api/*", func(c *fiber.Ctx) error {

		path := c.Params("*")

		// split path by /
		var parts = strings.Split(path, "/")

		// get last part of path
		val := parts[len(parts)-1]

		// return a sample JSON response
		return c.JSON(fiber.Map{
			"last path part": val,
		})
	})

	app.Get("/error", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusBadRequest, "Custom error message")
	})

	app.Static("/", "./public")

	log.Fatal(app.Listen(port))
}

func apiHandler(ctx *fiber.Ctx) error {

	log.Println("apiHandler")
	return ctx.Next()
}

func handler(ctx *fiber.Ctx) error {

	log.Println("handler")
	return ctx.Next()
}

func customHeaderMiddleware(c *fiber.Ctx) error {

	c.Set(
		"X-Custom-Header",
		"Random number: "+strconv.Itoa(rand.Intn(100)),
	)
	return c.Next()
}

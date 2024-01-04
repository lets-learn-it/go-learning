package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Line struct {
	P1 Point `json:"p1"`
	P2 Point `json:"p2"`
}

func main() {
	rand.Seed(time.Now().UnixNano())
	app := fiber.New()

	app.Get("/api/sse", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/event-stream")
		c.Set("Cache-Control", "no-cache")
		c.Set("Connection", "keep-alive")
		c.Set("Transfer-Encoding", "chunked")

		c.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
			fmt.Println("WRITER")
			var p1, p2 Point
			var line Line
			for {
				p1.X = rand.Intn(400)
				p1.Y = rand.Intn(400)

				p2.X = rand.Intn(400)
				p2.Y = rand.Intn(400)

				line.P1 = p1
				line.P2 = p2

				data, err := json.Marshal(line)
				if err != nil {
					log.Println(err)
				}

				fmt.Fprintf(w, "data: %s\n\n", data)

				err = w.Flush()
				if err != nil {
					// Refreshing page in web browser will establish a new
					// SSE connection, but only (the last) one is alive, so
					// dead connections must be closed here.
					fmt.Printf("Error while flushing: %v. Closing http connection.\n", err)

					break
				}
				time.Sleep(time.Second)
			}
		}))

		return nil
	})

	app.Static("/", "./static")

	log.Fatal(app.Listen(":8080"))
}

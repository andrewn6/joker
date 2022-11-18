package main

import (
    "log"
    "io/ioutil"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/jokes", func (c *fiber.Ctx) error {
        file, err := ioutil.ReadFile("./jokes/jokes.txt")       
        if err != nil {
           log.Fatalf("unable to read file: %v", err)
        }
        c.Set("Content-Type", "text")
        return c.Send(file)
    })

    log.Fatal(app.Listen(":8080"))
}


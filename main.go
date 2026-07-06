package main

import (
    "fmt"
    "log"
    "github.com/gabrielbismarck/GoMaker/internal/controller"
    "github.com/gabrielbismarck/GoMaker/pkg/index"
    "github.com/gofiber/fiber/v2"
)

func main() {

    myIndexer := index.NewIndexer("default")

	controller.SetIndexer(myIndexer)

	app := fiber.New()

	app.Post("/index", controller.AddDocumentToIndex)

	app.Post("/save-index", func(c *fiber.Ctx) error {
		if err := myIndexer.SaveIndex(); err != nil {
			log.Printf("Erro ao salvar o índice: %v", err)
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Erro ao salvar o índice: %v", err))
		}
		return c.SendString("Índice salvo com sucesso!")
	})

	log.Fatal(app.Listen(":3000"))

}

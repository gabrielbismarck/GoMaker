package controller

import (
    "github.com/gabrielbismarck/GoMaker/pkg/index"
    "github.com/gofiber/fiber/v2"
)

var globalIndexer *index.Indexer

func SetIndexer(idx *index.Indexer) {
    globalIndexer = idx
}

type Document struct {
    Url     string `json:"url" xml:"url" form:"url"`
    Content string `json:"content" xml:"content" form:"content"`
}

func AddDocumentToIndex(c *fiber.Ctx) error {
    doc := new(Document)

	if err := c.BodyParser(doc); err != nil {
		return err
	}

	if globalIndexer == nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Indexer not initialized")
	}

	globalIndexer.AddDocToIndex(doc.Url, doc.Content)

	return c.SendString("Document added")

}

package controller

import (
	"strings"

	"github.com/gabrielbismarck/GoMaker/pkg/search"
	"github.com/gofiber/fiber/v2"
)

// SearchQuery: Processa a requisição GET /search?q=palavra
func SearchQuery(c *fiber.Ctx) error {
	query := c.Query("q")
	queryType := c.Query("type", "SIMPLE")

	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Busca vazia"})
	}

	// Análise Léxica: quebra a frase em tokens (palavras)
	terms := strings.Fields(strings.ToLower(query))

	// Chamada da sua lógica (pkg/search) usando o globalIndexer
	results := search.Search(
		terms,
		queryType,
		globalIndexer.Index,
		globalIndexer.DocFreq,
		len(globalIndexer.Documents),
	)

	return c.JSON(results)
}

// --- PARTE DA MARIA ---
func AggregateDistributedResults(query string) []search.SearchResult {
	// Implementação
	return nil
}

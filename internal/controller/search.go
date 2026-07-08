package controller

import (
	"strings"

	"github.com/gabrielbismarck/GoMaker/pkg/search"
	"github.com/gofiber/fiber/v2"
)

// Abstração de Endereçamento: Lista de nós (servidores) para a busca distribuída.
// Para o seu teste, usaremos as portas do simulador que criamos.
var remoteNodes = []string{
	"http://localhost:3001", "http://localhost:3002", "http://localhost:3003",
	"http://localhost:3004", "http://localhost:3005", "http://localhost:3006",
	"http://localhost:3007", "http://localhost:3008", "http://localhost:3009",
	"http://localhost:3010",
}

// SearchQuery: Processa a requisição GET /search?q=palavra&dist=true
func SearchQuery(c *fiber.Ctx) error {
	query := c.Query("q")
	queryType := c.Query("type", "SIMPLE")
	distributed := c.Query("dist") == "true" // Verifica se o usuário quer busca distribuída

	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Busca vazia"})
	}

	var results []search.SearchResult

	if distributed {
		// Chamada para a busca distribuída
		results = AggregateDistributedResults(query)
	} else {
		// Busca Local
		terms := strings.Fields(strings.ToLower(query))
		results = search.Search(
			terms,
			queryType,
			globalIndexer.Index,
			globalIndexer.DocFreq,
			len(globalIndexer.Documents),
		)
	}

	return c.JSON(results)
}

// AggregateDistributedResults: Orquestrador da busca em múltiplos nós.
// Implementa o modelo CSP para Concorrência Nativa
func AggregateDistributedResults(query string) []search.SearchResult {
	// Operador 'make(chan)': Cria um CANAL para comunicação entre Goroutines.
	// O canal garante a sincronização e evita Race Conditions
	resultsChan := make(chan []search.SearchResult)

	// Dispara uma busca para cada nó remoto em paralelo.
	for _, url := range remoteNodes {
		// Operador 'go': Inicia uma Goroutine (Thread leve de 2KB)
		// Isso permite consultar dezenas de servidores simultaneamente sem travar o programa.
		go func(nodeURL string) {
			// Chamada de rede para o nó remoto
			res := search.SearchInRemoteNode(nodeURL, query)

			// Operador '<-': Envia o resultado da busca para o canal.
			resultsChan <- res
		}(url)
	}

	var allResults []search.SearchResult

	// Coleta os resultados conforme eles chegam no canal.
	for i := 0; i < len(remoteNodes); i++ {
		// Operador '<-': Recebe o dado vindo de uma Goroutine.
		// O loop aguarda (bloqueia) até que um resultado esteja disponível no canal.
		nodeRes := <-resultsChan

		// Operador '...': Descompacta o slice de resultados para adicionar à lista global.
		allResults = append(allResults, nodeRes...)
	}

	// Após reunir todos, chamo a função de ranking para ordenar por score TF-IDF.
	return search.RankDistributed(allResults)
}

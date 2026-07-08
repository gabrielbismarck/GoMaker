package search

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"sort"

	"github.com/gabrielbismarck/GoMaker/pkg/index"
)

// SearchResult: Abstração de Dados que agrupa o documento e sua nota (r-value).
type SearchResult struct {
	Document string
	Score    float64
}

// Search: Função principal. Orquestra a busca local e o ranking.
func Search(terms []string, queryType string, idx index.InvertedIndex, docFreq index.DocumentFrequency, numDocs int) []SearchResult {
	scores := make(map[string]float64)

	if queryType == "AND" {
		for _, doc := range intersectDocs(terms, idx) {
			scores[doc] = scoreDoc(terms, doc, idx, docFreq, numDocs)
		}
	} else {
		for _, term := range terms {
			for doc := range idx[term] {
				scores[doc] += scoreDoc([]string{term}, doc, idx, docFreq, numDocs)
			}
		}
	}

	return rankResults(scores)
}

// scoreDoc: Implementa o cálculo de relevância TF-IDF (Semântica Matemática).
func scoreDoc(terms []string, doc string, idx index.InvertedIndex, docFreq index.DocumentFrequency, numDocs int) float64 {
	score := 0.0
	for _, term := range terms {
		// Operador float64(): Coerção explícita para garantir precisão (Confiabilidade).
		tf := float64(idx[term][doc])
		// Operador math.Log: Calcula a raridade do termo (IDF) [Victor Lavrenko, 184].
		idf := math.Log(float64(numDocs) / float64(docFreq[term]))
		score += tf * idf
	}
	return score
}

// intersectDocs: Lógica de conjunto "E" usando filtragem por exclusão.
func intersectDocs(terms []string, index index.InvertedIndex) []string {
	if len(terms) == 0 {
		return nil
	}
	docs := make(map[string]bool)
	// Inicializa com a primeira palavra.
	for doc := range index[terms[0]] {
		docs[doc] = true
	}
	// Filtra com as demais palavras.
	for _, term := range terms[1:] {
		for doc := range docs {
			if _, exists := index[term][doc]; !exists {
				delete(docs, doc) // Operador 'delete': Remove o que não satisfaz a interseção.
			}
		}
	}
	result := []string{}
	for doc := range docs {
		result = append(result, doc)
	}
	return result
}

// rankResults: Ordena os resultados locais.
func rankResults(scores map[string]float64) []SearchResult {
	results := make([]SearchResult, 0, len(scores))
	for doc, score := range scores {
		results = append(results, SearchResult{Document: doc, Score: score})
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})
	return results
}

// SearchInRemoteNode: Abstração de Comunicação via Rede.
// Faz uma requisição HTTP GET para outro nó do cluster
func SearchInRemoteNode(serverURL string, query string) []SearchResult {
	// Operador fmt.Sprintf: Formata a URL de destino dinamicamente.
	fullURL := fmt.Sprintf("%s/search?q=%s", serverURL, query)

	// Operador http.Get: Realiza a chamada de rede (Comunicação entre Unidades).
	resp, err := http.Get(fullURL)
	if err != nil {
		// Confiabilidade: Se o servidor remoto falhar, retorna uma lista vazia
		// para não travar o sistema principal
		return []SearchResult{}
	}

	// Operador 'defer': Garante o fechamento do corpo da resposta (Gerência de Recursos).
	defer resp.Body.Close()

	var results []SearchResult
	// Operador json.NewDecoder: Transforma o fluxo de bytes JSON de volta em uma Struct Go.
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return []SearchResult{}
	}

	return results
}

// RankDistributed: Agregador Global de Scores.
// Junta os resultados de todos os nós em um ranking final único
func RankDistributed(allResults []SearchResult) []SearchResult {
	if len(allResults) == 0 {
		return []SearchResult{}
	}

	globalScores := make(map[string]float64)
	maxScore := 0.0

	// 1. FASE DE AGREGAÇÃO (REDUCE)
	for _, res := range allResults {
		globalScores[res.Document] += res.Score
		// Operador de Comparação '>': Identifica o valor de pico para normalização.
		if globalScores[res.Document] > maxScore {
			maxScore = globalScores[res.Document]
		}
	}

	// 2. FASE DE NORMALIZAÇÃO E MAPEAMENTO
	finalResults := make([]SearchResult, 0, len(globalScores))
	for doc, score := range globalScores {
		// Operador '/': Realiza a divisão em ponto flutuante para a escala [0..1].
		normalizedScore := score / maxScore

		finalResults = append(finalResults, SearchResult{
			Document: doc,
			Score:    normalizedScore,
		})
	}

	// 3. ORDENAÇÃO FINAL
	sort.Slice(finalResults, func(i, j int) bool {
		return finalResults[i].Score > finalResults[j].Score
	})

	return finalResults
}

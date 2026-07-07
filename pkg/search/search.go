package search

import (
	"math"
	"sort"
	"github.com/gabrielbismarck/GoMaker/pkg/index"
)

// SearchResult: Abstração de Dados que agrupa o documento e sua nota (r-value).
type SearchResult struct {
	Document string
	Score    float64
}

// Search: Função principal. Orquestra a busca e o ranking.
func Search(terms []string, queryType string, idx index.InvertedIndex, docFreq index.DocumentFrequency, numDocs int) []SearchResult {
	// Operador 'make': aloca o mapa de notas na memória.
	scores := make(map[string]float64)

	if queryType == "AND" {
		// Lógica de Interseção (AND)
		for _, doc := range intersectDocs(terms, idx) {
			scores[doc] = scoreDoc(terms, doc, idx, docFreq, numDocs)
		}
	} else {
		// Lógica de União (OR/SIMPLE)
		for _, term := range terms {
			// Operador 'range': percorre os documentos associados a cada palavra.
			for doc := range idx[term] {
				scores[doc] += scoreDoc([]string{term}, doc, idx, docFreq, numDocs)
			}
		}
	}

	return rankResults(scores)
}

// scoreDoc: Implementa o cálculo de relevância TF-IDF.
func scoreDoc(terms []string, doc string, idx index.InvertedIndex, docFreq index.DocumentFrequency, numDocs int) float64 {
	score := 0.0
	for _, term := range terms {
		// Operador float64(): conversão explícita para garantir precisão matemática.
		tf := float64(idx[term][doc])

		// math.Log: operador de logaritmo para normalizar pesos (Semântica).
		idf := math.Log(float64(numDocs) / float64(docFreq[term]))
		
		score += tf * idf
	}
	return score
}

// intersectDocs: Lógica de conjunto "E".
func intersectDocs(terms []string, index index.InvertedIndex) []string {
	if len(terms) == 0 { return nil }

	docs := make(map[string]bool)

	// Uso 'terms' para extrair a STRING individual.
	// O mapa 'index' não aceita o slice 'terms' inteiro.
	for doc := range index[terms[0]] { 
		docs[doc] = true
	}

	for _, term := range terms[1:] {
		for doc := range docs {
			if _, exists := index[term][doc]; !exists {
				delete(docs, doc) // Operador 'delete': limpa o conjunto.
			}
		}
	}

	result := []string{}
	for doc := range docs {
		result = append(result, doc)
	}
	return result
}

// rankResults: Ordena os resultados usando uma função Lambda (anônima).
func rankResults(scores map[string]float64) []SearchResult {
	results := make([]SearchResult, 0, len(scores))
	for doc, score := range scores {
		results = append(results, SearchResult{Document: doc, Score: score})
	}
	// sort.Slice: conceito de Funções de Ordem Alta.
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})
	return results
}

// --- PARTE DA MARIA (DISTRIBUÍDA) ---
func SearchInRemoteNode(serverURL string, query string) []SearchResult {
	// Implementação
	return nil 
}
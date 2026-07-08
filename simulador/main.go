package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gabrielbismarck/GoMaker/pkg/search"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Crio 10 instâncias de servidores (nós).
	// O operador 'for' define a iteração sobre o intervalo de portas.
	for i := 3001; i <= 3010; i++ {

		// OPERADOR 'go': Concorrência Nativa (Modelo CSP).
		// Dispara uma Goroutine (Thread leve de 2KB) para cada nó
		go func(porta int) {
			// Encapsulamento do servidor Fiber
			app := fiber.New(fiber.Config{DisableStartupMessage: true})

			app.Get("/search", func(c *fiber.Ctx) error {
				// Entrada: Recupera e normaliza a query.
				query := strings.ToLower(strings.TrimSpace(c.Query("q")))
				fmt.Printf("🔍 [Nó %d] Recebeu busca por: \"%s\"\n", porta, query)

				// Define a pasta exclusiva deste nó.
				caminhoPasta := fmt.Sprintf("./data/no-%d", porta)
				var resultadosLocais []search.SearchResult

				// Tenta acessar o hardware (disco) para ler o índice local.
				arquivos, err := os.ReadDir(caminhoPasta)
				if err != nil {
					// Se a pasta não existe, retorna vazio (Zero-Value).
					return c.JSON([]search.SearchResult{})
				}

				// O operador strings.Fields quebra a frase em palavras individuais.
				// Transforma "bismarck gabriel" em ["bismarck", "gabriel"].
				termosBusca := strings.Fields(query)

				// Itera sobre a coleção de arquivos da pasta.
				for _, f := range arquivos {
					if f.IsDir() {
						continue
					}

					caminhoArquivo := fmt.Sprintf("%s/%s", caminhoPasta, f.Name())
					conteudo, _ := os.ReadFile(caminhoArquivo)
					textoArquivo := strings.ToLower(string(conteudo))

					var scoreAcumulado float64 = 0.0
					encontrouAlgumTermo := false

					// Para cada palavra da busca, contamos as ocorrências no arquivo.
					for _, termo := range termosBusca {
						// O operador strings.Count retorna a frequência do lexema.
						ocorrencias := strings.Count(textoArquivo, termo)
						if ocorrencias > 0 {
							encontrouAlgumTermo = true
							// Acumula o score dinamicamente.
							scoreAcumulado += float64(ocorrencias)
						}
					}

					if encontrouAlgumTermo {
						// Adiciona o arquivo com sua nota real.
						resultadosLocais = append(resultadosLocais, search.SearchResult{
							Document: f.Name(),
							Score:    scoreAcumulado,
						})
						fmt.Printf("✅ [Nó %d] Encontrado em: %s | Score: %.2f\n", porta, f.Name(), scoreAcumulado)
					}
				}

				// Retorna o resultado serializado em JSON.
				return c.JSON(resultadosLocais)
			})

			// O método Listen mantém a Goroutine viva servindo a porta.
			if err := app.Listen(fmt.Sprintf(":%d", porta)); err != nil {
				log.Printf("❌ Erro fatal no nó %d: %v", porta, err)
			}
		}(i)
	}

	fmt.Println("🚀 Search GoMaker: Suporte a busca distribuida.")

	// Bloqueio infinito para manter o processo pai vivo.
	select {}
}

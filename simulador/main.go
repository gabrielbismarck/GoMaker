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
	// O operador 'for' define a faixa de portas (3001 a 3010).
	for i := 3001; i <= 3010; i++ {

		// OPERADOR 'go': Dispara uma Goroutine independente para cada nó.
		// Isso permite concorrência nativa com baixo custo (2KB por nó)
		go func(porta int) {
			// Encapsulamento do servidor Fiber sem mensagens de log de inicialização (Startup)
			app := fiber.New(fiber.Config{DisableStartupMessage: true})

			app.Get("/search", func(c *fiber.Ctx) error {
				// (Entrada): Recupera o termo e normaliza para minúsculo.
				query := strings.ToLower(strings.TrimSpace(c.Query("q")))

				// LOG DE VALIDAÇÃO: Importante para provar que o nó recebeu a mensagem.
				fmt.Printf("🔍 [Nó %d] Recebeu solicitação de busca por: \"%s\"\n", porta, query)

				// O nó só acessa sua respectiva pasta.
				caminhoPasta := fmt.Sprintf("./data/no-%d", porta)

				var resultadosLocais []search.SearchResult

				// Lê a lista de arquivos físicos no disco.
				arquivos, err := os.ReadDir(caminhoPasta)
				fmt.Printf("📂 [Nó %d] Vasculhando a pasta: %s (Encontrou %d arquivos)\n", porta, caminhoPasta, len(arquivos))
				if err != nil {
					// Se a pasta não existe, retorna vazio sem travar o sistema.
					return c.JSON([]search.SearchResult{})
				}

				// OPERADOR 'range': Itera sobre a coleção de arquivos encontrados.
				for _, f := range arquivos {
					// Ignora pastas acidentais, foca apenas em arquivos.
					if f.IsDir() {
						continue
					}

					caminhoArquivo := fmt.Sprintf("%s/%s", caminhoPasta, f.Name())
					conteudo, err := os.ReadFile(caminhoArquivo)
					if err != nil {
						continue
					}

					// MOTOR DE BUSCA (Análise de Conteúdo):
					// Normalizamos o conteúdo do arquivo para garantir o 'match'.
					textoArquivo := strings.ToLower(string(conteudo))

					if strings.Contains(textoArquivo, query) {
						// Adiciona o documento encontrado à lista.
						resultadosLocais = append(resultadosLocais, search.SearchResult{
							Document: f.Name(),
							Score:    1.0,
						})
						fmt.Printf("   ✅ [Nó %d] Termo encontrado no arquivo: %s\n", porta, f.Name())
					}
				}

				// (Saída): Retorna o slice como JSON para o Orquestrador.
				return c.JSON(resultadosLocais)
			})

			// O método Listen bloqueia a Goroutine, mantendo o nó ativo.
			if err := app.Listen(fmt.Sprintf(":%d", porta)); err != nil {
				log.Printf("❌ Erro fatal no nó %d: %v", porta, err)
			}
		}(i)
	}

	fmt.Println("🚀 Cluster GoMaker ON: 10 nós independentes prontos para busca real.")

	// Bloqueio infinito para manter o processo pai vivo enquanto as goroutines rodam.
	select {}
}

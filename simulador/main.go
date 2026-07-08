package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/gabrielbismarck/GoMaker/pkg/search"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Criando 10 nós virtuais
	for i := 3001; i <= 3010; i++ {
		go func(porta int) {
			app := fiber.New(fiber.Config{DisableStartupMessage: true})

			app.Get("/search", func(c *fiber.Ctx) error {
				// 1. Recupera o termo de busca (Abstração de Entrada)
				query := strings.ToLower(c.Query("q"))

				// 2. Define qual a "estante" de livros deste nó (Encapsulamento)
				caminhoPasta := fmt.Sprintf("./data/no-%d", porta)

				var resultadosLocais []search.SearchResult

				// 3. Lê os arquivos físicos da pasta
				arquivos, err := ioutil.ReadDir(caminhoPasta)
				if err != nil {
					// Se a pasta não existir, retorna lista vazia (Confiabilidade)
					return c.JSON([]search.SearchResult{})
				}

				for _, f := range arquivos {
					// 4. OPERADOR 'ReadFile': Carrega o conteúdo do disco para a RAM
					conteudo, _ := ioutil.ReadFile(caminhoPasta + "/" + f.Name())

					// 5. ANÁLISE LÉXICA: Verifica se a palavra existe no texto
					// strings.Contains é o nosso motor de busca simplificado
					if strings.Contains(strings.ToLower(string(conteudo)), query) {
						resultadosLocais = append(resultadosLocais, search.SearchResult{
							Document: f.Name(),
							Score:    1.0, // Define um score real para termos encontrados
						})
					}
				}

				// 6. SÍNTESE DE DADOS: Retorna apenas os arquivos que deram 'match'
				// Se nenhum arquivo tiver a palavra, retornará um JSON vazio '[]'
				return c.JSON(resultadosLocais)
			})

			if err := app.Listen(fmt.Sprintf(":%d", porta)); err != nil {
				log.Printf("Erro no nó %d: %v", porta, err)
			}
		}(i)
	}
	select {}
}

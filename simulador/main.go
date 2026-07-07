package main

import (
	"fmt"
	"log"
	"github.com/gabrielbismarck/GoMaker/pkg/index"
	"github.com/gofiber/fiber/v2"
)

func main() {
	inicioPorta := 3001
	fimPorta := 3003

	fmt.Println("🚀 Iniciando nós simulados...")

	for i := inicioPorta; i <= fimPorta; i++ {
		go func(porta int) {
			app := fiber.New(fiber.Config{
				DisableStartupMessage: true,
			})

			noID := fmt.Sprintf("no-%d", porta)
			
			// 1. Abstração de Dados: Criamos o indexador para o nó.
			idx := index.NewIndexer(noID)

			// 2. OPERADOR DE DESCARTE (_): 
			// Como o compilador de Go proíbe variáveis não utilizadas (Eficiência),
			// usamos o '_' para dizer ao tradutor: "Eu sei que 'idx' existe, 
			// mas não preciso acessar seus campos agora".
			_ = idx 

			fmt.Printf("✅ Nó ativo na porta %d\n", porta)

			app.Get("/search", func(c *fiber.Ctx) error {
				// Simulação de resposta estruturada (SearchResult)
				return c.Status(200).JSON([]fiber.Map{
					{
						"Document": fmt.Sprintf("livro_da_estante_%d.txt", porta),
						"Score":    0.95,
					},
				})
			})

			if err := app.Listen(fmt.Sprintf(":%d", porta)); err != nil {
				log.Printf("Erro no nó %d: %v", porta, err)
			}
		}(i)
	}
	// Bloqueia a Goroutine principal para manter os nós vivos.
	select {}
}

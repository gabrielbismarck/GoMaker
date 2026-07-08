package main

import (
	"fmt"
	"log"

	"github.com/gabrielbismarck/GoMaker/internal/controller"
	"github.com/gabrielbismarck/GoMaker/pkg/index"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Inicialização: Crio o objeto que guardará o índice na memória.
	myIndexer := index.NewIndexer("default")

	// Injeção de Dependência: Disponibilizamos o indexador para o controlador.
	controller.SetIndexer(myIndexer)

	// Inicialização do framework Fiber.
	app := fiber.New()

	// MIDDLEWARE DE CORS: Garante a Confiabilidade da comunicação.
	// Permite que o navegador aceite os dados JSON sem bloqueios de segurança
	app.Use(cors.New())

	// OPERADOR 'app.Static' (A SOLUÇÃO): Abstração de Sistema de Arquivos.
	// Mapeia a rota raiz "/" para a pasta atual "./".
	// Isso faz com que o arquivo 'interface.html' seja servido via HTTP.
	app.Static("/", "./")

	// --- ROTAS DO SISTEMA ---

	// Rota de Indexação: Recebe os documentos via POST.
	app.Post("/index", controller.AddDocumentToIndex)

	// Rota de Busca: Consulta de recuperação de dados via GET.
	app.Get("/search", controller.SearchQuery)

	// Rota de Persistência: Salva o índice no disco.
	app.Post("/save-index", func(c *fiber.Ctx) error {
		if err := myIndexer.SaveIndex(); err != nil {
			log.Printf("Erro ao salvar o índice: %v", err)
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Erro ao salvar o índice: %v", err))
		}
		return c.SendString("Índice salvo com sucesso!")
	})

	// 6. Eficiência: Inicia o servidor na porta 3000.
	// log.Fatal garante a terminação caso a porta esteja ocupada.
	log.Fatal(app.Listen(":3000"))
}

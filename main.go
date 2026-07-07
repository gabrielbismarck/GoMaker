// package main

// import (
//     "fmt"
//     "log"
//     "github.com/gabrielbismarck/GoMaker/internal/controller"
//     "github.com/gabrielbismarck/GoMaker/pkg/index"
//     "github.com/gofiber/fiber/v2"
// )

// func main() {

//     myIndexer := index.NewIndexer("default")

// 	controller.SetIndexer(myIndexer)

// 	app := fiber.New()

// 	app.Post("/index", controller.AddDocumentToIndex)

// 	app.Post("/save-index", func(c *fiber.Ctx) error {
// 		if err := myIndexer.SaveIndex(); err != nil {
// 			log.Printf("Erro ao salvar o índice: %v", err)
// 			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Erro ao salvar o índice: %v", err))
// 		}
// 		return c.SendString("Índice salvo com sucesso!")
// 	})

// 	log.Fatal(app.Listen(":3000"))

// }


package main

import (
    "fmt"
    "log"
    "github.com/gabrielbismarck/GoMaker/internal/controller"
    "github.com/gabrielbismarck/GoMaker/pkg/index"
    "github.com/gofiber/fiber/v2"
)

func main() {
    // 1. Inicialização e Abstração: Criamos o objeto que guardará o índice na memória.
    myIndexer := index.NewIndexer("default")

    // 2. Injeção de Dependência: Disponibilizamos o indexador para o controlador.
    controller.SetIndexer(myIndexer)

    // 3. Simplicidade: Inicializamos o framework Fiber usando o operador de inferência ':='.
    app := fiber.New()

    // --- ROTAS DO SISTEMA ---

    // Rota de Indexação (Parte do Lucas): Recebe os documentos.
    app.Post("/index", controller.AddDocumentToIndex)

    // CORREÇÃO - SUA PARTE: Rota de Busca (Sua responsabilidade).
    // Usamos o método GET pois buscas são consultas de recuperação de dados.
    app.Get("/search", controller.SearchQuery)

    // Rota de Persistência: Salva o índice no disco.
    app.Post("/save-index", func(c *fiber.Ctx) error {
        // Operador 'if' com tratamento de erro imediato (Confiabilidade).
        if err := myIndexer.SaveIndex(); err != nil {
            log.Printf("Erro ao salvar o índice: %v", err)
            return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Erro ao salvar o índice: %v", err))
        }
        return c.SendString("Índice salvo com sucesso!")
    })

    // 4. Eficiência: Inicia o servidor na porta 3000. 
    // log.Fatal garante que o programa pare se a porta estiver ocupada.
    log.Fatal(app.Listen(":3000"))
}

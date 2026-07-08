# Sistema de Busca e Indexação Distribuído - Go

Projeto acadêmico: demonstrar as premissas da linguagem Go por meio de um sistema backend real, comparando implementações com e sem controle de concorrência.

## Índice

- [Introdução](#introdução)
- [Sobre a Linguagem Go](#sobre-a-linguagem-go)
- [O Projeto](#o-projeto)
- [Referências](#referências)

---

# Introdução

Esta Wiki documenta o **Sistema de Busca e Indexação Distribuído**, desenvolvido em Go. Além de apresentar a arquitetura e o funcionamento do sistema, este material descreve os principais conceitos da linguagem Go, comparando-os com a linguagem C e justificando sua escolha para o desenvolvimento do projeto.

## Sobre a Linguagem Go

### Histórico e Versão
-  **Go** ou **Golang** foi criada no Google por ***Robert Griesemer***, ***Rob Pike*** e ***Ken Thompson***.
- A primeira versão estável (Go 1.0) foi lançada em março de 2012, e a versão atual é a **1.24**, lançada em 2025.
- **Motivação:** Os criadores observaram que as linguagens existentes na época (C++, Java, Python) não estavam atendendo bem às necessidades de desenvolvimento de software em larga escala no Google.
- Dentre as exigências buscavasse:
    - **Eficiência** e **performance** de linguagens compiladas (como C/C++).
    - **Simplicidade** e **facilidade** de uso de linguagens dinâmicas (como Python).
    - **Concorrência nativa** e eficiente para aproveitar os processadores multi-core modernos e lidar com sistemas distribuídos.

### Premissas da Linguagem 
- **Performance e Eficiência:**
    - Binários nativos e de alta performance.
    - Compilação rápida.
    - Coletor de lixo eficiente.
    
- **Simplicidade e Clareza:**
    - Sintaxe simples e concisa
    - Fácil aprendizado e manutenção

> Simplicidade: A linguagem possui poucas palavras-chave e uma sintaxe enxuta, tornando o código mais fácil de ler e manter.

- **Concorrência Nativa:**
    - Goroutines e Channels
    - Modelo CSP
    - Aproveitamento de múltiplos núcleos

> Concorrência: Seu grande diferencial é o suporte nativo à concorrência por meio de goroutines e channels, baseado no modelo CSP.

- **Segurança e Robustez:**
    - Tipagem forte
    - Gerenciamento automático de memória
    - Código confiável

> Segurança: A tipagem forte e o coletor de lixo ajudam a evitar erros comuns e tornam o desenvolvimento mais seguro.

### Domínios de Aplicação
- Programação de Sistemas
    - Desenvolvimento de ferramentas de infraestrutura e utilitários
    - Alto desempenho e eficiência para serviços em nuvem
    - Exemplos: Docker e Kubernetes

- Aplicações Comerciais (Backend)
    - Desenvolvimento de APIs REST, microsserviços e aplicações web
    - Capacidade de processar grande volume de requisições simultâneas
    - Exemplos: Uber, Netflix, Twitch e Google

### Usuário Característico
- Desenvolvedores de Backend e Infraestrutura
    - Construção de APIs, microsserviços e sistemas distribuídos
    - Desenvolvimento para computação em nuvem e contêineres
    - Foco em alto desempenho e escalabilidade

- Equipes que Trabalham com Concorrência
    - Processamento de múltiplas requisições simultâneas
    - Uso de goroutines e channels para paralelismo
    - Desenvolvimento de aplicações responsivas e eficientes

- Empresas e Startups
    - Infraestruturas com crescimento rápido
    - Equilíbrio entre produtividade e performance
    - Facilidade de adoção e manutenção da linguagem

### Construtores da Linguagem

Go oferece diversos recursos para desenvolver aplicações de forma simples, segura e eficiente. Muitos desses recursos possuem equivalentes em C, enquanto outros foram introduzidos para facilitar o desenvolvimento de software moderno.

- Struct (struct)
    - Go: Agrupa diferentes campos em um único tipo, permitindo representar estruturas de dados complexas.
    - C: Possui o mesmo conceito de struct.

> Diferença: Em Go, uma struct pode possuir métodos associados, permitindo combinar dados e comportamento sem a necessidade de classes.

- Funções (func)
    - Go: Blocos reutilizáveis de código que executam tarefas específicas e podem retornar múltiplos valores.
    - C: Também utiliza funções para modularizar o código.

> Diferença: O suporte nativo a múltiplos valores de retorno facilita o tratamento de erros e evita o uso excessivo de parâmetros por referência.

- Variáveis e Constantes (var e const)
    - Go: Permitem armazenar valores e definir constantes. O compilador exige que variáveis declaradas sejam utilizadas.
    - C: Possui variáveis e constantes com finalidade semelhante.

> Diferença: A verificação de variáveis não utilizadas ajuda a reduzir código morto e possíveis erros de programação.

- Declaração Curta (:=)
    - Go: Declara e inicializa uma variável com inferência automática de tipo.
    - C: Não possui um equivalente direto.

> Diferença: É uma das formas mais utilizadas para declarar variáveis locais, tornando o código mais conciso.

- Controle de Fluxo (if, for, switch)
    - Go: Controla a execução do programa por meio de condicionais e laços.
    - C: Possui estruturas equivalentes.

> Diferença: Go não possui while (o for desempenha esse papel) e dispensa o uso de parênteses nas condições.

- Ponteiros (* e &)
    - Go: Permitem acessar e modificar valores por referência.
    - C: Também utiliza ponteiros.

> Diferença: Go não permite aritmética de ponteiros, tornando o código mais seguro e reduzindo erros relacionados ao acesso à memória.

- Interfaces (interface)
    - Go: Definem um conjunto de métodos que um tipo deve implementar, permitindo polimorfismo e desacoplamento.
    - C: Não possui um recurso equivalente na linguagem.

> Diferença: Em Go, a implementação de uma interface é implícita: basta que um tipo implemente todos os métodos exigidos.

- Goroutines (go)
    - Go: Executam funções concorrentemente utilizando o runtime da linguagem.
    - C: Não possui suporte nativo; normalmente utiliza bibliotecas como POSIX Threads (pthread).

> Diferença: Goroutines são muito mais leves que threads do sistema operacional e são gerenciadas automaticamente pelo runtime de Go.

- Canais (chan)
    - Go: Permitem a comunicação segura entre goroutines por meio da troca de mensagens.
    - C: Não possui um mecanismo equivalente na linguagem.

> Diferença: Os canais implementam o modelo CSP (Communicating Sequential Processes), incentivando a comunicação entre processos concorrentes em vez do compartilhamento direto de memória.


Exemplo:
```go
type Ponto struct {
    X, Y int
}

func Mover(p Ponto, dx, dy int) Ponto {
    p.X += dx
    p.Y += dy
    return p
}

func main() {
    p1 := Ponto{X: 1, Y: 2}
    p2 := Mover(p1, 5, 3)
}
```

```c
struct Ponto {
    int X;
    int Y;
};

struct Ponto Mover(struct Ponto p, int dx, int dy) {
    p.X += dx;
    p.Y += dy;
    return p;
}

int main() {
    struct Ponto p1 = {1, 2};
    struct Ponto p2 = Mover(p1, 5, 3);
    return 0;
}
```

### Legibilidade
Legibilidade é a facilidade com que um programa pode ser lido e entendido.

- Simplicidade Global
    - Sintaxe enxuta e poucas palavras-chave
    - Apenas um tipo de laço (for)
    - Código mais fácil de ler e compreender

- Ortogonalidade: Capacidade de combinar um conjunto de estruturas primitivas que podem ser combinadas de forma a criar uma estrutura mais complexa.
    - Go é altamente ortogonal. 
    - Regras consistentes e poucas exceções
    - Uso de structs, interfaces, goroutines e channels

- Instruções de Controle:
    - Fluxo de execução simples com if, for e switch
    - Código mais organizado e previsível
    - Formatação automática com gofmt

- Tipos e Estruturas de Dados:
    - Go é uma linguagem fortemente tipada, o que aumenta a clareza sobre o tipo de dado que está sendo manipulado.
    - Oferece tipos primitivos claros (inteiros, floats, booleanos, strings) e estruturas de dados como structs, arrays, slices e maps que são intuitivas.
    - O uso de interfaces permite abstrações claras sem a complexidade da hierarquia de classes.

- Sintaxe:
    - Declarações simples e objetivas
    - Palavras-chave de fácil compreensão
    - Código padronizado automaticamente

| Aspecto                 | C                           | Go                         |
|  :-------------------:  |  :-----------------------:  |  :----------------------:  |
| Estruturas de repetição | `for`, `while` e `do-while` | Apenas `for`               |
| Formatação              | Definida pelo programador   | Padronizada pelo **gofmt** |
| Sintaxe                 | Mais flexível e detalhada   | Mais simples e uniforme    |
| Organização             | Maior liberdade de estilo   | Convenções padronizadas    |

Exemplo:
```go
package main

import "fmt"

type Ponto struct {
    X, Y int
}

func main() {
    nome := "GoLang"
    versao := 1.22

    p := Ponto{X: 10, Y: 20}

    fmt.Printf("%s v%.2f, Ponto: (%d, %d)\n", nome, versao, p.X, p.Y)
}
```

```c
#include <stdio.h>

typedef struct {
    int X;
    int Y;
} Ponto;

int main() {
    char nome[] = "C Lang";
    float versao = 99.0;

    Ponto p = {10, 20};

    printf("%s v%.2f, Ponto: (%d, %d)\n", nome, versao, p.X, p.Y);

    return 0;
}
```

### Capacidade de Escrita
A capacidade de escrita mede a facilidade com que uma linguagem permite criar programas para um determinado domínio.

- **Expressividade**
    - Declaração simplificada de variáveis (:=)
    - Múltiplos valores de retorno para resultados e erros
    - Concorrência nativa com goroutines e channels

- **Abstração**
    - Funções como principal mecanismo de modularização
    - Uso de structs para representar dados
    - Interfaces para abstração e desacoplamento

| Aspecto                 | C                              | Go                            |
|  :-------------------:  |  :--------------------------:  |  :--------------------------: |
| Declaração de variáveis | Mais verbosa                   | Sintaxe simplificada (`:=`)   |
| Retorno de funções      | Geralmente um único valor      | Múltiplos valores de retorno  |
| Concorrência            | Threads e bibliotecas externas | Goroutines e channels nativos |
| Abstração               | Structs e funções              | Structs, funções e interfaces |

```go
package main

import "fmt"

func DividirComResto(a, b int) (int, int) {
    quociente := a / b
    resto := a % b
    return quociente, resto
}

func main() {
    num, den := 17, 5

    q, r := DividirComResto(num, den)

    fmt.Printf("%d dividido por %d é %d com resto %d\n", num, den, q, r)
}
```

```c
#include <stdio.h>

void DividirComResto(int a, int b, int *quociente_ptr, int *resto_ptr) {
    *quociente_ptr = a / b;
    *resto_ptr = a % b;     
}

int main() {
    int num = 17, den = 5;
    int q_c, r_c;
    
    DividirComResto(num, den, &q_c, &r_c);

    printf("%d dividido por %d é %d com resto %d\n", num, den, q_c, r_c);

    return 0;
}
```

###  Confiabilidade
Um programa é confiável se funciona conforme o esperado em todas as condições. A confiabilidade é crucial para sistemas que precisam operar sem falhas, especialmente em ambientes de produção.

- Verificação de Tipos
    - Linguagem fortemente e estaticamente tipada
    - Erros de tipo identificados durante a compilação
    - Inferência de tipos sem perda de segurança

- Tratamento Explícito de Erros:
    - Retorno de erros por meio do tipo error
    - Desenvolvedor deve tratar possíveis falhas
    - Uso de panic apenas para situações excepcionais

- Segurança no Uso de Ponteiros
    - Possui ponteiros, mas sem aritmética de ponteiros
    - Evita acessos inválidos à memória
    - Garbage Collector aumenta a segurança da aplicação

| Aspecto                  | C                            | Go                             |
|  :---------------------: | :-------------------------:  |  :---------------------------: |
| Verificação de tipos     | Estática, porém menos segura | Estática e fortemente tipada   |
| Tratamento de erros      | Códigos de retorno           | Tipo `error` explícito         |
| Ponteiros                | Permite aritmética           | Não permite aritmética         |
| Gerenciamento de memória | Manual                       | Automático (Garbage Collector) |

```go
package main

import "fmt"

func CriarESomar(a, b int) *int {    
    resultado := a + b
    return &resultado 
}

func main() {
    ptr := CriarESomar(5, 3)
    fmt.Println("Soma:", *ptr)
}
```

- Em Go, não é possível criar um "wild pointer" por aritmética de ponteiros ou acessar memória liberada, pois o Garbage Colletor cuida da liberação.
- Não há "memory leaks" acidentais por esquecimento de free() pois não há necessidade de liberar explicitamente.

```c
#include <stdio.h>
#include <stdlib.h>

int* CriarESomar(int a, int b) {
    int* resultado_ptr = (int*)malloc(sizeof(int));
    if (resultado_ptr == NULL) {
        // Mensagem de erro de alocação
        return NULL;
    }
    *resultado_ptr = a + b;
    return resultado_ptr;
}

int main() {
    // Memory Leak
    int* ptr_c = CriarESomar(5, 3);
    if (ptr_c != NULL) {
        printf("Soma: %d\n", *ptr_c);
        /* free(ptr_c); */ 
        // Se esta linha for esquecida, ocorre um memory leak após esquecimento de liberar a memória
    }

    // Wild Pointer: Acesso a memória após liberação
    int* wild_ptr = (int*)malloc(sizeof(int));
    *wild_ptr = 100;
    free(wild_ptr);
    
    *wild_ptr = 200; // Acessar memória após free() (wild pointer)
    printf("Valor do wild pointer: %d\n", *wild_ptr); // Comportamento indefinido
    
    return 0;
}
```

### Custo
O custo total de uma linguagem de programação abrange desde o treinamento da equipe até a manutenção do software em produção, incluindo o impacto de falhas.

Go foi projetada para ser uma linguagem de baixo custo total, otimizando diversos fatores:
- Treinamento:
    - Sintaxe simples e poucas palavras-chave
    - Curva de aprendizado reduzida para desenvolvedores
    - Recursos padronizados e fáceis de compreender

- Escrita e Compilação:
    - Código escrito de forma objetiva e produtiva.
    - Compilação extremamente rápida.
    - Execução em código nativo com alto desempenho.

- Manutenção:
    - Código legível e padronizado pelo `gofmt`.
    - Estrutura simples, facilitando alterações futuras.
    - Menor esforço para corrigir bugs e adicionar funcionalidades

| Aspecto |  C  |  Go |
| :-----: | :-: | :-: |
| Treinamento              | Maior complexidade      | Curva de aprendizado menor   |
| Compilação               | Rápida                  | Muito rápida                 |
| Execução                 | Muito alta              | Muito alta (próxima ao C)    |
| Gerenciamento de memória | Manual                  | Automático (GC)              |
| Manutenção               | Mais suscetível a erros | Código mais simples e seguro 

### Portabilidade
A portabilidade refere-se à facilidade com que um programa pode ser movido e executado em diferentes ambientes de hardware e sistemas operacionais, sem a necessidade de grandes modificações.

- Compilação para Código Nativo
    - Geração de executáveis nativos para cada plataforma
    - Não depende de máquina virtual ou interpretador
    - Alto desempenho em diferentes sistemas operacionai

- Compilação Cruzada (Cross-Compilation):
    - Compilação para diferentes sistemas operacionais
    - Suporte a múltiplas arquiteturas, como x86 e ARM
    - Processo simples com poucas configurações.

- Executáveis Autônomos
    - Binários estáticos com bibliotecas incorporadas
    - Pouca dependência de componentes externos
    - Implantação simples e rápida

| Aspecto            | C                                       | Go                               |
|  :---------------: |  :-----------------------------------:  |  :----------------------------:  |
| Compilação         | Nativa                                  | Nativa                           |
| Compilação cruzada | Exige ferramentas específicas           | Suporte nativo                   |
| Dependências       | Pode depender de bibliotecas do sistema | Executáveis estáticos por padrão |
| Portabilidade      | Depende do compilador e das bibliotecas | Elevada e padronizada            |

### Generalidade
A generalidade (ou aplicabilidade) de uma linguagem refere-se à sua capacidade de ser utilizada em uma ampla variedade de domínios de aplicação, sem ser excessivamente especializada.

Como já foi dito anteriomente, Go é uma linguagem de propósito geral, projetada para ser eficaz em diversos cenários:
- Desenvolvimento de Sistemas
- Ferramentas de Rede:
- DevOps e Ferramentas de Infraestrutura: Docker e Kubernets
- Computação em Nuvem:
- Outros Domínios: Machine Learning, Sistemas Embarcados

| Aspecto                       | C                             | Go                                     |
|  :--------------------------: |  :--------------------------: |  :----------------------------------:  |
| Propósito                     | Geral, com foco em sistemas   | Geral, com foco em sistemas e serviços |
| Desenvolvimento Web           | Pouco utilizado               | Muito utilizado                        |
| Ferramentas de Infraestrutura | Possível, porém mais complexo | Um dos principais domínios             |
| Computação em Nuvem           | Uso limitado                  | Amplamente utilizada                   |


# O Projeto
Sistema de Busca e Indexação Distribuído
- Backend desenvolvido em Go
- Coleta, processamento e indexação de documentos
- Suporte a sites e arquivos PDF

- Coleta e Processamento Concorrente
    - Uso de goroutines para processar múltiplas fontes simultaneamente
    - Maior desempenho na coleta de documentos
    - Redução do tempo de processamento

- Processamento Paralelo
    - Parsing e indexação executados em paralelo
    - Melhor aproveitamento dos múltiplos núcleos do processador
    - Maior velocidade na construção do índice

- Sincronização
    - Comunicação entre goroutines utilizando channels
    - Proteção do índice compartilhado com sync.RWMutex

## Por que Go?
- Concorrência nativa: Uso de goroutines para executar múltiplas tarefas simultaneamente
- Comunicação eficiente: Channels para troca segura de informações entre goroutines
- Processamento paralelo: Melhor aproveitamento de processadores multicore
- Sincronização de dados: Utilização de sync.RWMutex para acesso seguro ao índice compartilhado
- Alto desempenho: Compilação para código nativo e baixo consumo de recursos

## Como Funciona o Projeto?

O sistema funciona em duas etapas principais:

- **Indexação:**
    - Coleta documentos de sites e arquivos PDF
    - Extrai e normaliza palavras
    - Armazena as informações em um índice invertido 

- **Busca:**
    - Consulta o índice local ou distribuído
    - Calcula relevância utilizando TF-IDF
    - Retorna os documentos ordenados por relevância

## Principais Estruturas
- `Indexer`:
  - Gerencia todo o índice de documentos
  - Mantém índice invertido, frequência de termos e documentos
  - Protegido por `sync.RWMutex`

- `InvertedIndex`: Relaciona cada palavra aos documentos onde aparece
- `DocumentFrequency`: Armazena em quantos documentos cada termo está presente

```go
type InvertedIndex map[string]map[string]int

type DocumentFrequency map[string]int

type Indexer struct {
    Index     InvertedIndex
    DocFreq   DocumentFrequency
    Documents []string
    mu        sync.RWMutex
    name      string       
}
```

### Inicialização do Sistema
- `NewIndexer()`: Cria e inicializa uma nova instância do indexador.
- `createEmptyIndex()` : Inicializa as estruturas de dados do índice vazias.
- `loadIndex()`: Carrega um índice previamente salvo do disco para a memória.

```go
func NewIndexer(name string) *Indexer {
    idx := &Indexer{
        name: name,
    }
    idx.createEmptyIndex()

    if err := idx.loadIndex(); err != nil {
        if errors.Is(err, os.ErrNotExist) {
            log.Printf("Arquivo de índice '%s.db' não encontrado. Criando um novo índice vazio.", name)
        } else {
            log.Printf("Erro ao carregar o índice de '%s.db': %v. Iniciando com um índice vazio.", name, err)
        }
    } else {
        log.Printf("Índice carregado com sucesso de '%s.db'.", name)
    }

    return idx
}

func (idx *Indexer) createEmptyIndex() {
    idx.Index = make(InvertedIndex)
    idx.DocFreq = make(DocumentFrequency)
    idx.Documents = make([]string, 0)
}

func (idx *Indexer) loadIndex() error {
    indexFilePath := fmt.Sprintf("/tmp/%s.db", idx.name)
    if _, err := os.Stat(indexFilePath); errors.Is(err, os.ErrNotExist) {
        return os.ErrNotExist
    }

    var savedIndex SavedIndex
    err := idx.readStructFromFile(indexFilePath, &savedIndex)
    if err != nil {
        return fmt.Errorf("erro ao ler o índice do arquivo: %w", err)
    }

    idx.Index = savedIndex.Index
    idx.DocFreq = savedIndex.DocFreq
    idx.Documents = savedIndex.Documents
    return nil
}
```

## Processo de Indexação
- `AddDocToIndex()`: Processa um documento e adiciona suas palavras ao índice invertido.
- `SaveIndex()`: Salva o estado atual do índice em um arquivo no disco.
- `writeStructToFile()`: Serializa as estruturas de dados e grava o índice em arquivo.

```go
func (idx *Indexer) AddDocToIndex(url string, content string) {
    idx.mu.Lock() 
    defer idx.mu.Unlock()

    idx.Documents = append(idx.Documents, url)

    reader := strings.NewReader(content)
    seenTerms := make(map[string]bool)
    scanner := bufio.NewScanner(reader)
    scanner.Split(bufio.ScanWords)

    for scanner.Scan() {
        word := strings.ToLower(strings.Trim(scanner.Text(), ",.!?&<>;:=§$%{}[]()|"))

        if idx.Index[word] == nil {
            idx.Index[word] = make(map[string]int)
        }
        idx.Index[word][url]++

        if !seenTerms[word] {
            idx.DocFreq[word]++
            seenTerms[word] = true
        }
    }
}

func (idx *Indexer) SaveIndex() error {
    idx.mu.RLock()
    defer idx.mu.RUnlock()

    savedIndex := SavedIndex{
        Index:     idx.Index,
        DocFreq:   idx.DocFreq,
        Documents: idx.Documents,
    }

    indexFilePath := fmt.Sprintf("/tmp/%s.db", idx.name)
    err := idx.writeStructToFile(indexFilePath, savedIndex)
    if err != nil {
        return fmt.Errorf("erro ao salvar o índice: %w", err)
    }
    log.Printf("Índice salvo com sucesso em '%s'.", indexFilePath)
    return nil
}

func (idx *Indexer) writeStructToFile(filename string, data interface{}) error {
    var buf bytes.Buffer
    enc := gob.NewEncoder(&buf)
    err := enc.Encode(data)
    if err != nil {
        return fmt.Errorf("erro ao serializar as structs: %w", err)
    }

    err = os.WriteFile(filename, buf.Bytes(), 0644)
    if err != nil {
        return fmt.Errorf("erro ao gravar o arquivo '%s': %w", filename, err)
    }
    return nil
}
```

## Recuperação do Índice
- `loadIndex()`: Carrega o índice salvo do disco para a memória da aplicação.
- `readStructFromFile()`: Lê e desserializa os dados armazenados no arquivo de índice.

```go
func (idx *Indexer) loadIndex() error {
    indexFilePath := fmt.Sprintf("/tmp/%s.db", idx.name)
    if _, err := os.Stat(indexFilePath); errors.Is(err, os.ErrNotExist) {
        return os.ErrNotExist
    }

    var savedIndex SavedIndex
    err := idx.readStructFromFile(indexFilePath, &savedIndex)
    if err != nil {
        return fmt.Errorf("erro ao ler o índice do arquivo: %w", err)
    }

    idx.Index = savedIndex.Index
    idx.DocFreq = savedIndex.DocFreq
    idx.Documents = savedIndex.Documents
    return nil
}

func (idx *Indexer) readStructFromFile(filename string, data interface{}) error {
    content, err := os.ReadFile(filename)
    if err != nil {
        return fmt.Errorf("erro ao ler o arquivo '%s': %w", filename, err)
    }

    buf := bytes.NewBuffer(content)
    dec := gob.NewDecoder(buf)
    err = dec.Decode(data)
    if err != nil {
        return fmt.Errorf("erro ao desserializar as structs do arquivo '%s': %w", filename, err)
    }
    return nil
}
```

## Fluxo da Busca
- `SearchQuery()`: Recebe a consulta do usuário e coordena a execução da busca.
- `Search()`: Localiza os documentos que correspondem aos termos pesquisados.
- `scoreDoc()`: Calcula a relevância de cada documento utilizando o algoritmo TF-IDF.

```go
func SearchQuery(c *fiber.Ctx) error {
	query := c.Query("q")
	queryType := c.Query("type", "SIMPLE")
	distributed := c.Query("dist") == "true"

	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Busca vazia"})
	}

	var results []search.SearchResult

	if distributed {
		results = AggregateDistributedResults(query)
	} else {
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

func scoreDoc(terms []string, doc string, idx index.InvertedIndex, docFreq index.DocumentFrequency, numDocs int) float64 {
	score := 0.0
	for _, term := range terms {
		tf := float64(idx[term][doc])
		idf := math.Log(float64(numDocs) / float64(docFreq[term]))
		score += tf * idf
	}
	return score
}
```

## Busca Distribuída
- `AggregateDistributedResults()`: Coordena a busca em todos os nós e reúne os resultados obtidos.
- `SearchInRemoteNode()`: Envia a consulta para um nó remoto e recupera seus resultados.
- `RankDistributed()`: Consolida e ordena os resultados de todos os nós por relevância.

```go
func AggregateDistributedResults(query string) []search.SearchResult {
	resultsChan := make(chan []search.SearchResult)

	for _, url := range remoteNodes {
		go func(nodeURL string) {
			res := search.SearchInRemoteNode(nodeURL, query)
			resultsChan <- res
		}(url)
	}

	var allResults []search.SearchResult

	for i := 0; i < len(remoteNodes); i++ {
		nodeRes := <-resultsChan
		allResults = append(allResults, nodeRes...)
	}

	return search.RankDistributed(allResults)
}

func SearchInRemoteNode(serverURL string, query string) []SearchResult {
	fullURL := fmt.Sprintf("%s/search?q=%s", serverURL, query)
	resp, err := http.Get(fullURL)
	
    if err != nil {
		return []SearchResult{}
	}
	
	defer resp.Body.Close()

	var results []SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return []SearchResult{}
	}

	return results
}

func RankDistributed(allResults []SearchResult) []SearchResult {
	globalScores := make(map[string]float64)
	
	for _, res := range allResults {
		globalScores[res.Document] += res.Score
	}

	return rankResults(globalScores)
}
```

### Inicialização da Aplicação
- `main()`
- Configuração das rotas
- Inicialização do servidor Fiber

```go
package main

import (
    "fmt"
    "log"
    "github.com/gabrielbismarck/GoMaker/internal/controller"
    "github.com/gabrielbismarck/GoMaker/pkg/index"
    "github.com/gofiber/fiber/v2"
)

func main() {

    myIndexer := index.NewIndexer("default")

	controller.SetIndexer(myIndexer)

	app := fiber.New()

	app.Post("/index", controller.AddDocumentToIndex)

	app.Post("/save-index", func(c *fiber.Ctx) error {
		if err := myIndexer.SaveIndex(); err != nil {
			log.Printf("Erro ao salvar o índice: %v", err)
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Erro ao salvar o índice: %v", err))
		}
		return c.SendString("Índice salvo com sucesso!")
	})

	log.Fatal(app.Listen(":3000"))

}
```

## Referências
- Go official website: https://go.dev/
- CTII418- Linguagem Go (IFSP): https://cbt.ifsp.edu.br/images/Documentos/2021/CTII/CTII418_Go.pdf
- COX-BUDAY, Katherine. *Concurrency in Go*. O'Reilly Media, 2017.
- DOXSEY, Caleb. *Introducing Go*. O'Reilly Media, 2016.

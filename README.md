# Sistema de Aluguel de Equipamentos — Go

Projeto acadêmico: demonstrar as premissas da linguagem Go por meio de um sistema backend real, comparando implementações com e sem controle de concorrência.

## Critérios de Avaliação

Lista baseada na planilha disponibilizada pelo professor.

- [ X ] Linguagem: histórico e versão
- [ X ] Projeto: premissas, usuário característico, domínio de aplicação
- [ X ] Construtores (com exemplos)
- [ X ] Legibilidade (fatores, exemplos e conclusão)
- [ X ] Capacidade de escrita (fatores, exemplos e conclusão)
- [ X ] Confiabilidade (fatores, exemplos e conclusão)
- [ X ]  Custo e outros critérios de avaliação da linguagem
- [ ] Projeto: apresentação e explicação do código
- [ ] Site / Demonstração / Wiki (Colocar o que cada integrante fez)
- [ ] Vídeo (duração: ~10 min)

## Slides

### Slide - Capa

### Slide - Histórico e Versão
-  **Go** ou **Golang** foi criada no Google por ***Robert Griesemer***, ***Rob Pike*** e ***Ken Thompson***.
- A primeira versão estável (Go 1.0) foi lançada em março de 2012, e a versão atual é a **1.24**, lançada em 2025.
- **Motivação:** Os criadores observaram que as linguagens existentes na época (C++, Java, Python) não estavam atendendo bem às necessidades de desenvolvimento de software em larga escala no Google.
- Dentre as exigências buscavasse:
    - **Eficiência** e **performance** de linguagens compiladas (como C/C++).
    - **Simplicidade** e **facilidade** de uso de linguagens dinâmicas (como Python).
    - **Concorrência nativa** e eficiente para aproveitar os processadores multi-core modernos e lidar com sistemas distribuídos.

### Slide - Premissas da Linguagem 
- **Performance e Eficiência:**
    - Binários nativos e de alta performance.
    - Compilação rápida.
    - Coletor de lixo eficiente.
    
> Explicação: Go foi projetada para oferecer desempenho próximo ao C/C++, gerando binários nativos e com tempos de compilação muito rápidos.

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

### Slide - Dominíos de Aplicação
- Programação de Sistemas
    - Desenvolvimento de ferramentas de infraestrutura e utilitários
    - Alto desempenho e eficiência para serviços em nuvem
    - Exemplos: Docker e Kubernetes

- Aplicações Comerciais (Backend)
    - Desenvolvimento de APIs REST, microsserviços e aplicações web
    - Capacidade de processar grande volume de requisições simultâneas
    - Exemplos: Uber, Netflix, Twitch e Google

### Slide - Usuário Característico - Não sei o que colocar aqui
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

### Slide - Construtores

Go oferece um conjunto de construtores fundamentais para a construção de software eficiente e robusto. 

Muitos têm paralelos em C, mas com características próprias que visam simplicidade e segurança.

- Tipos Estruturados (struct):
    - Go: Agrupa campos de diferentes tipos em uma única unidadeBase para modelar dados complexos.
    - C: Similar ao struct em C
    
> Diferença é que em Go, structs podem ter métodos associados e são usadas diretamente como tipos.

- Funções (func):
    - Go: Blocos de código reutilizáveis que executam uma tarefa específica. Podem retornar múltiplos valores.
    - C: Similar às funções em C. 
    
> A principal diferença em Go é o suporte nativo a múltiplos valores de retorno, o que simplifica o tratamento de erros.

- Variáveis e Constantes (var, const):
    - Go: Declaração de armazenamento para valores. Go exige que variáveis declaradas sejam usadas, evitando código morto.
    - C: Similar à declaração de variáveis e constantes.

- Operador de Declaração Curta (:=):
    -  Go: Não existe um equivalente direto em C. Declara e inicializa uma variável, inferindo seu tipo automaticamente. Usado para declarações locais.
    
> Explicação: É uma forma concisa e idiomática de declarar e atribuir um valor a uma nova variável em Go, muito utilizada no dia a dia.

- Declarações de Controle de Fluxo (if, for, switch):
    - Go: Permitem controlar a ordem de execução do código.
    - C: Similar às estruturas de controle de C. Go não possui o while (usa for para isso) e não exige parênteses nas condições.

- Ponteiros (*, &):
    - Go: Permitem manipular diretamente endereços de memória.
    - C: Similar aos ponteiros em C. 
    
> Diferença: Go não possui aritmética de ponteiros, o que aumenta a segurança e reduz a complexidade.

- Interfaces (interface):
    - Go: Não existe um equivalente direto em C. Define um conjunto de métodos que um tipo deve implementar. Permite polimorfismo e design flexível, desacoplando a implementação da interface.
    
> Explicação: Em Go, um tipo implementa uma interface implicitamente ao ter todos os métodos da interface.

- Goroutines (go):
    - Go: Não existe um equivalente direto em C. Funções leves e concorrentes que rodam em paralelo. Gerenciadas pelo runtime de Go.

> Explicação: Permitem escrever código concorrente de forma simples e eficiente, sem a complexidade de threads de sistema operacional.

- Canais (chan):
    - Go: Não existe um equivalente direto em C. Meio de comunicação seguro entre goroutines. Permitem que goroutines enviem e recebam valores umas das outras.

>  Explicação: Implementam o modelo CSP (Communicating Sequential Processes), onde a comunicação é a forma preferida de compartilhar memória.

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

### Slide - Legibilidade
Legibilidade é a facilidade com que um programa pode ser lido e entendido.

- Simplicidade Global
    - Sintaxe enxuta e poucas palavras-chave
    - Apenas um tipo de laço (for)
    - Código mais fácil de ler e compreender

> Falar: "A legibilidade de Go começa pela sua simplicidade. A linguagem possui poucas palavras-chave e evita oferecer várias maneiras de fazer a mesma coisa. Um exemplo é que existe apenas um tipo de laço, o for, substituindo estruturas como while e do-while. Isso torna o código mais uniforme e facilita sua leitura."

- Ortogonalidade: Capacidade de combinar um conjunto de estruturas primitivas que podem ser combinadas de forma a criar uma estrutura mais complexa.
    - Go é altamente ortogonal. 
    - Regras consistentes e poucas exceções
    - Uso de structs, interfaces, goroutines e channels

> Falar: "Outro aspecto importante é a ortogonalidade. Em Go, um pequeno conjunto de recursos pode ser combinado para resolver diferentes problemas. Como as regras da linguagem possuem poucas exceções, o comportamento do código se torna previsível e mais fácil de entender."

- Instruções de Controle:
    - Fluxo de execução simples com if, for e switch
    - Código mais organizado e previsível
    - Formatação automática com gofmt

> Falar: "Go também favorece a legibilidade por meio de suas estruturas de controle. O fluxo do programa é simples e organizado, utilizando principalmente if, for e switch. Além disso, a ferramenta gofmt formata automaticamente todo o código, garantindo um padrão único entre diferentes projetos e desenvolvedores."

- Tipos e Estruturas de Dados:
    - Go é uma linguagem fortemente tipada, o que aumenta a clareza sobre o tipo de dado que está sendo manipulado.
    - Oferece tipos primitivos claros (inteiros, floats, booleanos, strings) e estruturas de dados como structs, arrays, slices e maps que são intuitivas.
    - O uso de interfaces permite abstrações claras sem a complexidade da hierarquia de classes.

> Falar: "A tipagem forte facilita a identificação do tipo de cada variável e reduz ambiguidades. Go também oferece estruturas de dados intuitivas, como structs, slices e maps, além de interfaces para abstração. Isso torna o código mais organizado e fácil de interpretar."

- Sintaxe:
    - Declarações simples e objetivas
    - Palavras-chave de fácil compreensão
    - Código padronizado automaticamente

> Falar: "Por fim, a própria sintaxe contribui para a legibilidade. As declarações de funções, variáveis e constantes são diretas e objetivas. Como toda a equipe utiliza o gofmt, praticamente todo código Go segue o mesmo estilo, permitindo que qualquer desenvolvedor consiga ler e compreender um projeto com facilidade."

| Aspecto                 | C                           | Go                         |
|  :-------------------:  |  :-----------------------:  |  :----------------------:  |
| Estruturas de repetição | `for`, `while` e `do-while` | Apenas `for`               |
| Formatação              | Definida pelo programador   | Padronizada pelo **gofmt** |
| Sintaxe                 | Mais flexível e detalhada   | Mais simples e uniforme    |
| Organização             | Maior liberdade de estilo   | Convenções padronizadas    |

> Falar: "Comparando com C, Go prioriza a padronização e a simplicidade. Enquanto C oferece mais liberdade de escrita, isso pode fazer com que diferentes programadores adotem estilos muito distintos. Em Go, a sintaxe enxuta e a padronização promovida pelo gofmt tornam o código mais consistente, facilitando sua leitura, manutenção e colaboração em equipe."

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

### Slide - Capacidade de Escrita
A capacidade de escrita mede a facilidade com que uma linguagem permite criar programas para um determinado domínio.

- **Expressividade**
    - Declaração simplificada de variáveis (:=)
    - Múltiplos valores de retorno para resultados e erros
    - Concorrência nativa com goroutines e channels

> Falar: "A capacidade de escrita está relacionada à facilidade de desenvolver programas utilizando a linguagem. Em Go, isso é alcançado por uma sintaxe simples e objetiva. A declaração curta de variáveis reduz a quantidade de código, enquanto os múltiplos valores de retorno facilitam o tratamento de erros. Outro grande diferencial é a concorrência nativa, que permite executar tarefas em paralelo utilizando goroutines e channels de maneira muito mais simples do que em outras linguagens."

- **Abstração**
    - Funções como principal mecanismo de modularização
    - Uso de structs para representar dados
    - Interfaces para abstração e desacoplamento

> Falar: "Go também oferece mecanismos eficientes de abstração. As funções permitem dividir o programa em partes menores e reutilizáveis. Já as structs organizam os dados de forma simples, enquanto as interfaces permitem definir comportamentos sem depender de implementações específicas. Isso deixa o código mais organizado, flexível e fácil de manter."

| Aspecto                 | C                              | Go                            |
|  :-------------------:  |  :--------------------------:  |  :--------------------------: |
| Declaração de variáveis | Mais verbosa                   | Sintaxe simplificada (`:=`)   |
| Retorno de funções      | Geralmente um único valor      | Múltiplos valores de retorno  |
| Concorrência            | Threads e bibliotecas externas | Goroutines e channels nativos |
| Abstração               | Structs e funções              | Structs, funções e interfaces |

> Falar: "Comparando com C, Go oferece recursos que tornam o desenvolvimento mais produtivo. Enquanto em C o programador normalmente precisa recorrer a bibliotecas para trabalhar com concorrência e possui mecanismos de abstração mais limitados, Go já oferece essas funcionalidades de forma nativa. O resultado é um código mais limpo, modular e fácil de desenvolver, sem perder desempenho."

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

### Slide -  Confiabilidade
Um programa é confiável se funciona conforme o esperado em todas as condições. A confiabilidade é crucial para sistemas que precisam operar sem falhas, especialmente em ambientes de produção.

- Verificação de Tipos
    - Linguagem fortemente e estaticamente tipada
    - Erros de tipo identificados durante a compilação
    - Inferência de tipos sem perda de segurança

> Falar: "A primeira característica que aumenta a confiabilidade de Go é a verificação de tipos. Como ela é uma linguagem fortemente e estaticamente tipada, muitos erros são detectados ainda na compilação, antes mesmo da execução do programa. Isso reduz bastante a chance de falhas em produção. Além disso, Go permite inferência de tipos usando o operador :=, o que deixa o código mais simples sem comprometer a segurança."

- Tratamento Explícito de Erros:
    - Retorno de erros por meio do tipo error
    - Desenvolvedor deve tratar possíveis falhas
    - Uso de panic apenas para situações excepcionais

> Falar: "Diferente de linguagens que utilizam exceções com try e catch, Go trabalha com tratamento explícito de erros. Normalmente uma função retorna o resultado e um valor do tipo error. Isso obriga o desenvolvedor a verificar se ocorreu algum problema antes de continuar a execução, tornando o código mais previsível e reduzindo erros que poderiam passar despercebidos."

- Segurança no Uso de Ponteiros
    - Possui ponteiros, mas sem aritmética de ponteiros
    - Evita acessos inválidos à memória
    - Garbage Collector aumenta a segurança da aplicação

> Falar: "Go também oferece maior segurança no uso de ponteiros. A linguagem permite utilizá-los, porém não permite aritmética de ponteiros, que é uma das principais causas de erros em linguagens como C. Além disso, o Garbage Collector gerencia automaticamente a memória, reduzindo problemas como vazamentos e acessos inválidos."

| Aspecto                  | C                            | Go                             |
|  :---------------------: | :-------------------------:  |  :---------------------------: |
| Verificação de tipos     | Estática, porém menos segura | Estática e fortemente tipada   |
| Tratamento de erros      | Códigos de retorno           | Tipo `error` explícito         |
| Ponteiros                | Permite aritmética           | Não permite aritmética         |
| Gerenciamento de memória | Manual                       | Automático (Garbage Collector) |

> Falar: "Comparando com C, Go foi projetada para reduzir erros comuns de programação. Enquanto em C o programador é responsável por gerenciar toda a memória e pode manipular ponteiros livremente, em Go essas operações são mais controladas. Isso torna a linguagem mais segura e confiável, especialmente em aplicações grandes e sistemas distribuídos, onde uma falha pode causar impactos significativos."

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

### Slide - Custo
O custo total de uma linguagem de programação abrange desde o treinamento da equipe até a manutenção do software em produção, incluindo o impacto de falhas.

Go foi projetada para ser uma linguagem de baixo custo total, otimizando diversos fatores:
- Treinamento:
    - Sintaxe simples e poucas palavras-chave
    - Curva de aprendizado reduzida para desenvolvedores
    - Recursos padronizados e fáceis de compreender

> Explicação: "Um dos objetivos da linguagem Go foi reduzir o custo de treinamento dos desenvolvedores. Ela possui uma sintaxe simples, poucas palavras-chave e evita recursos muito complexos presentes em outras linguagens. Por isso, quem já possui experiência com linguagens como C, C++ ou Java consegue aprender Go de forma relativamente rápida, reduzindo o tempo necessário para que uma equipe se torne produtiva."        


- Escrita e Compilação:
    - Código escrito de forma objetiva e produtiva.
    - Compilação extremamente rápida.
    - Execução em código nativo com alto desempenho.

> Explicação: "Outro fator importante é o custo durante o desenvolvimento. Go possui uma sintaxe objetiva, permitindo escrever programas com menos código e de forma mais organizada. Além disso, seu compilador é extremamente rápido, diminuindo o tempo entre escrever, compilar e testar a aplicação. Mesmo com essa velocidade de compilação, os programas são convertidos para código de máquina nativo, alcançando um desempenho muito próximo ao obtido em C e C++."

- Manutenção:
    - Código legível e padronizado pelo `gofmt`.
    - Estrutura simples, facilitando alterações futuras.
    - Menor esforço para corrigir bugs e adicionar funcionalidades

> Explicação: "Por fim, Go foi projetada para facilitar a manutenção do software. Sua sintaxe é simples, o código possui um padrão de formatação garantido pela ferramenta gofmt e a linguagem evita recursos que tornam os programas difíceis de entender. Isso faz com que qualquer desenvolvedor consiga ler, modificar e evoluir o código com mais facilidade, diminuindo o custo de manutenção ao longo do ciclo de vida do sistema."

| Aspecto |  C  |  Go |
| :-----: | :-: | :-: |
| Treinamento              | Maior complexidade      | Curva de aprendizado menor   |
| Compilação               | Rápida                  | Muito rápida                 |
| Execução                 | Muito alta              | Muito alta (próxima ao C)    |
| Gerenciamento de memória | Manual                  | Automático (GC)              |
| Manutenção               | Mais suscetível a erros | Código mais simples e seguro 

> Explicação: "Em comparação com a linguagem C, Go oferece uma curva de aprendizado menor, gerenciamento automático de memória e maior facilidade de manutenção. Embora C ainda apresente um desempenho ligeiramente superior em aplicações extremamente críticas, Go mantém uma performance muito próxima, oferecendo em troca mais produtividade, segurança e menor custo de desenvolvimento. Por esse motivo, ela se tornou uma das principais escolhas para aplicações de backend, computação em nuvem e sistemas distribuídos."

### Slide - Outros Critérios de Avaliação da Linguagem: Portabilidade
A portabilidade refere-se à facilidade com que um programa pode ser movido e executado em diferentes ambientes de hardware e sistemas operacionais, sem a necessidade de grandes modificações.

- Compilação para Código Nativo
    - Geração de executáveis nativos para cada plataforma
    - Não depende de máquina virtual ou interpretador
    - Alto desempenho em diferentes sistemas operacionai

> Falar: "A portabilidade de Go começa pela forma como a linguagem é compilada. Diferente de linguagens que dependem de uma máquina virtual, como Java, Go gera executáveis nativos para cada sistema operacional. Isso significa que o programa pode ser executado diretamente, oferecendo melhor desempenho e facilitando sua distribuição."


- Compilação Cruzada (Cross-Compilation):
    - Compilação para diferentes sistemas operacionais
    - Suporte a múltiplas arquiteturas, como x86 e ARM
    - Processo simples com poucas configuraçõeso.

> Falar: "Um dos grandes diferenciais de Go é o suporte à compilação cruzada. Com um único comando é possível gerar executáveis para Windows, Linux, macOS e até diferentes arquiteturas de processadores, sem precisar configurar vários ambientes de desenvolvimento. Isso facilita bastante a distribuição de aplicações."

- Executáveis Autônomos
    - Binários estáticos com bibliotecas incorporadas
    - Pouca dependência de componentes externos
    - Implantação simples e rápida

> Falar: "Outra vantagem é que os programas Go normalmente são distribuídos como um único executável. As bibliotecas necessárias já ficam incorporadas ao binário, eliminando a necessidade de instalar dependências ou ambientes de execução na máquina de destino. Isso simplifica bastante a implantação."

| Aspecto            | C                                       | Go                               |
|  :---------------: |  :-----------------------------------:  |  :----------------------------:  |
| Compilação         | Nativa                                  | Nativa                           |
| Compilação cruzada | Exige ferramentas específicas           | Suporte nativo                   |
| Dependências       | Pode depender de bibliotecas do sistema | Executáveis estáticos por padrão |
| Portabilidade      | Depende do compilador e das bibliotecas | Elevada e padronizada            |

> Falar: "Comparando com C, as duas linguagens geram código nativo e possuem excelente desempenho. Porém, Go facilita muito a portabilidade ao oferecer compilação cruzada integrada e gerar executáveis praticamente independentes de bibliotecas externas. Isso reduz problemas na hora de distribuir aplicações para diferentes plataformas e torna o processo de implantação muito mais simples."

### Slide - Outros Critérios de Avaliação da Linguagem: Generalidade
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

> Falar: "Comparando com C, ambas são linguagens de propósito geral. Entretanto, Go foi projetada pensando nas necessidades atuais, como computação em nuvem, microsserviços e sistemas distribuídos. Já C continua sendo mais utilizada em sistemas embarcados, sistemas operacionais e aplicações de baixo nível. Por isso, apesar de ambas serem gerais, cada uma se destaca em domínios diferentes."

### Slide - O que é o nosso projeto ?
Sistema de Busca e Indexação Distribuído
- Backend desenvolvido em Go
- Coleta, processamento e indexação de documentos
- Suporte a sites e arquivos PDF

> Falar: "Nosso projeto consiste em um sistema backend desenvolvido em Go para realizar a busca e indexação de conteúdos distribuídos. O sistema é capaz de coletar documentos de diferentes fontes, como páginas da web e arquivos PDF, processar essas informações e criar um índice que permite realizar buscas de forma rápida e eficiente."

- Coleta e Processamento Concorrente
    - Uso de goroutines para processar múltiplas fontes simultaneamente
    - Maior desempenho na coleta de documentos
    - Redução do tempo de processamento

> Falar: "Como os documentos podem vir de várias fontes independentes, utilizamos as goroutines para realizar a coleta de forma concorrente. Enquanto uma goroutine está lendo um site, outra pode estar processando um PDF, permitindo que várias tarefas aconteçam ao mesmo tempo e reduzindo significativamente o tempo de execução."

- Processamento Paralelo
    - Parsing e indexação executados em paralelo
    - Melhor aproveitamento dos múltiplos núcleos do processador
    - Maior velocidade na construção do índice

> Falar: "Depois que os documentos são coletados, eles precisam ser analisados e indexados. Essas tarefas exigem bastante processamento e, por isso, também são executadas em paralelo. Dessa forma, conseguimos aproveitar melhor os processadores multicore e acelerar a criação do índice de busca."

- Sincronização
    - Comunicação entre goroutines utilizando channels
    - Proteção do índice compartilhado com sync.RWMutex

> Falar: "Como várias goroutines trabalham ao mesmo tempo, é necessário sincronizar o acesso aos dados compartilhados. Para isso, utilizamos os channels, responsáveis pela comunicação entre as tarefas, e o sync.RWMutex, que protege o índice durante as atualizações".

### Slide - Por que Go ?

| Requisito do projeto                  | Solução oferecida pelo Go |
|  :---------------------------------:  |  :---------------------:  |
| Coletar várias fontes simultaneamente | Goroutines                |
| Comunicação entre tarefas             | Channels                  |
| Processamento paralelo                | Escalonamento automático  |
| Proteção de dados compartilhados      | `sync.RWMutex`            |

> Falar: "Escolhemos Go porque ela atende exatamente às necessidades do projeto. A linguagem oferece suporte nativo à concorrência, facilita o processamento paralelo e fornece mecanismos seguros para sincronização entre tarefas. Isso permite desenvolver um sistema distribuído eficiente, escalável e com alto desempenho, exatamente o que é necessário para um serviço de busca e indexação."

### Slide - Como Ele funciona ?

O nosso projeto opera em duas fases principais: **Indexação** e **Busca**.

### Indexação
A indexação é o processo de coletar, analisar e armazenar informações de forma estruturada para permitir a recuperação rápida e eficiente de dados. Em nosso projeto, a indexação transforma o conteúdo bruto dos documentos em um formato otimizado para buscas.

**O que é a Indexação?**
- Em termos simples, a indexação cria um "mapa" ou "índice" de todas as palavras encontradas nos documentos, associando cada palavra aos documentos onde ela aparece. 
- Isso é análogo ao índice remissivo de um livro, que lista termos importantes e as páginas onde eles podem ser encontrados. 
- Sem um índice, uma busca exigiria a leitura completa de cada documento, o que seria inviável para grandes volumes de dados.

**Como foi aplicada em nosso Projeto ?**
- Em nosso projeto, a indexação é gerenciada pela `struct Indexer` (localizada em `pkg/index/index.go`). Esta estrutura encapsula todo o estado do índice, incluindo:
    - `Index`: Um mapa invertido que associa cada palavra (token) a uma lista de documentos onde ela ocorre.
    - `DocFreq`: Registra a frequência de cada palavra nos documentos.
    - `Documents`: Armazena os documentos originais (URL e conteúdo) por um ID único.
    - `mu`: Um `sync.RWMutex` para garantir acesso seguro e concorrente ao índice.

### Slide - Explicação do código da parte do index

### Slide - Busca

### Slide - Distríbuida

### Slide - Possíveis Melhorias

### Slide - Referências
- Go official website: https://go.dev/
- CTII418- Linguagem Go (IFSP): https://cbt.ifsp.edu.br/images/Documentos/2021/CTII/CTII418_Go.pdf
- COX-BUDAY, Katherine. *Concurrency in Go*. O'Reilly Media, 2017.
- DOXSEY, Caleb. *Introducing Go*. O'Reilly Media, 2016.

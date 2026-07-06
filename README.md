# Sistema de Aluguel de Equipamentos — Go

Projeto acadêmico: demonstrar as premissas da linguagem Go por meio de um sistema backend real, comparando implementações com e sem controle de concorrência.

## Critérios de Avaliação

Lista baseada na planilha disponibilizada pelo professor.

- [ ] Linguagem: histórico e versão
- [ ] Projeto: premissas, usuário característico, domínio de aplicação
- [ ] Construtores (com exemplos)
- [ ] Legibilidade (fatores, exemplos e conclusão)
- [ ] Capacidade de escrita (fatores, exemplos e conclusão)
- [ ] Confiabilidade (fatores, exemplos e conclusão)
- [ ] Custo e outros critérios de avaliação da linguagem
- [ ] Projeto: apresentação e explicação do código
- [ ] Site / Demonstração / Wiki (Colocar o que cada integrante fez)
- [ ] Vídeo (duração: ~10 min)

## Linguagem - Histórico e Versão

A linguagem Go foi criada em 2007 por Robert Griesemer, Rob Pike e Ken
Thompson como projeto interno do Google, tornou-se open source em novembro
de 2009 e atingiu estabilidade com o lançamento da versão 1.0 em março de
2012. A versão atual é a **1.24**, lançada em 2025. <sources>[1]</sources>

A motivação central foi resolver limitações práticas em escala: ciclos de
compilação lentos com C++, dificuldade de escrever software concorrente seguro
e a ausência de uma linguagem que reunisse tipagem estática, simplicidade de
sintaxe e uso eficiente de hardware multinúcleo. Em 2008 o projeto deixou de
ser atividade paralela e passou a ser desenvolvido em tempo integral dentro da
empresa.

Go se inspira na sintaxe do C, mas incorpora coleta de lixo automática,
inferência de tipos e um modelo de concorrência baseado em goroutines e
channels, derivado do formalismo CSP (Communicating Sequential Processes).

## Linguagem - Domínios de Aplicação

Go se consolidou em domínios onde concorrência, desempenho e simplicidade de
implantação são requisitos centrais.

Em **serviços de nuvem e rede** a linguagem tem sua maior presença: Docker e
Kubernetes são escritos em Go, e empresas como Dropbox migraram infraestrutura
crítica de Python para Go buscando ganhos de desempenho. É a escolha dominante
para microserviços e arquiteturas serverless.

No desenvolvimento de **CLIs**, o diferencial é o binário único e
autossuficiente gerado pelo compilador, sem runtime externo a instalar.
Ferramentas como Hugo e as CLIs do GitHub e da Stripe foram construídas nesse
modelo.

Em **DevOps e SRE**, ferramentas como Prometheus (monitoramento) e Terraform
(infraestrutura como código) são referências construídas em Go, evidenciando a
adoção pela comunidade de operações.

Go também aparece em bancos de dados (CockroachDB), bioinformática e
**backends web**, com suporte nativo a HTTP/2 e integrações com MySQL, MongoDB
e Elasticsearch. Empresas como Medium, Netflix, Uber e Riot Games figuram
entre os adotantes.

## Linguagem - Usuários Característicos
O perfil típico de usuário Go, chamado de *gopher*, é o de engenheiro que lida com sistemas distribuídos, alta carga ou infraestrutura de
nuvem. Os grupos mais representativos são engenheiros de sistemas em larga
escala, equipes de DevOps e SRE, desenvolvedores de microserviços e APIs, e
desenvolvedores de CLIs que precisam de binários portáteis sem dependências.

Entre as empresas usuárias estão Google, Microsoft, Docker, Cloudflare,
PayPal, American Express, Netflix, Uber, Riot Games e Wildlife Studios.

## Linguagem - Premissas e Diretivas

**Concorrência como cidadã de primeira classe.** O mantra central é: *"não
comunique compartilhando memória; compartilhe memória comunicando"*.
Goroutines são extremamente leves (cerca de 2 KB de stack inicial, contra
megabytes de uma thread de SO), permitindo que milhares rodem simultaneamente.
Channels são o mecanismo preferido de coordenação — intrinsecamente
thread-safe e mais fáceis de compor do que locks explícitos.

**Simplicidade acima de expressividade.** Go deliberadamente omite herança,
sobrecarga de operadores e generics complexos (generics foram adicionados
apenas na v1.18) para manter o código legível por qualquer membro da equipe.

**Binário único e implantação trivial.** O compilador empacota todas as
dependências em um executável estático, eliminando problemas de versionamento
de runtime em produção.

**Garbage collection de baixa latência.** O GC opera com pausas entre 10 e
100 microssegundos, tornando Go viável para sistemas de tempo quase-real.

**Tipagem estática com inferência.** O operador `:=` permite inferência de
tipos sem abrir mão das garantias da tipagem estática.

## Construtores

> Em desenvolvimento.

## Legibilidade

> Em desenvolvimento.

## Capacidade de Escrita

> Em desenvolvimento.

## Confiabilidade

> Em desenvolvimento.

## Custo e Outros Critérios de Avaliação da Linguagem

> Em desenvolvimento.

## Projeto

O projeto é um sistema backend de gerenciamento de aluguel e agendamento de
equipamentos, acessado por múltiplos usuários simultaneamente via requisições
HTTP.

**Domínio:** Backend web, onde a linguagem já demonstrou maturidade em produção em empresas como Medium e Uber.

**Usuário característico:** Equipes ou plataformas que precisam de um serviço
de reservas confiável sob carga simultânea, onde dois usuários não podem
alugar a mesma unidade no mesmo horário.

**Alinhamento com as premissas:** o problema central, múltiplos usuários
acessando o mesmo recurso ao mesmo tempo, é exatamente o cenário para o qual
Go foi projetado. O sistema será implementado em duas versões: uma sem
controle de concorrência, que evidenciará race conditions (o estoque pode
ficar negativo), e outra utilizando goroutines, mutex e channels, demonstrando
o modelo CSP na prática. A comparação de performance entre as duas versões sob
carga será feita com ferramentas de load testing.

## Referências

- Go official website: https://go.dev/
- CTII418- Linguagem Go (IFSP): https://cbt.ifsp.edu.br/images/Documentos/2021/CTII/CTII418_Go.pdf
- COX-BUDAY, Katherine. *Concurrency in Go*. O'Reilly Media, 2017.
- DOXSEY, Caleb. *Introducing Go*. O'Reilly Media, 2016.

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

### Slide - Legibilidade: Aqui tem que ter a comparação com C ?
Legibilidade é a facilidade com que um programa pode ser lido e entendido.

- Simplicidade Global:
    - Go possui uma linguagem pequena e concisa, com poucas palavras-chave e uma sintaxe minimalista.
    - Evita a multiplicidade de recursos, oferecendo uma forma idiomática e clara de realizar a maioria das operações.
    - Exemplo: Apenas um tipo de loop (for), sem while ou do-while.

- Ortogonalidade: Capacidade de combinar um conjunto de estruturas primitivas que podem ser combinadas de forma a criar uma estrutura mais complexa.
    - Go é altamente ortogonal. Um pequeno conjunto de conceitos (como structs, interfaces, goroutines, channels) pode ser combinado de forma consistente.
    - As regras são claras e com poucas exceções, o que reduz a surpresa e facilita a previsão do comportamento do código.
    - A ausência de herança de classes complexas e sobrecarga de operadores contribui para essa ortogonalidade.

- Instruções de Controle:
    - Go promove um fluxo de controle claro e sequencial.
    - Não possui goto, incentivando estruturas de controle mais legíveis (if, for, switch).
    - A formatação de código é padronizada pela ferramenta gofmt, garantindo consistência visual em todos os projetos Go.

- Tipos e Estruturas de Dados:
    - Go é uma linguagem fortemente tipada, o que aumenta a clareza sobre o tipo de dado que está sendo manipulado.
    - Oferece tipos primitivos claros (inteiros, floats, booleanos, strings) e estruturas de dados como structs, arrays, slices e maps que são intuitivas.
    - O uso de interfaces permite abstrações claras sem a complexidade da hierarquia de classes.

- Sintaxe:
    - A sintaxe de Go é projetada para ser concisa e expressiva.
    - Uso de palavras-chave claras (func, type, var, const, package, import).
    - Declarações de variáveis e funções são diretas.
    -  formatação automática pelo gofmt garante um estilo uniforme, eliminando discussões sobre estilo e focando na lógica.

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

### Slide - Capacidade de Escrita: Aqui tem que ter a comparação com C ?
A capacidade de escrita mede a facilidade com que uma linguagem permite criar programas para um determinado domínio.

- **Expressividade:** Go busca um equilíbrio entre concisão e clareza. Embora não tenha todos os operadores compactos de C (como ++ e -- que são apenas pós-fixados e não retornam valor), ela oferece:
    - **Declaração Curta (:=):** Simplifica a declaração e inicialização de variáveis.
    - **Múltiplos Valores de Retorno:** Permite que funções retornem resultados e erros de forma idiomática, tornando o tratamento de erros mais expressivo e menos verboso.
    - **Concorrência Nativa (Goroutines e Canais):** A sintaxe go para iniciar uma goroutine e o uso de chan para comunicação tornam a escrita de código concorrente surpreendentemente simples e expressiva, sem a complexidade de threads e locks manuais.

- Abstração:
    - **Abstração de Processo (Subprogramas):** Go utiliza funções (func) como principal mecanismo de abstração de processo. Elas podem ser passadas como argumentos, retornadas de outras funções e são a base para a modularização do código.
    - **Abstração de Dados (Tipos e Interfaces):**
        - structs permitem definir tipos de dados complexos, encapsulando dados relacionados.
        - Interfaces são um mecanismo poderoso de abstração em Go. Elas permitem definir comportamentos (métodos) sem especificar a implementação, promovendo o polimorfismo e o desacoplamento. Isso facilita a escrita de código flexível e extensível.

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

### Slide -  Confiabilidade: Aqui tem que ter a comparação com C ?
Definição: Um programa é confiável se funciona conforme o esperado em todas as condições. A confiabilidade é crucial para sistemas que precisam operar sem falhas, especialmente em ambientes de produção.

- Verificação de Tipos (Forte e Estática):
    - Go é uma linguagem fortemente tipada e estaticamente tipada. Isso significa que a maioria dos erros de tipo é detectada em tempo de compilação, antes mesmo do programa ser executado.
    - Essa verificação rigorosa evita uma classe inteira de bugs que poderiam surgir em tempo de execução, tornando o desenvolvimento mais seguro e a depuração mais barata.
    - A inferência de tipo (:=) simplifica a escrita sem comprometer a segurança dos tipos.

- Manipulação de Erros (Não Exceções):
    - Go não possui um mecanismo de exceções tradicional (como try-catch). Em vez disso, adota um modelo de tratamento de erros explícito usando múltiplos valores de retorno (o segundo valor geralmente é um error).
    - Isso força o desenvolvedor a considerar e lidar com possíveis falhas em cada ponto onde elas podem ocorrer, tornando o código mais robusto e menos propenso a erros não tratados.
    - O mecanismo panic/recover existe, mas é reservado para erros verdadeiramente excepcionais e irrecuperáveis, não para fluxo de controle normal.

- Aliasing (Ponteiros Seguros):
    - Go possui ponteiros, mas não permite aritmética de ponteiros. Isso elimina uma fonte comum de erros e vulnerabilidades presentes em linguagens como C, onde a manipulação incorreta de ponteiros pode levar a acessos de memória inválidos e aliasing perigoso.
    - A ausência de aliasing arbitrário e a segurança de memória (garantida pelo coletor de lixo e ausência de aritmética de ponteiros) contribuem significativamente para a confiabilidade.

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

### Slide - Custo: Aqui tem que ter a comparação com C ?
- Definição: O custo total de uma linguagem de programação abrange desde o treinamento da equipe até a manutenção do software em produção, incluindo o impacto de falhas.

- Go foi projetada para ser uma linguagem de baixo custo total, otimizando diversos fatores:
    - Treinamento:
        - Go possui uma sintaxe simples e um conjunto reduzido de palavras-chave, o que facilita e acelera o aprendizado para desenvolvedores com experiência em outras linguagens (especialmente C-like).
        - A ortogonalidade e a ausência de recursos complexos (como herança de classes, genéricos complexos, sobrecarga de operadores) reduzem a curva de aprendizado e o tempo para se tornar produtivo.

    - Escrita e Compilação:
        - Escrita: A capacidade de escrita de Go, com sua sintaxe concisa, múltiplos retornos e concorrência nativa, permite que os desenvolvedores escrevam código funcional e eficiente mais rapidamente.
        - Compilação: O compilador de Go é extremamente rápido. Isso acelera o ciclo de desenvolvimento (escrever, compilar, testar), permitindo iterações mais rápidas e feedback imediato.
        - Execução: Embora a compilação seja rápida, os programas Go compilam para código de máquina nativo, resultando em execução de alta performance, comparável a C/C++.

    - Má Confiabilidade:
        - A alta confiabilidade de Go, garantida por sua verificação de tipos estática, gerenciamento automático de memória (GC), tratamento explícito de erros e segurança de ponteiros, minimiza a ocorrência de bugs em produção.
        - Isso reduz drasticamente os custos associados a falhas de sistema, como tempo de inatividade, perda de dados, retrabalho e potenciais danos financeiros ou de reputação.

    - Manutenção:
        - A legibilidade intrínseca do código Go (sintaxe simples, gofmt para padronização, ausência de recursos complexos) torna os programas mais fáceis de entender, depurar e modificar.
        - A capacidade de escrita clara e idiomática resulta em código mais consistente e menos propenso a "truques" difíceis de manter.
        - A confiabilidade inerente significa menos tempo gasto corrigindo bugs inesperados, liberando recursos para o desenvolvimento de novas funcionalidades.
        - Todos esses fatores combinados fazem da manutenção em Go um processo significativamente mais eficiente e menos custoso.


### Slide - Outros Critérios de Avaliação da Linguagem: Portabilidade
A portabilidade refere-se à facilidade com que um programa pode ser movido e executado em diferentes ambientes de hardware e sistemas operacionais, sem a necessidade de grandes modificações.

- Compilação para Código Nativo:
    - Go compila diretamente para código de máquina nativo, sem a necessidade de uma máquina virtual (JVM, .NET CLR) ou interpretador em tempo de execução.
    - Isso permite gerar executáveis autônomos e otimizados para diversas arquiteturas e sistemas operacionais.

- Compilação Cruzada (Cross-Compilation):
    - Go possui suporte nativo e robusto para compilação cruzada. Um desenvolvedor pode compilar um programa Go em um sistema operacional (ex: Linux) para ser executado em outro (ex: Windows ou macOS) ou para diferentes arquiteturas de CPU (ex: ARM, x86) com um simples comando.
    - Isso elimina a complexidade de configurar ambientes de desenvolvimento específicos para cada plataforma de destino.

- Ausência de Dependências Externas:
    - Os executáveis Go são estaticamente linkados por padrão, o que significa que todas as bibliotecas necessárias são empacotadas no próprio binário.
    - Isso reduz significativamente as dependências externas e facilita a implantação, pois o executável pode ser copiado e executado diretamente na maioria dos sistemas, sem a necessidade de instalar runtimes ou bibliotecas adicionais.

- Padronização da Linguagem:
    - A especificação da linguagem Go é bem definida e implementada de forma consistente em todas as plataformas suportadas, garantindo que o código se comporte da mesma maneira onde quer que seja executado.

### Slide - Outros Critérios de Avaliação da Linguagem: Generalidade
A generalidade (ou aplicabilidade) de uma linguagem refere-se à sua capacidade de ser utilizada em uma ampla variedade de domínios de aplicação, sem ser excessivamente especializada.

- Go é uma linguagem de propósito geral, projetada para ser eficaz em diversos cenários:
    - Desenvolvimento de Sistemas:
        - Go é amplamente utilizada para construir sistemas de backend de alta performance, como APIs RESTful, microsserviços, servidores web e ferramentas de linha de comando.
        - Sua concorrência nativa e eficiência a tornam ideal para lidar com grandes volumes de requisições e processamento paralelo.

    - Ferramentas de Rede:
        - Com bibliotecas padrão robustas para rede (net/http, net), Go é uma escolha popular para o desenvolvimento de ferramentas de rede, proxies, balanceadores de carga e outros componentes de infraestrutura.

    - DevOps e Ferramentas de Infraestrutura:
        - A capacidade de compilar para binários estáticos e a facilidade de compilação cruzada tornam Go excelente para criar ferramentas de DevOps, automação, CLI (Command Line Interface) e agentes de monitoramento. 
        - Exemplos incluem Docker, Kubernetes e Terraform, que são escritos em Go.

    - Computação em Nuvem:
        - Go é a linguagem preferida para muitos projetos e serviços em ambientes de computação em nuvem, devido à sua eficiência, escalabilidade e facilidade de implantação.

    - Outros Domínios:
        - Embora menos comum, Go também pode ser usada em processamento de dados, machine learning (com bibliotecas específicas), e até mesmo em alguns projetos embarcados onde a performance e o baixo consumo de recursos são cruciais.

### Slides - Explicação do Projeto


### Slide - Referências

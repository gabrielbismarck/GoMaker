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
- [ ] Site / Demonstração
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

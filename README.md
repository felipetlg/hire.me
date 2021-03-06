# Hire.me
Um pequeno projeto para testar suas habilidades como programador.

## Instruções Gerais

1. *Clone* este repositório
2. Em seu *fork*, atenda os casos de usos especificados e se desejar também os bonus points
3. Envio um e-mail para rh@bemobi.com.br com a seu Nome e endereço do repositorio.

## Projeto

O projeto consiste em reproduzir um encurtador de URL's (apenas sua API), simples e com poucas funções, porém com espaço suficiente para mostrar toda a gama de desenho de soluções, escolha de componentes, mapeamento ORM, uso de bibliotecas de terceiros, uso de GIT e criatividade.

O projeto consiste de dois casos de uso: 

1. Shorten URL
2. Retrieve URL

### 1 - Shorten URL
![Short URL](http://i.imgur.com/MFB7VP4.jpg)

1. Usuario chama a API passando a URL que deseja encurtar e um parametro opcional **CUSTOM_ALIAS**
    1. Caso o **CUSTOM_ALIAS** já exista, um erro especifico ```{ERR_CODE: 001, Description:CUSTOM ALIAS ALREADY EXISTS}``` deve ser retornado.
    2. Toda URL criada sem um **CUSTOM_ALIAS** deve ser reduzida a um novo alias, **você deve sugerir um algoritmo para isto e o porquê.**
    
2. O Registro é colocado em um repositório (*Data Store*)
3. É retornado para o cliente um resultado que contenha a URL encurtada e outros detalhes como
    1. Quanto tempo a operação levou
    2. URL Original

Exemplos (Você não precisa seguir este formato):

* Chamada sem CUSTOM_ALIAS
```
PUT http://shortener/create?url=http://www.bemobi.com.br

{
   "alias": "XYhakR",
   "url": "http://shortener/u/XYhakR",
   "statistics": {
       "time_taken": "10ms",
   }
}
```

* Chamada com CUSTOM_ALIAS
```
PUT http://shortener/create?url=http://www.bemobi.com.br&CUSTOM_ALIAS=bemobi

{
   "alias": "bemobi",
   "url": "http://shortener/u/bemobi",
   "statistics": {
       "time_taken": "12ms",
   }
}
```

* Chamada com CUSTOM_ALIAS que já existe
```
PUT http://shortener/create?url=http://www.github.com&CUSTOM_ALIAS=bemobi

{
   "alias": "bemobi",
   "err_code": "001",
   "description": "CUSTOM ALIAS ALREADY EXISTS"
}
```

### 2 - Retrieve URL
![Retrieve URL](http://i.imgur.com/f9HESb7.jpg)

1. Usuario chama a API passando a URL que deseja acessar
    1. Caso a **URL** não exista, um erro especifico ```{ERR_CODE: 002, Description:SHORTENED URL NOT FOUND}``` deve ser retornado.
2. O Registro é lido de um repositório (*Data Store*)
3. Esta tupla ou registro é mapeado para uma entidade de seu projeto
3. É retornado para o cliente um resultado que contenha a URL final, a qual ele deve ser redirecionado automaticamente

## Stack Tecnológico

Não há requerimentos específicos para linguagens, somos poliglotas. Utilize a linguagem que você se sente mais confortável.

## Bonus Points

1. Crie *testcases* para todas as funcionalidades criadas
2. Crie um *endpoint* que mostre as dez *URL's* mais acessadas 
3. Crie um *client* para chamar sua API
4. Faça um diagrama de sequencia da implementação feita nos casos de uso (Dica, use o https://www.websequencediagrams.com/)
5. Monte um deploy da sua solução utilizando containers 

## Realização do projeto

1. A API foi desenvolvida em Go
2. Foi criada uma interface simples em html+javascript que usa a API
3. O banco de dados foi criado usando sqlite3 com uma tabela criada da seguinte forma:
```CREATE TABLE urls(alias TEXT PRIMARY KEY, longUrl TEXT NOT NULL, shortUrl TEXT NOT NULL, visits INTEGER DEFAULT 0)```

## Build & Run

* Com o Docker:
1. É necessária a instalação do Docker
    1. Verificar instruções em https://docs.docker.com/get-docker/
3. Navegar até a pasta do projeto clonado
2. **Build:** ```$ docker build -t api_shortener .```
4. **Run:** ```$ docker run -d -p 8080:8080 api_shortener```

* Sem o Docker:
1. Com a instalação do Go (https://golang.org/dl/)
2. Navegar até a pasta do projeto clonado
3. Executar ```$ go get``` (busca as dependências definidas em "go.mod")
4. **Run:** na pasta do projeto clonado: ```$ go run main.go```

## Usando a API

Após seguir os passos da sessão anterior:

* Com a interface:

Foi criada uma interface onde é possível usar o navegador web e ir até: http://localhost:8080/, onde é possível:
1. Criar uma nova URL curta ao inserir a URL original e um alias (opcional)
2. Após a criação é possível usar a URL curta da seguinte forma: http://localhost:8080/s/{ALIAS} (onde {ALIAS} deve ser substituído pelo criado)
3. Visualizar e acessar as top 10 URL's curtas que ficam disponíveis na interface

* Sem a interface:

Aqui vão alguns exemplos usando curl
-- Com alias
```
curl -L -X POST 'localhost:8080' -H 'Content-Type: application/json' --data-raw '{
"alias": "bemobi",
"longUrl": "https://www.bemobi.com.br/careers.html"
}'
```

-- Sem alias
```
curl -L -X POST 'localhost:8080' -H 'Content-Type: application/json' --data-raw '{
"longUrl": "https://www.google.com/maps/place/Bemobi/@-22.9461593,-43.1852983,17z/data=!3m1!4b1!4m5!3m4!1s0x997ff0ef611383:0x4b66002fd48e7656!8m2!3d-22.9461643!4d-43.1831043"
}'
```

## Testes unitários

A camada de serviços possui testes unitários e para rodá-los:
1. Navegar até "hire.me/service"
2. Executar ```$ go test -v```

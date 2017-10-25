# Widgets API

API de serviço de widgets e usuarios para ser consumido pelo projeto [Widgets SPA](https://github.com/RedVentures/widgets-spa).

## Tecnologias utilizadas

A API foi produzida usando a linguagem [Go](https://golang.org/), com persistência de dados em [MySQL](https://www.mysql.com/). Foram também usadas as seguintes bibliotecas:

  - [gorilla/mux](github.com/gorilla/mux) - Gerador de rotas
  - [dgrijalva/jwt-go](github.com/dgrijalva/jwt-go) - Geração e autênticação de *json web tokens *
  - [jinzhu/gorm](github.com/jinzhu/gorm) - ORM e gerenciador de migrações
  - [joho/godotenv](github.com/joho/godotenv) - Carregamento de arquivo .env para variáveis de sistema
  - [validator.v9](gopkg.in/go-playground/validator.v9) - Validação de modelos

## Como executar

Para executar a API é necessário ter um aquivo .env com as credenciais do banco de dados, ou carregar as credenciais em variáveis de sistema. O arquivo [.env.example](./.env.example) contém exemplo de credenciais.

### Variáveis para banco de dados
- DB_HOST - Host do banco a ser acessado
- DB_PORT - Porta usada para acessar o banco de dados
- DB_DATABASE - Banco usado
- DB_USERNAME - Usuário do banco de dados 
- DB_PASSWORD - Senha do banco de dados
- DB_AUTO_MIGRATE - Boolean definindo se deve ser realizada a migração das tabelas 
- DB_AUTO_POPULATE - Boolean definindo se deve ser realizada uma população inicial ao crias as tabelas
 
### Executando Nativamente

Para executar a API, basta executar os comandos:

```sh
$ go get github.com/guilhermeartem/widgets-api
$ widgets-api
```

A API pode ser acessada em [localhost:4000](localhost:4000), e a documentação das rotas em [localhost:4000/docs](localhost:4000/docs).

### Executando com docker-compose

É provido um Dockerfile para geração de uma imagem para Docker e um arquivo docker-compose para facilitar o uso. Para executar desse modo, deve ser feito:

```sh
$ go get github.com/guilhermeartem/widgets-api
$ cd $GOPATH/src/github.com/guilhermeartem/widgets-api
$ docker-compose up
```

A API pode ser acessada em [localhost:4000](localhost:4000), e a documentação das rotas em [localhost:4000/docs](localhost:4000/docs).

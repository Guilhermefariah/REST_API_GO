# Rest API GO

Este é um projeto de uma API RESTful escrita em Go, utilizando o framework Gin. A API é responsável por gerenciar produtos, permitindo operações de CRUD (criação, leitura, atualização e exclusão).

## Endpoints

### GET /ping

Retorna uma mensagem de "pong" para verificar se a API está funcionando corretamente.

### GET /products

Retorna uma lista de produtos.

### POST /product

Cria um novo produto.

### GET /product/:id

Retorna um produto específico pelo ID.

## Modelo de Dados

O modelo de dados utilizado é o seguinte:

```go
type Product struct {
    ID    int     `json:"id_product"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}
```

## Dependências

* Gin: framework web para Go
* PostgreSQL: banco de dados relacional
* github.com/lib/pq: driver de banco de dados para PostgreSQL

## Configuração

Para executar o projeto, é necessário ter o Go e o PostgreSQL instalados. Além disso, é necessário criar um banco de dados e uma tabela de produtos com as seguintes colunas:

```sql
CREATE TABLE product (
    id SERIAL PRIMARY KEY,
    product_name VARCHAR(255),
    price DECIMAL(10, 2)
);
```

## Execução

Para executar o projeto, basta executar o comando `go run main.go`. A API estará disponível em `http://localhost:8000`.

## Testes

Os testes estão disponíveis no diretório `tests`. Para executá-los, basta executar o comando `go test`.

## Contribuição

Contribuições são bem-vindas! Se você tiver alguma sugestão ou correção, por favor, abra uma issue ou envie um pull request.
# API Rest Genesis

API Rest que realiza conversão de moedas utilizando Go, Fiber, GORM, Docker e MySQL.

## Requisitos

- Docker
- Docker Compose

## Como executar

1. Clone o repositório:

```sh
git clone https://github.com/SEU_USUARIO/api-rest-genesis.git
```

2. Entre na pasta: 

```sh
cd api-rest-genesis
```

3. Construa a imagem Docker e inicie: 

```sh
docker-compose up --build
```

4. Acesse a API em http://localhost:8080

<br>
<br>

# Endpoints GET /convert/:amount/:from/:to/:rate

Converte uma quantidade de uma moeda para outra, com base em uma taxa de conversão.

Parâmetros:
- amount: valor a ser convertido (número decimal)
- from: código da moeda de origem (string)
- to: código da moeda de destino (string)
- rate: taxa de conversão (número decimal)

<br>
<br>

# Outros Endpoints:

GET / - Exibe mensagem de boas-vindas e instruções de como utilizar a API.
GET /logs - Exibe o histórico de conversões realizadas.

Tecnologias utilizadas
- Go
- Fiber
- GORM
- MySQL
- Docker
- Docker Compose

### Licença MIT

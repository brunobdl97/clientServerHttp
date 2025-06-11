# 💵 Projeto Cotação Dólar (Go)

Este projeto é composto por:

- 🖥️ Um **servidor HTTP** que fornece a cotação atual do dólar.
- 📥 Um **cliente** que consome esse endpoint e salva o valor em um arquivo de texto.

## 📋 Pré-requisitos

Antes de começar, certifique-se de ter os seguintes itens instalados:

- ✅ [Go](https://golang.org/dl/) (versão **1.18** ou superior)
- ✅ [SQLite3](https://www.sqlite.org/download.html)
- ✅ [Git](https://git-scm.com/) (opcional, para clonar o repositório)

## ⚙️ Instalação

Clone o repositório e instale as dependências:

```bash
git clone https://github.com/brunobdl97/clientServerHttp.git
cd clientServerHttp
go mod tidy
```

## 🚀 Inicializando o Servidor
No diretório server, execute:
```bash
cd server
go run main.go
```
O servidor estará disponível em:
🔗 http://localhost:8080/cotacao

## 📡 Executando o Cliente
Em outro terminal, acesse o diretório client e execute:
```bash
cd ../client
go run main.go
```
O cliente irá requisitar a cotação do servidor e salvar o valor no arquivo cotacao.txt.

## 🗒️ Observações
🗃️ O banco de dados SQLite será criado automaticamente como quote.db no diretório do servidor.

📄 O arquivo cotacao.txt será gerado no diretório do cliente com a última cotação do dólar.

## 📁 Estrutura do Projeto
```bash
clientServerHttp/
├── client/
│   └── main.go
├── server/
│   ├── main.go
│   └── quote.db (gerado automaticamente)
└── go.mod
```

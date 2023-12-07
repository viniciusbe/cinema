# Guia de compilação e execução

Este guia passo a passo irá ajudá-lo a compilar e executar a aplicação. Siga as etapas abaixo:

## 1. Pré-requisitos

Certifique-se de ter os seguintes pré-requisitos instalados em seu sistema:

- [Go 1.21.1](https://golang.org/doc/install)
- [Neo4j 5.14.0](https://neo4j.com/download/)

## 2. Obtendo o Código da Aplicação

Clone o repositório da sua aplicação Go do GitHub ou qualquer outra fonte.

   ```bash
   git clone https://github.com/viniciusbe/cinema.git
   cd cinema
   git checkout neo4j-db
   ```

## 3. Configurando o banco de dados

O arquivo `backup.txt` mostra os scripts usados para fazer a migração do PostgreSQL.

## 4. Configuração da Aplicação

Abra o arquivo de configuração da sua aplicação (`internal/handlers/neo4j/connection.go`) e atualize as configurações do banco de dados com as informações do banco de dados Neo4j que você configurou anteriormente (nome do banco de dados, nome de usuário, senha, host, porta, etc.).

## 5. Compilação e Execução da Aplicação

1. Acesse a pasta que contém o arquivo principal da aplicação:

   ```bash
   cd ./cmd/cinemacli
   ```

2. Compile a aplicação:

   ```bash
   go build
   ```

3. Execute a aplicação:

   ```bash
   ./cinemacli
   ```

Agora, sua aplicação deve estar em execução e conectada ao banco de dados.

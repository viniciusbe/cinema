# Guia de Compilação e Execução da Aplicação em Go com PostgreSQL 15

Este guia passo a passo irá ajudá-lo a compilar e executar a aplicação em Go que utiliza o banco de dados PostgreSQL 15. Siga as etapas abaixo:

## 1. Pré-requisitos

Certifique-se de ter os seguintes pré-requisitos instalados em seu sistema:

- [Go 1.21.1](https://golang.org/doc/install)
- [PostgreSQL 15](https://www.postgresql.org/download/)

## 1. Obtendo o Código da Aplicação

1. Clone o repositório da sua aplicação Go do GitHub ou qualquer outra fonte.

   ```bash
   git clone https://github.com/viniciusbe/cinema.git
   cd cinema
   ```

## 2. Configurando o banco de dados

1. Caso queira iniciar o projeto com alguns registros no banco, use o arquivo de backup `backup.sql` localizado na raiz do repositório. Se não, pule para a próxima etapa. Obs: não é necessário criar as tabelas no banco, a aplicação já é resposável por criá-las na primeira execução.

## 3. Configuração da Aplicação

1. Abra o arquivo de configuração da sua aplicação (`internal\handlers\gormdb\connection.go`) e atualize as configurações do banco de dados com as informações do banco de dados PostgreSQL que você configurou anteriormente (nome do banco de dados, nome de usuário, senha, host, porta, etc.).

## 4. Compilação e Execução da Aplicação

1. Acesse a pasta que contém o arquivo principal da aplicação:

   ```bash
   cd .\cmd\cinemacli
   ```

2. Compile a aplicação:

   ```bash
   go build
   ```

3. Execute a aplicação:

   ```bash
   ./cinemacli
   ```

Agora, sua aplicação em Go deve estar em execução e conectada ao banco de dados.

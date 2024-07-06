# Client and Server Challenge

Sistemas em *Golang* para obter a cotação do dólar em real - USDBRL - e armazená-lo em um arquivo local e em um banco de dados.

**Utilizando para isso:** Requisições Http, Context, Banco de dados e Manipulação de arquivos.

## Requisitos básicos

- Go
- Sqlite3

## Procedimentos iniciais - Banco de dados

Antes de executar os serviços é necessário executar as seguintes queries no banco **./db/db-cotacao.db**:

```sql
CREATE TABLE cotacao (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    bid text
);

INSERT INTO cotacao (bid) VALUES (0);
inserir na Tabela cotacao
```

**Tabela cotacao inicialmente:**

| ID    | BID  |
| --- | --- |
| 1   | 0 |


## Procedimentos gerais para utilização dos sistemas

- Inicializar o go utilizando os comandos **go mod init** e **go mod tidy**;
- Subir o server utilizando o comando: **go run server/main.go**;
- Executar o client executando: **go run client/main.go**.

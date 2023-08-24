package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Cotacao struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func main() {
	http.HandleFunc("/cotacao", FindCotacaoHandler)
	http.ListenAndServe(":8080", nil)
}

func FindCotacaoHandler(w http.ResponseWriter, r *http.Request) {
	ctxHttp, cancelHttp := context.WithTimeout(r.Context(), 200*time.Millisecond)
	defer cancelHttp()

	cotacao, err := GetCotacao(ctxHttp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	db, _ := sql.Open("sqlite3", "../db/db-cotacao.db")
	defer db.Close()

	ctxDB, cancelDB := context.WithTimeout(r.Context(), 10*time.Millisecond)
	defer cancelDB()
	persistirCotacao(ctxDB, db, cotacao)

	json.NewEncoder(w).Encode(cotacao.USDBRL.Bid)
}

func GetCotacao(ctx context.Context) (*Cotacao, error) {
	resp, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	cotacaoResp, err := http.DefaultClient.Do(resp)
	if err != nil {
		log.Println("API cotacao -", err)
		return nil, err
	}
	defer cotacaoResp.Body.Close()

	body, err := io.ReadAll(cotacaoResp.Body)
	if err != nil {
		return nil, err
	}

	var cotacao Cotacao
	err = json.Unmarshal(body, &cotacao)
	if err != nil {
		return nil, err
	}

	return &cotacao, nil
}

func persistirCotacao(ctxDB context.Context, db *sql.DB, cotacao *Cotacao) {
	updateCotacao := `UPDATE cotacao SET (bid) = (?)`
	stmt, err := db.Prepare(updateCotacao)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = stmt.ExecContext(ctxDB, cotacao.USDBRL.Bid)
	if err != nil {
		log.Println("Update -", err)
	}
}

// Comandos usados para criar / inicializar o banco de dados:

// CREATE TABLE cotacao (
//    id INTEGER PRIMARY KEY AUTOINCREMENT,
//    bid text
// );

// INSERT INTO cotacao (bid) VALUES (0);

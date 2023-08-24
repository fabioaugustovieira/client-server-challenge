package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctxHttp, cancel := context.WithTimeout(context.Background(), time.Duration(300*time.Millisecond))
	defer cancel()

	req, err := http.NewRequestWithContext(ctxHttp, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Println("Server -", err)
	}

	serverResp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer serverResp.Body.Close()

	arquivoCotacao, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}
	defer arquivoCotacao.Close()

	body, err := io.ReadAll(serverResp.Body)
	if err != nil {
		panic(err)
	}

	var bid string
	err = json.Unmarshal(body, &bid)
	if err != nil {
		panic(err)
	}

	cotacaoString := fmt.Sprintf("Dólar: %s", bid)
	_, err = arquivoCotacao.WriteString(cotacaoString)
	if err != nil {
		panic(err)
	}
}

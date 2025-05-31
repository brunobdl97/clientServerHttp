package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Quote struct {
	Dolar string `json:"dolar"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	startRequest := time.Now()
	request, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}

	response, err := http.DefaultClient.Do(request)
	elapsed := time.Since(startRequest)
	fmt.Printf("Tempo de resposta do endpoint: %d ms\n", elapsed.Milliseconds())

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var quote Quote
	if err := json.NewDecoder(response.Body).Decode(&quote); err != nil {
		panic(err)
	}

	f, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("Dólar: " + quote.Dolar)
	io.Copy(os.Stdout, response.Body)
}

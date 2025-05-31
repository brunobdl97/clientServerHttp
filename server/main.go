package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/cotacao", handler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Erro ao iniciar o servidor:")
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	db, err := OpenDB()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	defer db.Close()

	ctx := r.Context()

	response, err := CallQuote(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	if err := SaveQuote(ctx, db, response); err != nil {
		log.Fatal(err.Error())
		return
	}

	buildResponse(w, response)
}

func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./quote.db")
	if err != nil {
		return nil, err
	}

	createTableQuotes := `
	CREATE TABLE IF NOT EXISTS quotes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code TEXT,
		codein TEXT,
		name TEXT,
		high TEXT,
		low TEXT,
		varBid TEXT,
		pctChange TEXT,
		bid TEXT,
		ask TEXT,
		timestamp TEXT,
		create_date TEXT
	);`

	_, err = db.Exec(createTableQuotes)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func buildResponse(w http.ResponseWriter, response *DollarQuote) {
	quote := map[string]interface{}{
		"dolar": response.Quote.Bid,
	}

	parsedResponse, err := json.Marshal(quote)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(parsedResponse)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

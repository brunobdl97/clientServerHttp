package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

const (
	DollarCote = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
)

type DollarQuote struct {
	Quote struct {
		ID         string `json:"id"`
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

func CallQuote(ctx context.Context) (*DollarQuote, error) {
	request, err := http.NewRequest(http.MethodGet, DollarCote, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, err
	}

	var quote *DollarQuote
	if err := json.NewDecoder(response.Body).Decode(&quote); err != nil {
		return nil, err
	}

	select {
	case <-time.After(200 * time.Millisecond):
		if quote == nil {
			return nil, errors.New("request timeout to get dollar quote")
		}

		return quote, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return quote, nil
	}
}

func SaveQuote(ctx context.Context, db *sql.DB, quote *DollarQuote) error {
	query := `INSERT INTO quotes (code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	stmt, err := db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(
		quote.Quote.Code,
		quote.Quote.Codein,
		quote.Quote.Name,
		quote.Quote.High,
		quote.Quote.Low,
		quote.Quote.VarBid,
		quote.Quote.PctChange,
		quote.Quote.Bid,
		quote.Quote.Ask,
		quote.Quote.Timestamp,
		quote.Quote.CreateDate,
	)
	if err != nil {
		return err
	}

	select {
	case <-time.After(10 * time.Millisecond):
		affected, err := result.RowsAffected()
		if err != nil {
			return errors.New("sssssssssssssssss timeout to get dollar quote")
		}

		if affected == 0 {
			return errors.New("request timeout to get dollar quote")
		}

		return nil
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	_ "modernc.org/sqlite"
)

const REQUEST_TIMEOUT = time.Millisecond * 200
const DATABASE_TIMEOUT = time.Millisecond * 10
const DATABASE_URL = "../data/goexpert_sqlite.db"
const SERVER_PORT = ":8080"

type Endpoint string

const (
	cotacaoDolarEndpoint = Endpoint("https://economia.awesomeapi.com.br/json/last/USD-BRL")
)

type Cotacao struct {
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
}

type CotacaoResponse struct {
	USDBRL Cotacao `json:"USDBRL"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", handleCotacao)
	http.ListenAndServe(SERVER_PORT, mux)
}

func handleCotacao(w http.ResponseWriter, r *http.Request) {
	requestCtx, cancel := context.WithTimeout(context.Background(), REQUEST_TIMEOUT)
	defer cancel()

	cotacaoResponse, err := getCotacaoDolar(requestCtx)
	if err != nil {
		if requestCtx.Err() == context.DeadlineExceeded {
			http.Error(w, "Request to economia.awesomeapi timed out", http.StatusRequestTimeout)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dbCtx, cancel := context.WithTimeout(context.Background(), DATABASE_TIMEOUT)
	defer cancel()

	db, err := sql.Open("sqlite", DATABASE_URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	err = insertCotacaoOnDB(dbCtx, db, cotacaoResponse)
	if err != nil {
		if requestCtx.Err() == context.DeadlineExceeded {
			http.Error(w, "Database timed out", http.StatusGatewayTimeout)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cotacaoResponse)
}

func getCotacaoDolar(ctx context.Context) (*CotacaoResponse, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", string(cotacaoDolarEndpoint), nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var cotacaoResponse CotacaoResponse
	err = json.NewDecoder(res.Body).Decode(&cotacaoResponse)
	if err != nil {
		return nil, err
	}

	return &cotacaoResponse, nil
}

func insertCotacaoOnDB(ctx context.Context, db *sql.DB, cotacao *CotacaoResponse) error {
	stmt, err := db.Prepare("insert into currency_exchange_BRL (coin, value) values (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, cotacao.USDBRL.Code, cotacao.USDBRL.Bid)
	if err != nil {
		return err
	}
	return nil
}

package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const SERVER_ENDPOINT = "http://localhost:8080"
const REQUEST_TIMEOUT = time.Millisecond * 300

type Endpoint string

const (
	cotacaoEndpoint = Endpoint("/cotacao")
)

func (e Endpoint) getApiURL() string {
	return SERVER_ENDPOINT + string(e)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), REQUEST_TIMEOUT)
	defer cancel()

	err := saveCotacao(ctx)
	if err != nil {
		panic(err)
	}
}

func saveCotacao(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, "GET", cotacaoEndpoint.getApiURL(), nil)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status code", res.StatusCode)
		fmt.Printf("Error: ")
		io.Copy(os.Stderr, res.Body)
		return nil
	}

	f, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = io.Copy(f, res.Body)
	if err != nil {
		return err
	}
	return nil
}

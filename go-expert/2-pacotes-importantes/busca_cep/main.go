package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradoruro string `json:"logradouro"`
	Complemento string `json:"complemento"`
}

func main() {
	http.HandleFunc("/", BuscaCEPHandler)
	err := http.ListenAndServe(":8080", nil)
	panic(err)
}

func BuscaCEPHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cep, error := BuscaCEP(cepParam)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cep)
}

func BuscaCEP(cep string) (*ViaCep, error) {
	resp, error := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()
	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}
	var c ViaCep
	error = json.Unmarshal(body, &c)
	return &c, nil
}

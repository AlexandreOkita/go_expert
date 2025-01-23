package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s" validate:"gt=0"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(string(res))

	encoder := json.NewEncoder((os.Stdout))
	err = encoder.Encode(conta)
	if err != nil {
		println(err)
	}

	jsonPuro := []byte(`{"n":2, "s":-200}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(contaX)
}

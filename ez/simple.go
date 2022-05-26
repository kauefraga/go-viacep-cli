package main

import (
	"fmt"

	"github.com/imroc/req/v3"
)

func main() {
	var cep string

	fmt.Print("Enter a cep: ")
	fmt.Scan(&cep)

	res, err := req.Get("https://viacep.com.br/ws/" + cep + "/json/")

	if err != nil {
		println(err)
	}

	fmt.Println(res)
}

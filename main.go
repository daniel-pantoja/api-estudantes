package main

import (
	"log"

	"github.com/daniel-pantoja/api-estudante/api"
)

func main() {
	servidor := api.NovoServidor()

	servidor.ConfigRotas()

	if err := servidor.ComecaServidor(); err != nil {
		log.Fatal(err)
	}
}
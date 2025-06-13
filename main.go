package main

import (
	"github.com/rs/zerolog/log"

	"github.com/daniel-pantoja/api-estudante/api"
)

func main() {
	servidor := api.NovoServidor()

	servidor.ConfigRotas()

	if err := servidor.ComecaServidor(); err != nil {
		log.Fatal().Err(err).Msgf("Falha ao iniciar o servidor")
	}
}
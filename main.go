package main

import (
	"github.com/paula-michele-brisa/backend-campeonato/config"
	"github.com/paula-michele-brisa/backend-campeonato/router"
)

func main() {
	config.OpenDB()
	router.SetupRouter()

}

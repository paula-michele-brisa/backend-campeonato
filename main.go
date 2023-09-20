package main

import "github.com/paula-michele-brisa/backend-campeonato/router"

func main() {
	initialize := router.SetUpRouter()

	initialize.Run()

}

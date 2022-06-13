package main

import (
	"github.com/luizclaudioholanda/api-go-gin/database"
	"github.com/luizclaudioholanda/api-go-gin/routes"
)

func main() {
	database.ConectaBancoDeDados()

	routes.HandleRequests()
}

package main

import (
	"./config"
	"./database"
	"./server"
)

func main() {
	config.Init()
	database.Init()
	server.Init()
}

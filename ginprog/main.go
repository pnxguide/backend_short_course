package main

import (
	"./config"
	"./server"
)

func main() {
	config.Init()
	server.Init()
}

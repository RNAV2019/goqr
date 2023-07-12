package main

import (
	"rnav2022/goqr/model"
	"rnav2022/goqr/server"
)

func main() {
	model.Setup()
	server.SetupAndListen()
}

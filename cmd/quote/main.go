package main

import (
	"quote/internal/handler"
	env "quote/pkg/init"
)

func main() {
	env.Init()
	//go run.Bot()
	handler.Route()
}

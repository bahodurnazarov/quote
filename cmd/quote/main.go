package main

import (
	"log"
	run "quote/internal/bot"
	lg "quote/pkg/utils"
)

func main() {
	go run.Bot()
	lg.Errl.Println("hdf")
	lg.Server.Println("32")
	log.Println("3ed")
}

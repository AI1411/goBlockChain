package main

import (
	"flag"
	"log"
)

func init() {
	//logを生成するための処理
	log.SetPrefix("Blockchain:  ")
}

func main() {
	port := flag.Uint("port", 5000, "TCP Port number blockchain server")
	flag.Parse()
	app := NewBlockchainServer(uint16(*port))
	app.Run()
}

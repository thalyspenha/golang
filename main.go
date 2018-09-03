package main

import (
	"log"

	"github.com/golang/api"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	api := api.App{}
	api.StartServer()

}

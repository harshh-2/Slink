package main

import (
	"log"
	"github.com/harshh-2/slink/internal/app"
)

func main() {
	srv := app.NewApp()
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
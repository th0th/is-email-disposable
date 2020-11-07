package main

import (
	"github.com/th0th/is-email-disposable/domain"
	"github.com/th0th/is-email-disposable/http"
	"log"
)

func main() {
	d, err := domain.NewDomain()

	if err != nil {
		log.Panic(err)
	}

	server := http.NewServer(d)

	server.Run()
}

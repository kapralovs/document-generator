package main

import (
	"log"

	"github.com/kapralovs/document-generator/internal/server"
)

func main() {
	s := server.New()

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}

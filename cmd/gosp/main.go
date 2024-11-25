package main

import (
	"log"

	"github.com/gabefiori/gosp/internal/cli"
)

func main() {
	if err := cli.Run(); err != nil {
		log.Fatal(err)
	}
}

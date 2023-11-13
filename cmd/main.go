package main

import (
	"log"

	"github.com/threadedstream/pixelblender/cmd/simple"
)

func main() {
	if err := simple.Main(); err != nil {
		log.Fatal(err)
	}
}

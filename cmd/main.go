package main

import (
	"log"

	"github.com/qiloz/nexus-shaihulud-scan/cmd/flagParser"
)

func main() {
	_, err := flagParser.Init()
	if err != nil {
		log.Fatal(err)
	}
}

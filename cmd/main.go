package main

import (
	"log"

	"github.com/qiloz/nexus-shaihulud-scan/cmd/flagParser"
	"github.com/qiloz/nexus-shaihulud-scan/internal/nexusPkgPuller"
)

func main() {
	flags, err := flagParser.Init()
	if err != nil {
		log.Fatal(err)
	}

	err = nexusPkgPuller.GetRepositoryPackages(flags)
	if err != nil {
		log.Fatal(err)
	}
}

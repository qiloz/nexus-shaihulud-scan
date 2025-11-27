package main

import (
	"log"
	"os"

	"github.com/qiloz/nexus-shaihulud-scan/cmd/flagParser"
	"github.com/qiloz/nexus-shaihulud-scan/cmd/utilEnv"
	"github.com/qiloz/nexus-shaihulud-scan/internal/nexusPkgPuller"
)

func main() {
	flags, err := flagParser.Init()
	if err != nil {
		log.Fatal(err)
	}

	utilEnv.PrintStartupCfg(flags)

	err = nexusPkgPuller.GetRepositoryPackages(flags)
	if err != nil {
		log.Fatal(err)
	}

	if flags.NxNoScanMode {
		os.Exit(0)
	}
}

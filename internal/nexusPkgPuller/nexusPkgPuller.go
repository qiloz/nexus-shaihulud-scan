package nexusPkgPuller

import (
	"fmt"
	"os"
	"path"

	"github.com/qiloz/nexus-shaihulud-scan/cmd/utilEnv"
	"github.com/qiloz/nexus-shaihulud-scan/internal/nexusApiProxy"
	"github.com/qiloz/nexus-shaihulud-scan/internal/utilityFlags"
)

func GetRepositoryPackages(flags utilityFlags.Flags) error {
	utilDirPath, err := utilEnv.CreateUtilDir()
	if err != nil {
		return err
	}

	var pkgRepoListPath = path.Join(utilDirPath, "nx-"+flags.NxRepoName+"-npm-pkgs.csv")

	f, err := os.Create(pkgRepoListPath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString("Package Name,Version\n")
	if err != nil {
		return err
	}

	var continuationToken string

	fmt.Print("Please wait: ")
	for {
		assetData, err := nexusApiProxy.GetNxRepoAssets(flags, continuationToken)
		if err != nil {
			return err
		}

		for _, asset := range assetData.Items {
			_, err = f.WriteString(asset.Npm.Name + "," + asset.Npm.Version + "\n")
			if err != nil {
				return err
			}
		}

		if assetData.ContinuationToken == "" {
			break
		}
		fmt.Print("|")
		continuationToken = assetData.ContinuationToken
	}
	return nil
}

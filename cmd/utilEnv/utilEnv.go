package utilEnv

import (
	"fmt"
	"os"
	"path"

	"github.com/qiloz/nexus-shaihulud-scan/internal/utilityFlags"
)

const tempDirName = ".nx-shaihd-scan"

func PrintStartupCfg(flags utilityFlags.Flags) {
	fmt.Println("SHA1-Hulud nexus repository scanner\n ----------")
	fmt.Printf("Nexus Instance URI: %s\n", flags.NxInstanceUri)
	fmt.Printf("Target repository name: %s\n", flags.NxRepoName)

	if flags.NxAnonymousMode {
		fmt.Printf("Anonymous mode: %t \n", flags.NxAnonymousMode)
	}
	fmt.Println(" ----------")
}

func CreateUtilDir() (string, error) {
	fmt.Print("> Creating utility dir: ")
	err := os.Mkdir(tempDirName, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return "", fmt.Errorf("can't create temp dir %s: %v", tempDirName, err)
	}

	osWdPath, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("can't get current executable path: %v", err)
	}

	tempFolderPath := path.Join(osWdPath, tempDirName)
	fmt.Println(tempFolderPath)
	return tempFolderPath, nil
}

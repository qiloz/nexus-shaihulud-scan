package utilEnv

import (
	"fmt"
	"os"
	"path"
)

const tempDirName = ".nx-shaihd-scan"

func CreateUtilDir() (string, error) {
	err := os.Mkdir(tempDirName, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return "", fmt.Errorf("can't create temp dir %s: %v", tempDirName, err)
	}

	osWdPath, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("can't get current executable path: %v", err)
	}

	tempFolderPath := path.Join(osWdPath, tempDirName)
	return tempFolderPath, nil
}

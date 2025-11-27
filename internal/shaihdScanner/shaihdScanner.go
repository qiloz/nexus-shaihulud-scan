package shaihdScanner

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/qiloz/nexus-shaihulud-scan/cmd/utilEnv"
	"github.com/qiloz/nexus-shaihulud-scan/internal/utilityFlags"
)

const CompromisedPkgSource = "https://docs.google.com/spreadsheets/d/16aw6s7mWoGU7vxBciTEZSaR5HaohlBTfVirvI-PypJc/export?format=csv&gid=1289659284"
const CompromisedListPkgFilename = "shaihd-compromised-checklist.csv"

func Init(flags utilityFlags.Flags, repoPkgsFilePath string) error {
	malwareListPath, err := pullCompromisedPkg()
	if err != nil {
		return err
	}

	err = runScanner(repoPkgsFilePath, malwareListPath)
	if err != nil {
		return err
	}

	return nil
}

func pullCompromisedPkg() (string, error) {
	var client = http.Client{}

	fmt.Println("Retrieve SHA1-hulud-compromised packages from: ", CompromisedPkgSource)
	req, err := http.NewRequest("Get", CompromisedPkgSource, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	malwareListPath, err := writePkgDataToFile(buf)

	return malwareListPath, err
}

func writePkgDataToFile(data []byte) (string, error) {
	utilDirPath, err := utilEnv.CreateUtilDir()
	if err != nil {
		return "", err
	}

	var pkgRepoListPath = path.Join(utilDirPath, CompromisedListPkgFilename)

	f, err := os.Create(pkgRepoListPath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return "", err
	}

	fmt.Println("> Malware pkg list assets was written to: " + pkgRepoListPath)

	return pkgRepoListPath, nil
}

func runScanner(repoPkgListPath string, malwarePkgListPath string) error {
	repoPkgsFile, err := os.Open(repoPkgListPath)
	if err != nil {
		return err
	}
	defer repoPkgsFile.Close()

	malwareFile, err := os.Open(malwarePkgListPath)
	if err != nil {
		return err
	}
	defer malwareFile.Close()

	repoLineScanner := bufio.NewScanner(repoPkgsFile)

	fmt.Println("> Comparing repo pkg list with malware list")

	var isMalwareFound bool

	for repoLineScanner.Scan() {
		repoLine := repoLineScanner.Text()

		_, err := malwareFile.Seek(0, 0)
		if err != nil {
			return fmt.Errorf("failed to seek malware file: %w", err)
		}

		malwareLineScanner := bufio.NewScanner(malwareFile)

		for malwareLineScanner.Scan() {
			if repoLineScanner.Text() == "Package Name,Version" {
				continue
			}

			malwareLine := malwareLineScanner.Text()
			if repoLine == malwareLine {
				isMalwareFound = true
				fmt.Println("> Malware pkg list asset found: " + repoLineScanner.Text())
			}
		}
	}

	if !isMalwareFound {
		fmt.Println("> OK! Malware pkg assets not found at your repository.")
	}

	return nil
}

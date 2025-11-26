package flagParser

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/qiloz/nexus-shaihulud-scan/internal/utilityFlags"
)

func Init() (utilityFlags.Flags, error) {
	nxUsername := flag.String("u", "", "Nexus username")
	nxPassword := flag.String("p", "", "Nexus password")
	nxInstanceUri := flag.String("uri", "", "Nexus instance uri")
	nxRepoName := flag.String("repo", "", "Nexus repository name for malware scanning")
	flag.Parse()

	if nxUsername == nil || len(*nxUsername) == 0 {
		return utilityFlags.Flags{}, errors.New("nexus username is not specified (-u)")
	}

	if nxPassword == nil || len(*nxPassword) == 0 {
		return utilityFlags.Flags{}, errors.New("nexus password  is not specified (-p)")
	}

	if nxInstanceUri == nil || len(*nxInstanceUri) == 0 {
		return utilityFlags.Flags{}, errors.New("nexus instance URI is not specified (-uri)")
	}
	var err error
	*nxInstanceUri, err = normalizeUri(nxInstanceUri)
	if err != nil {
		return utilityFlags.Flags{}, err
	}

	if nxRepoName == nil || len(*nxRepoName) == 0 {
		return utilityFlags.Flags{}, errors.New("target nexus repository name is not specified (-repo)")
	}
	return utilityFlags.Flags{
		NxUsername: *nxUsername, NxPassword: *nxPassword, NxInstanceUri: *nxInstanceUri, NxRepoName: *nxRepoName,
	}, nil
}

func normalizeUri(uri *string) (string, error) {
	if len(*uri) == 0 {
		return "", errors.New("nexus repository uri cannot be empty")
	}

	if (*uri)[len(*uri)-1] == '/' {
		*uri = (*uri)[:len(*uri)-1]
	}

	if !strings.Contains(*uri, "http://") || !strings.Contains(*uri, "https://") {
		*uri = fmt.Sprintf("https://%s", *uri)
	}
	return *uri, nil
}

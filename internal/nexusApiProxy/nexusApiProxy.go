package nexusApiProxy

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/qiloz/nexus-shaihulud-scan/internal/utilityFlags"
)

type NxAssetsResult struct {
	Items             []NxAssetItem `json:"items,omitempty"`
	ContinuationToken string        `json:"continuationToken,omitempty"`
}

type NxAssetItem struct {
	Npm NxNpmData `json:"npm"`
}

type NxNpmData struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func GetNxRepoAssets(flags utilityFlags.Flags, continuationToken string) (NxAssetsResult, error) {
	client := &http.Client{}

	var requestUri = flags.NxInstanceUri + "/service/rest/v1/search/assets?repository=" + flags.NxRepoName
	if continuationToken != "" {
		requestUri += "&continuationToken=" + continuationToken
	}

	req, err := http.NewRequest("GET", requestUri, nil)
	if err != nil {
		return NxAssetsResult{}, err
	}

	if flags.NxUsername != "" && flags.NxPassword != "" {
		req.SetBasicAuth(flags.NxUsername, flags.NxPassword)
	}

	resp, err := client.Do(req)
	if err != nil {
		return NxAssetsResult{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return NxAssetsResult{}, err
	}

	var assetData NxAssetsResult
	err = json.Unmarshal(body, &assetData)
	if err != nil {
		return NxAssetsResult{}, err
	}

	return assetData, nil
}

package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type WhitelistUpdateData struct {
	Add    []string `json:"add,omitempty"`
	Remove []string `json:"remove,omitempty"`
	APIKey string   `json:"apiKey"`
}

type WhitelistUpdateResult struct {
	Status    string `json:"status"`
	ErrorCode string `json:"errorCode"`
}

func UpdateWhitelist(data *WhitelistUpdateData) *WhitelistUpdateResult {
	updateData, _ := json.Marshal(data)

	response := patch(whitelistPatchUrl, updateData)

	defer response.Body.Close()

	resultJson, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln("An error occurred while parsing the IP whitelist update result", err)
	}

	var result WhitelistUpdateResult
	err = json.Unmarshal(resultJson, &result)

	if err != nil {
		log.Fatalln("An error occurred while parsing the IP whitelist update result", err)
	}

	return &result
}

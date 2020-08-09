package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type LoginData struct {
	Username *string `json:"username,omitempty"`
	UserID   *int    `json:"userID,omitempty"`
	Password *string `json:"password"`
}

type LoginResult struct {
	Status    string `json:"status"`
	APIKey    string `json:"apiKey"`
	Expires   int    `json:"expires"`
	SessionID string `json:"sessionID"`
	Cookie    bool   `json:"cookie"`
	ErrorCode string `json:"errorCode"`
}

func Login(data *LoginData) (result *LoginResult) {
	loginJson, _ := json.Marshal(data)

	println(string(loginJson))

	response := post(loginUrl, loginJson)

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln("An error occurred while parsing a response", err)
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalln("An error occurred while parsing a response", err)
	}

	return result
}

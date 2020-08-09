package updater

import (
	"errors"
	"fmt"
	"github.com/StevenKGER/myreflex-ipupdate/api"
	"log"
	"time"
)

func Run(username *string, userID *int, password *string) {
	oldIP := getOldIP()
	var currentIP string

	for {
	ip:
		for {
			currentIP = api.GetCurrentIP()
			if currentIP != oldIP {
				break ip
			}

			log.Println("Waiting 5 minutes until next check, if the IP address has changed....")
			time.Sleep(5 * time.Minute)
		}

		log.Println("Logging in....")
		result := login(username, userID, password)

		log.Println("Updating whitelist...")
		if currentIP == "" { // initial IP has to be inserted
			updateResult := updateWhitelist(nil, []string{oldIP}, result.APIKey)
			println(updateResult.ErrorCode)

			if updateResult == nil {
				oldIP = ""
				log.Println("Whitelist couldn't be updated, trying again in 5 minutes.")
			} else {
				log.Println("Whitelist updated!")
			}
		} else {
			updateResult := updateWhitelist([]string{currentIP}, []string{oldIP}, result.APIKey)

			if updateResult == nil {
				log.Print("Whitelist couldn't be updated, trying again in 5 minutes.")
			} else {
				log.Println("Whitelist updated!")
			}

			oldIP = currentIP
			saveNewIP(currentIP)
		}
	}
}

func updateWhitelist(currentIP []string, oldIP []string, apiKey string) *api.WhitelistUpdateResult {
	updateResult := api.UpdateWhitelist(&api.WhitelistUpdateData{
		Add:    currentIP,
		Remove: oldIP,
		APIKey: apiKey,
	})
	return updateResult
}

func login(username *string, userID *int, password *string) *api.LoginResult {
	result := api.Login(&api.LoginData{
		Username: username,
		UserID:   userID,
		Password: password,
	})

	if result.Status != "success" {
		panic(errors.New(fmt.Sprintf("invalid login information provided, please check your login information (%s)",
			result.ErrorCode)))
	}
	return result
}

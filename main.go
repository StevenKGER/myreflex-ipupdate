package main

import (
	"errors"
	"flag"
	"github.com/StevenKGER/myreflex-ipupdate/updater"
	"log"
	"os"
	"strconv"
)

func main() {
	username, userID, password := initializeCommandArgs()

	log.Println("Hello!")

	if *username == "" && *userID == 0 {
		userNameFromEnv := os.Getenv("MYREFLEX_USERNAME")
		userIDFromEnv := os.Getenv("MYREFLEX_USERID")

		if userNameFromEnv == "" {
			if userIDFromEnv == "" {
				panic(errors.New("no username or userID provided, please see the help page"))
			}

			userIDFromEnvAsInt, err := strconv.Atoi(userIDFromEnv)
			if err != nil {
				panic(err)
			}

			userID = &userIDFromEnvAsInt
		}

		username = &userNameFromEnv
	}

	if *userID == 0 {
		userID = nil
	}

	if *username == "" {
		username = nil
	}

	if *password == "" {
		passwordFromEnv := os.Getenv("MYREFLEX_PASSWORD")

		if passwordFromEnv == "" {
			panic(errors.New("no password provided, please see the help page"))
		}

		password = &passwordFromEnv
	}

	updater.Run(username, userID, password)

}

func initializeCommandArgs() (*string, *int, *string) {
	username := flag.String("username", "",
		"username for authenticating to MyReflex API services; can also be set as env MYREFLEX_USERNAME")
	userID := flag.Int("userID", 0,
		"userID for authenticating to MyReflex API services; can also be set as env MYREFLEX_USERID")
	password := flag.String("password", "",
		"password for authenticating to MyReflex API services; can also be set as env MYREFLEX_PASSWORD")

	flag.Parse()
	return username, userID, password
}

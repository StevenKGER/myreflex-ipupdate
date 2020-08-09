package updater

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const header = "# Do not change this file manually except you want some weird errors when updating the whitelist.\n"
const fileName = "oldIP.txt"

func saveNewIP(ip string) {
	if err := ioutil.WriteFile(fileName, []byte(header+ip), 0644); err != nil {
		log.Fatalln("An error occurred while saving the new ip", err)
	}
}

func getOldIP() string {
	byteContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatalln("An error occurred while reading the old IP", err)
		}

		return ""
	}

	var ip string

	content := string(byteContent)
	if strings.Contains(content, header) {
		ip = strings.TrimSpace(strings.Split(content, header)[1])
	} else {
		ip = strings.TrimSpace(content)
	}

	return ip
}

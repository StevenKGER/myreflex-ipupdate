package api

import (
	"io/ioutil"
	"log"
)

func GetCurrentIP() string {
	response := get(ipUrl)

	if response.StatusCode != 200 {
		log.Fatalln("An server-side error occurred while parsing the IP result")
		return ""
	}

	defer response.Body.Close()

	ip, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln("An error occurred while parsing the IP result", err)
	}

	return string(ip)
}

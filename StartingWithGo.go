package main

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Users  []string
	Groups []string
}

func main() {
	log.Println("In main Function")
	file, _ := os.Open("config/conf.json")
	log.Println("file ", file)
	decoder := json.NewDecoder(file)
	//log.Println("decoder ", decoder)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Println("error:", err)
	}
	log.Println(configuration.Groups)
	log.Println(configuration.Users)
	log.Println("Finished")
}

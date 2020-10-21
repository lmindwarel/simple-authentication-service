package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	utils "github.com/lmindwarel/quizzbox-utils"
	"quizzbox.fr/authentificator/authenticator"
	"quizzbox.fr/authentificator/datastore"
)

// Config is the core configuration
type Config struct {
	LogPath       string               `json:"logPath"`
	Datastore     datastore.Config     `json:"datastore"`
	Authenticator authenticator.Config `json:"authenticator"`
}

func main() {
	var err error

	configFileName := "config.json"
	if len(os.Args) > 1 {
		configFileName = os.Args[1]
	}

	_, err = os.Stat(configFileName)
	if os.IsNotExist(err) {
		panic(errors.New("please provide config.json or give the path in arg"))
	}

	fmt.Printf("Reading config file...")
	configFile, err := os.Open(configFileName)
	if err != nil {
		panic(err)
	}

	var config Config
	parser := json.NewDecoder(configFile)
	if err = parser.Decode(&config); err != nil {
		panic(err)
	}
	fmt.Printf("ok\n")

	fmt.Printf("Initialize logger...")
	utils.InitLogger(config.LogPath)
	fmt.Printf("ok\n")

	fmt.Printf("Initialize datastore...")
	ds, err := datastore.New(config.Datastore)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ok\n")

	fmt.Printf("Initialize authenticator...")
	a := authenticator.New(config.Authenticator, ds)
	fmt.Printf("ok\n")

	fmt.Printf("Starting authenticator...")

	err = a.StartServer()
	if err != nil {
		panic(err)
	}
	fmt.Printf("ok\n")

	utils.Standby()
}

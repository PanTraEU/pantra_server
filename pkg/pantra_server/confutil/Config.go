package configUtil

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	confFiles = []string{
		"./pantra_server.json",
		"./etc/pantra_server/pantra_server.json",
		"/etc/pantra_server/pantra_server.json",
	}
)

type Configuration struct {
	DbPath   string `json:"db_path"`
	DataPath string `json:"data_path"`
}

func getConfFile() (string, error) {

	for index, element := range confFiles {
		log.Debug(index, "=>", element)
		if _, err := os.Stat(element); err == nil {
			log.Info("will use config file: ", element)
			return element, nil
		}
	}

	return "", fmt.Errorf("no config file found: %s", confFiles)
}

func GetConfig() Configuration {
	cFile, cerr := getConfFile()
	if cerr != nil {
		panic(cerr.Error())
	}

	file, ferr := os.Open(cFile)
	defer file.Close()
	if ferr != nil {
		panic(ferr.Error())
	}

	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}

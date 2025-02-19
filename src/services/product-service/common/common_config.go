package common

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Configuration struct {
	Port                string `json:"port"`
	EnableGinConsoleLog bool   `json:"enableConsoleLog"`
	EnableGinFileLog    bool   `json:"enableGinFileLog"`
	LogFileName         string `json:"logFileName"`
	LogMaxSize          int    `json:"logMaxSize"`
	LogMaxBackups       int    `json:"logMaxBackups"`
	LogMaxAge           int    `json:"logMaxAge"`
	MgAddrs             string `json:"mgAddrs"`
	MgDbName            string `json:"mgDbName"`
	MgDbUserName        string `json:"mgUsername"`
	MgDbPassword        string `json:"mgDbPassword"`
	JwtSecretPassword   string `json:"jwtSecretPassword"`
	Isser               string `json:"isser"`
}

var Config *Configuration

func LoadConfig() error {
	file, err := os.Open("../pkg/config/config.json")

	if err != nil{
		return err
	}

	Config = new(Configuration)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)

	if err != nil {
		return err
	}

	logFile := &lumberjack.Logger{
		Filename: 	Config.LogFileName,
		MaxSize: 	Config.LogMaxSize,
		MaxBackups: Config.LogMaxBackups,
		MaxAge: 	Config.LogMaxAge,
	}

	log.SetOutput(logFile)
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})
	return nil
}
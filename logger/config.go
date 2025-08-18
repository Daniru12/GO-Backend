package logger

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type LogConfig struct {
	Level         string `yaml:"level" json:"level"`
	RemoteLogging bool   `yaml:"remote_logging" json:"remote_logging"`
	FilePath      bool   `yaml:"file_path_enabled" json:"file_path_enabled"`
}

var logConfig *LogConfig

func init() {
	loadConfig()

	if os.Getenv(`LOG_LEVEL`) != `` {
		logConfig.Level = os.Getenv(`LOG_LEVEL`)
	}

	if os.Getenv(`LOG_FILE_PATH`) != `` && os.Getenv(`LOG_FILE_PATH`) == `1` {
		logConfig.FilePath = true
	}

}

// loadConfig Load logging configurations
func loadConfig() {

	file, err := ioutil.ReadFile(`config/logger.yaml`)
	if err != nil {
		log.Fatalln(`core/logger: Cannot open config file `, err)
	}

	logConfig = new(LogConfig)

	err = yaml.Unmarshal(file, &logConfig)
	if err != nil {
		log.Fatalln(`core/logger: Cannot decode config file `, err)
	}
}

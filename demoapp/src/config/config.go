package config

import (
	"encoding/json"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Config struct {
    Server struct {
        Port string `yaml:"port"`
        Timeout time.Duration `yaml:"timeout"`
    } `yaml:"server"`
}

func New(log *logrus.Logger) (*Config) {
    config := &Config{}

    file, err := os.Open("config/config.yaml")
    if err != nil {
        log.Fatal("Error opening config file:", err)
    }
    defer file.Close()

    d := yaml.NewDecoder(file)
    if err := d.Decode(&config); err != nil {
        log.Fatal("Error decoding config file:", err)
    }

    configJSON, _ := json.MarshalIndent(config, "", "  ")
    log.Debugf("Log config structure for convenience:\n%s", string(configJSON))

    return config
}
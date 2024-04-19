package config

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"

	toml "github.com/BurntSushi/toml"
)

type Config struct {
	DB_PATH   string `toml:"DB_PATH"`
	PORT      string `toml:"PORT"`
	MEDIA_DIR string `toml:"MEDIA_DIR"`
	API_KEY   string `toml:"API_KEY"`
}

func LoadConfig() (Config, error) {
	path := "config.toml"
	if ConfigFileExists(path) {
		config, err := ReadConfig(path)
		return config, err
	}

	log.Println("No configuration file found")
	log.Println("Loading default configuration")

	// return default config
	defaultConfig := Config{
		DB_PATH:   "./database.db",
		PORT:      ":3002",
		MEDIA_DIR: "./temp",
		API_KEY:   "",
	}

	err := Save(defaultConfig, path)

	return defaultConfig, err

}

func ConfigFileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func ReadConfig(configFilePath string) (Config, error) {
	var config Config
	content, err := os.ReadFile(configFilePath)

	if err != nil {
		return config, err
	}

	_, err = toml.Decode(string(content), &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func Save(config Config, path string) error {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(config); err != nil {
		return err
	}

	err := os.WriteFile(path, buf.Bytes(), os.ModePerm)

	return err

}

func findAvailablePort() string {
	// create listener on a random port
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	// get system assigned port
	addr := listener.Addr().(*net.TCPAddr)
	return fmt.Sprintf(":%d", addr.Port)
}

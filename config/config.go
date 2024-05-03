package config

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"

	toml "github.com/BurntSushi/toml"
)

var CONFIG_PATH = "config.toml"

type Config struct {
	DB_PATH      string `toml:"DB_PATH"`
	PORT         string `toml:"PORT"`
	MEDIA_DIR    string `toml:"MEDIA_DIR"`
	SUBTITLE_DIR string `toml:"SUBTITLE_DIR"`
	API_KEY      string `toml:"API_KEY"`
}

// LoadConfig loads the configuration from the specified file path.
// If the configuration file exists, it reads the configuration from the file.
// If the file does not exist, it creates a default configuration and saves it to the file.
// The function returns the loaded or default configuration and any error that occurred.
func LoadConfig() (Config, error) {
	path := "config.toml"
	if configFileExists(path) {
		config, err := readConfig(path)
		return config, err
	}

	log.Println("No configuration file found")
	log.Println("Loading default configuration")

	// return default config
	defaultConfig := Config{
		DB_PATH:      "./database.db",
		PORT:         ":3002",
		MEDIA_DIR:    "./media",
		SUBTITLE_DIR: "./subtitles",
		API_KEY:      "",
	}

	err := defaultConfig.Save(CONFIG_PATH)

	return defaultConfig, err

}

func configFileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func readConfig(configFilePath string) (Config, error) {
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

func (c *Config) Save(configFilePath string) error {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(c); err != nil {
		return err
	}

	err := os.WriteFile(configFilePath, buf.Bytes(), os.ModePerm)

	return err

}

func findAvailablePort() string {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	addr := listener.Addr().(*net.TCPAddr)
	return fmt.Sprintf(":%d", addr.Port)
}

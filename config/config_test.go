package config

import (
	"os"
	"testing"
)

func TestCreateConfigFile(t *testing.T) {
	config := Config{
		DB_PATH:      "./testPath",
		PORT:         ":2000",
		API_KEY:      "test_key",
		SUBTITLE_DIR: "test_subtitle_dir",
		MEDIA_DIR:    "test_media_dir",
	}
	tmpFilePath := "tmp.toml"
	if err := config.Save(tmpFilePath); err != nil {
		t.Errorf("Error saving config file: %s", err)
	}

	_, err := os.Stat(tmpFilePath)
	if err != nil {
		t.Errorf("Created file config not found: %s", err)
	}

	err = os.Remove(tmpFilePath)
	if err != nil {
		t.Errorf("Error removing tmp file: %s", err)
	}

}

func TestReadConfigFile(t *testing.T) {
	config := Config{
		DB_PATH:      "./testPath",
		PORT:         ":2000",
		API_KEY:      "test_key",
		SUBTITLE_DIR: "test_subtitle_dir",
		MEDIA_DIR:    "test_media_dir",
	}
	tmpFilePath := "tmp2.toml"
	if err := config.Save(tmpFilePath); err != nil {
		t.Errorf("Error saving config file: %s", err)
	}

	writedConfig, err := readConfig(tmpFilePath)
	if err != nil {
		t.Errorf("Error reading file: %s", err)
	}

	if writedConfig.DB_PATH != config.DB_PATH {
		t.Errorf("Error reading \"DB_PATH\": expect=%s, got=\"%s\"", config.DB_PATH, writedConfig.DB_PATH)
	}
	if writedConfig.API_KEY != config.API_KEY {
		t.Errorf("Error reading \"API_KEY\": expect=%s, got=\"%s\" ", config.API_KEY, writedConfig.API_KEY)
	}
	if writedConfig.PORT != config.PORT {
		t.Errorf("Error reading \"PORT\": expect=%s, got=\"%s\" ", config.PORT, writedConfig.PORT)
	}

	err = os.Remove(tmpFilePath)
	if err != nil {
		t.Errorf("Error removing tmp file: %s", err)
	}

}

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig()
	if err != nil {
		t.Errorf("Error loading config file: %s", err)
	}

	if config.DB_PATH == "" {
		t.Errorf("Error loading DB_PATH: %s", config.DB_PATH)
	}
	if config.PORT == "" {
		t.Errorf("Error loading PORT: %s", config.PORT)
	}

	if config.MEDIA_DIR == "" {
		t.Errorf("Error loading MEDIA_DIR: %s", config.MEDIA_DIR)
	}
	if config.SUBTITLE_DIR == "" {
		t.Errorf("Error loading SUBTITLE_DIR: %s", config.SUBTITLE_DIR)
	}

	if _, err := os.Stat("./config.toml"); err != nil {
		t.Errorf("Error loading config file: %s", err)
	}

	os.Remove("./config.toml")
}

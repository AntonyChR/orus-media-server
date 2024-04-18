package config

import (
	"os"
	"testing"
)

func TestCreateConfigFile(t *testing.T) {
	config := Config{DB_PATH: "./testPath", PORT: ":2000", API_KEY: "test_key"}
	tmpFilePath := "tmp.toml"
	if err := Save(config, tmpFilePath); err != nil {
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
	config := Config{DB_PATH: "./testPath", PORT: ":2000", API_KEY: "test_key"}
	tmpFilePath := "tmp2.toml"
	if err := Save(config, tmpFilePath); err != nil {
		t.Errorf("Error saving config file: %s", err)
	}

	writedConfig, err := ReadConfig(tmpFilePath)
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

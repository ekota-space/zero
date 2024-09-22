package setup

import (
	"encoding/json"
	"errors"
	"io/fs"
	"log"
	"os"
	"path"

	"github.com/ekota-space/zero/pkgs/common"
	"github.com/ekota-space/zero/pkgs/root/db"
)

type SetupConfig struct {
	Running  bool `json:"running"`
	Finished bool `json:"finished"`
}

func writeToJsonFile(data SetupConfig) error {
	jsonByte, err := json.Marshal(data)

	if err != nil {
		return err
	}

	configPath := path.Join(os.TempDir(), "zero", "config.json")

	err = os.MkdirAll(path.Dir(configPath), 0777)

	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, jsonByte, 0644)

	if err != nil {
		return err
	}

	return nil
}

func readJsonFile() (SetupConfig, error) {
	data := SetupConfig{}

	configPath := path.Join(os.TempDir(), "zero", "config.json")

	jsonFile, err := os.ReadFile(configPath)

	if err != nil {
		return data, err
	}

	err = json.Unmarshal(jsonFile, &data)

	if err != nil {
		return data, err
	}

	return data, nil
}

func GlobalSetup() {
	config, err := readJsonFile()

	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		log.Fatalf("❌ Error reading config file: %v", err)
	}

	if config.Running || config.Finished {
		log.Println("🟢 Setup already ran")
		return
	}
	config.Running = true
	err = writeToJsonFile(config)

	if err != nil {
		log.Fatalf("❌ Error writing config file: %v", err)
	}

	log.Println("⚪️ Running migrations")

	common.SetupTestEnvironmentVars()
	db.RunMigrations()

	log.Println("✅ Migration complete")

	config.Finished = true
	config.Running = false

	err = writeToJsonFile(config)

	if err != nil {
		log.Fatalf("❌ Error writing config file: %v", err)
	}
}

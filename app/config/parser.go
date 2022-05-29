package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func Parse(configDir string) *Config {
	// set config directory
	dir := getDir(configDir)

	return &Config{
		AppConfig: parseAppConfig(dir),
	}
}

// parseAppConfig parses application configurations.
func parseAppConfig(dir string) AppConfig {

	cfg := AppConfig{}
	file := read(dir + "app.yaml")

	err := yaml.Unmarshal(file, &cfg)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}
	return cfg
}

// getDir returns config directory path after analysing and correcting.
func getDir(dir string) string {

	// get last char of dir path
	c := dir[len(dir)-1]
	if os.IsPathSeparator(c) {
		return dir
	}

	return dir + string(os.PathSeparator)
}

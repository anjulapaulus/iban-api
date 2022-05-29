package config

// Config is the main config struct that holds all other config structs.
type Config struct {
	AppConfig AppConfig
}

// AppConfig holds application configurations.
type AppConfig struct {
	Port int `yaml:"port"`
}

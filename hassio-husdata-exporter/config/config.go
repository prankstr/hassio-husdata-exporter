package config

import (
	"log/slog"
	"os"
	"strconv"

	"golang.org/x/text/language"
)

var logLevels map[string]slog.Level = map[string]slog.Level{
	"Error": slog.LevelError,
	"Warn":  slog.LevelWarn,
	"Info":  slog.LevelInfo,
	"Debug": slog.LevelDebug,
}

var languages map[string]language.Tag = map[string]language.Tag{
	"Swedish": language.Swedish,
	"English": language.English,
}

type Config struct {
	HusdataHostname string
	PollInterval    int
	Lang            language.Tag
	LogLevel        slog.Level
	Port            int
	ExternalServer  bool
}

// New returns a new Config struct
func New() *Config {
	config := Config{
		HusdataHostname: getEnvOrDefault("HUSDATA_HOSTNAME", ""),
		PollInterval:    getEnvOrDefaultInt("POLL_INTERVAL", 15),
		Lang:            languages[getEnvOrDefault("LANG", "Swedish")],
		LogLevel:        logLevels[getEnvOrDefault("LOG_LEVEL", "Info")],
		Port:            getEnvOrDefaultInt("PORT", 9101),
	}

	if config.HusdataHostname == "" {
		slog.Error("HUSDATA_HOSTNAME cannot be empty")
		os.Exit(1)
	}

	// Handle the debug setting
	extServerStr := getEnvOrDefault("EXTERNAL_SERVER", "false")
	extServer, err := strconv.ParseBool(extServerStr)
	if err == nil {
		config.ExternalServer = extServer
	} else {
		slog.Error("Cannot parse the EXTERNAL_SERVER variable")
	}

	return &config
}

// Simple helper function to read an environment or return a default value
func getEnvOrDefaultInt(key string, defaultVal int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}

	return defaultVal
}

// Simple helper function to read an environment or return a default value
func getEnvOrDefault(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

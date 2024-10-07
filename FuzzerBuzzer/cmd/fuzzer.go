package main

import (
	"fmt"
	"os"

	"Hacktoberfest2024/FuzzerBuzzer/internal/fuzz_logic" // Correct module path
	"Hacktoberfest2024/FuzzerBuzzer/internal/http"       // Correct module path

	// Correct module path

	"gopkg.in/yaml.v2" // For handling YAML configuration
)

// Config structure for holding configuration details
type Config struct {
	TargetURL string            `yaml:"target_url"`
	Headers   map[string]string `yaml:"headers"`
}

func main() {
	// Load configuration
	config := loadConfig("config/config.yaml")
	fmt.Printf("Starting fuzzing on target: %s\n", config.TargetURL)

	// Create an HTTP client
	client := http.NewClient(config.Headers)

	// Generate inputs and start fuzzing
	fuzzer := fuzz_logic.NewFuzzer(client, config.TargetURL)

	// Start the fuzzing process
	fuzzer.Start()
}

// Function to load the configuration from a YAML file
func loadConfig(filePath string) Config {
	var config Config
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		os.Exit(1)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("Error parsing config file: %v\n", err)
		os.Exit(1)
	}
	return config
}

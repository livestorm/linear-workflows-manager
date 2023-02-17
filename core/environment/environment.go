package environment

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	defaultEnvFile = ".env"
	quotesRegex    = regexp.MustCompile(`\A"(.*)"\z`)
	// Use local variables map instead of os.Setenv as global variables
	// from other packages load before this is initialised.
	variables = map[string]string{}
)

func processValue(value string) string {
	// Remove spaces
	processedVal := strings.TrimSpace(value)

	// Check for surrounding quotes and remove
	quotes := quotesRegex.FindStringSubmatch(value)
	if quotes != nil {
		processedVal = strings.Trim(processedVal, "\"")
	}

	return processedVal
}

func processLine(line string) (string, string, error) {
	keyValuePair := strings.SplitN(line, "=", 2)
	if len(keyValuePair) != 2 {
		return "", "", errors.New("Can't separate key from value")
	}

	// Return key and value
	return strings.TrimSpace(keyValuePair[0]), processValue(keyValuePair[1]), nil
}

func parseFile(content []byte) (map[string]string, error) {
	envLines := strings.Split(string(content), "\n")
	envVariables := make(map[string]string)

	for _, line := range envLines {
		// Remove comment lines and empty lines
		if strings.HasPrefix(line, "#") || strings.Index(line, "=") == -1 {
			continue
		}

		// Process line
		key, value, err := processLine(line)
		if err != nil {
			log.Printf("Failed to parse line: %s", line)
			continue
		}

		// Ensure key exists and store (allowing values to be empty in case we want an empty override)
		if key != "" {
			envVariables[key] = value
		}
	}
	return envVariables, nil
}

func loadFile(filename string) error {
	filePath := fmt.Sprintf("%s", filename)

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		// If the env specific file does not exist, we don't want the process to fail.
		if filename != defaultEnvFile {
			return nil
		}
		return err
	}

	envVariables, err := parseFile(fileContent)
	if err != nil {
		return err
	}

	for key, value := range envVariables {
		variables[key] = value
	}
	return nil
}

func loadServerVariables() error {
	serverVariables := os.Environ()
	for _, variable := range serverVariables {
		key, value, err := processLine(variable)
		if err != nil {
			return err
		}
		variables[key] = value
	}
	return nil
}

func Load() error {
	// Load server environment variables
	err := loadServerVariables()
	if err != nil {
		return err
	}

	// Load variables from file
	fileForEnvironment := fmt.Sprintf("%s.env", strings.ToLower(os.Getenv("ENVIRONMENT")))
	fileNames := []string{defaultEnvFile, fileForEnvironment}

	for _, file := range fileNames {
		_ = loadFile(file)
	}
	return nil
}

func init() {
	err := Load()
	if err != nil {
		log.Fatalf("Failed to load env variables. Killing app. Error: %s", err.Error())
	}
}

func Get(name string) string {
	return variables[name]
}

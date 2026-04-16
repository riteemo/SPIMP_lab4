package config

import (
    "bufio"
    "os"
    "strings"
)

type Config struct {
    APIKey string
}

func LoadConfig() (*Config, error) {
    loadEnvFile(".env")

    apiKey := os.Getenv("EXCHANGE_API_KEY")
    if apiKey == "" {
        apiKey = "demo"
    }

    return &Config{APIKey: apiKey}, nil
}

func loadEnvFile(filename string) {
    file, err := os.Open(filename)
    if err != nil {
        return 
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())

        if line == "" || strings.HasPrefix(line, "#") {
            continue
        }

        parts := strings.SplitN(line, "=", 2)
        if len(parts) == 2 {
            key := strings.TrimSpace(parts[0])
            value := strings.TrimSpace(parts[1])
            os.Setenv(key, value)
        }
    }
}
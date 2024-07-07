package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/viper"
)

type Config struct {
	BearerToken string `mapstructure:"bearer_token"`
	APIURL      string `mapstructure:"api_url"`
}

const requestDelay = 1 * time.Second

func loadConfig() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("could not get home directory: %w", err)
	}
	configPath := filepath.Join(homeDir, "sanpush", "config.yaml")

	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("could not read config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("could not unmarshal config: %w", err)
	}

	if config.BearerToken == "" || config.APIURL == "" {
		return nil, errors.New("config file is missing required fields: bearer_token or api_url")
	}

	return &config, nil
}

func validateMessage(message string) error {
	if len(message) == 0 {
		return errors.New("message cannot be empty")
	}
	if len(message) > 600 {
		return errors.New("message cannot exceed 600 characters")
	}
	return nil
}

func sendMessage(message string) error {
	if err := validateMessage(message); err != nil {
		return err
	}

	config, err := loadConfig()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/send-message", config.APIURL)
	escapedMessage := strings.ReplaceAll(message, "\n", "\\n")
	jsonPayload := fmt.Sprintf(`{"message": "%s"}`, escapedMessage)

	payloadBytes := []byte(jsonPayload)

	// Log the payload to debug if it's correctly formatted
	//fmt.Printf("Payload to be sent: %s\n", payloadBytes)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("could not create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+config.BearerToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}

	s := spinner.New(spinner.CharSets[14], 60*time.Millisecond)
	s.Prefix = "\n Sending "
	s.Color("green", "bold")
	s.Start()
	time.Sleep(3 * time.Second)
	s.Stop()

	time.Sleep(requestDelay)

	resp, err := client.Do(req)
	if err != nil {
		s.Stop()
		return fmt.Errorf("could not send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			s.Stop()
			return fmt.Errorf("could not read response body: %w", err)
		}
		s.Stop()
		return fmt.Errorf("unexpected response code: %d, response: %s", resp.StatusCode, string(bodyBytes))
	}

	fmt.Print("\n✅ Message sent successfully...\n\n")
	return nil
}

func reloadPage() error {
	config, err := loadConfig()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/reload", config.APIURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("could not create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+config.BearerToken)

	client := &http.Client{Timeout: 10 * time.Second}

	s := spinner.New(spinner.CharSets[6], 60*time.Millisecond)
	s.Prefix = "\n Reload "
	s.Color("green", "bold")
	s.Start()
	time.Sleep(3 * time.Second)
	s.Stop()

	time.Sleep(requestDelay)

	resp, err := client.Do(req)
	if err != nil {
		s.Stop()
		return fmt.Errorf("could not send request: %w", err)
	}
	defer resp.Body.Close()

	s.Stop()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("could not read response body: %w", err)
		}
		return fmt.Errorf("unexpected response code: %d, response: %s", resp.StatusCode, string(bodyBytes))
	}

	fmt.Print("\n🍪 Page reloaded successfully\n\n")
	return nil
}

func showUsage() {
	fmt.Println(`Usage:
  sanpush send-message "<message>" - Sends a message to the Webhook Server.
  sanpush reload - Reloads the page.

Commands:
  send-message "<message>" - Send a message to the configured API endpoint.
  reload - Reload the page at the configured API endpoint.

Flags:
  -h, --help - Show this help message.

Examples:
  sanpush send-message "Hello world."
  sanpush reload`)
}

func main() {
	if len(os.Args) < 2 {
		showUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "send-message":
		if len(os.Args) < 3 {
			fmt.Println("Error: No message provided.")
			showUsage()
			return
		}
		message := os.Args[2]
		if err := sendMessage(message); err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
	case "reload":
		if err := reloadPage(); err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
	case "-h", "--help":
		showUsage()
	case "version":
		fmt.Println("sanpush version 1.0.0")
	default:
		fmt.Printf("Error: Unknown command '%s'.\n", command)
		showUsage()
	}
}

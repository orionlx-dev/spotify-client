package spotify

import "fmt"

// Config holds configuration for the Spotify API client.
type Config struct {
	// ClientID is the Spotify application client ID.
	ClientID string

	// ClientSecret is the Spotify application client secret.
	ClientSecret string
}

// Validate checks if the configuration is valid.
func (c Config) Validate() error {
	if c.ClientID == "" {
		return fmt.Errorf("ClientID is required")
	}
	if c.ClientSecret == "" {
		return fmt.Errorf("ClientSecret is required")
	}
	return nil
}

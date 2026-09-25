package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	spotify "github.com/orionlx-dev/spotify-client"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Create configuration
	cfg := spotify.Config{
		ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
	}

	// Validate config
	if err := cfg.Validate(); err != nil {
		log.Fatalf("Invalid configuration: %v", err)
	}

	// Create Spotify client
	client := spotify.NewClient(cfg)

	// Example: Get access token
	token, err := client.GetAccessToken()
	if err != nil {
		log.Fatalf("Failed to get access token: %v", err)
	}

	fmt.Printf("✓ Access token obtained successfully!\n")
	fmt.Printf("  Token type: %s\n", token.TokenType)
	fmt.Printf("  Expires in: %d seconds\n\n", token.ExpiresIn)

	// Example: Get a track
	trackID := "11dFghVXANMlKmJXsNCbNl" // "Cut To The Feeling" by Carly Rae Jepsen
	track, err := client.GetTrack(trackID)
	if err != nil {
		log.Fatalf("Failed to get track: %v", err)
	}

	fmt.Printf("✓ Track Information:\n")
	fmt.Printf("  Name: %s\n", track.Name)
	fmt.Printf("  Artists: ")
	for i, artist := range track.Artists {
		if i > 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("%s", artist.Name)
	}
	fmt.Printf("\n  Album: %s\n", track.Album.Name)
	fmt.Printf("  Duration: %d ms (%.2f minutes)\n", track.DurationMs, float64(track.DurationMs)/60000.0)
	fmt.Printf("  Explicit: %v\n", track.Explicit)
	fmt.Printf("  Popularity: %d/100\n\n", track.Popularity)

	// Example: Search for tracks
	searchQuery := "Doxy Miles Davis"
	searchResults, err := client.SearchTracks(searchQuery, 5)
	if err != nil {
		log.Fatalf("Failed to search tracks: %v", err)
	}

	fmt.Printf("✓ Search Results for '%s' (found %d):\n", searchQuery, searchResults.Total)
	for i, t := range searchResults.Items {
		fmt.Printf("  %d. %s by %s\n", i+1, t.Name, t.Artists[0].Name)
	}
}

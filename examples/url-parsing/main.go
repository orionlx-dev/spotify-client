package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	spotify "github.com/orionlx-dev/spotify-client"
)

func main() {
	// Load environment variables
	_ = godotenv.Load("../../.env")

	// Create client
	config := spotify.Config{
		ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
	}
	client := spotify.NewClient(config)

	fmt.Println("=== Spotify URL Parsing Examples ===\n")

	// Example 1: Parse track URL
	trackURL := "https://open.spotify.com/track/11dFghVXANMlKmJXsNCbNl?si=abc123"
	fmt.Printf("Input: %s\n", trackURL)

	trackID, err := spotify.ExtractTrackID(trackURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Extracted Track ID: %s\n", trackID)

	track, err := client.GetTrack(trackID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Track: %s by %s\n\n", track.Name, track.Artists[0].Name)

	// Example 2: Parse track URI
	trackURI := "spotify:track:3n3Ppam7vgaVa1iaRUc9Lp"
	fmt.Printf("Input: %s\n", trackURI)

	trackID2, err := spotify.ExtractTrackID(trackURI)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Extracted Track ID: %s\n", trackID2)

	track2, err := client.GetTrack(trackID2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Track: %s by %s\n\n", track2.Name, track2.Artists[0].Name)

	// Example 3: Parse playlist URL
	playlistURL := "https://open.spotify.com/playlist/37i9dQZF1DXcBWIGoYBM5M"
	fmt.Printf("Input: %s\n", playlistURL)

	playlistID, err := spotify.ExtractPlaylistID(playlistURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Extracted Playlist ID: %s\n", playlistID)

	playlist, err := client.GetPlaylist(playlistID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Playlist: %s\n", playlist.Name)
	fmt.Printf("Total tracks: %d\n\n", playlist.Tracks.Total)

	// Example 4: Generic parsing with ParseSpotifyURL
	inputs := []string{
		"https://open.spotify.com/track/6rqhFgbbKwnb9MLmUQDhG6",
		"spotify:playlist:37i9dQZF1DX0XUsuxWHRQd",
		"https://open.spotify.com/album/4aawyAB9vmqN3uQ7FjRGTy",
	}

	fmt.Println("=== Generic URL Parsing ===")
	for _, input := range inputs {
		resourceType, id, err := spotify.ParseSpotifyURL(input)
		if err != nil {
			log.Printf("Error parsing %s: %v\n", input, err)
			continue
		}
		fmt.Printf("Type: %-10s ID: %s\n", resourceType, id)
	}
	fmt.Println()

	// Example 5: User input simulation
	fmt.Println("=== Simulating User Input ===")
	userInputs := []string{
		"https://open.spotify.com/track/11dFghVXANMlKmJXsNCbNl?si=xyz",
		"11dFghVXANMlKmJXsNCbNl", // Just the ID
		"spotify:track:11dFghVXANMlKmJXsNCbNl",
	}

	for _, input := range userInputs {
		fmt.Printf("\nUser input: %s\n", input)

		resourceType, id, err := spotify.ParseSpotifyURL(input)
		if err != nil {
			log.Printf("Could not parse: %v\n", err)
			continue
		}

		fmt.Printf("Detected: %s with ID %s\n", resourceType, id)

		if resourceType == "track" {
			track, err := client.GetTrack(id)
			if err != nil {
				log.Printf("Failed to fetch track: %v\n", err)
				continue
			}
			fmt.Printf("✓ Found: %s by %s\n", track.Name, track.Artists[0].Name)
		} else if resourceType == "playlist" {
			playlist, err := client.GetPlaylist(id)
			if err != nil {
				log.Printf("Failed to fetch playlist: %v\n", err)
				continue
			}
			fmt.Printf("✓ Found: %s (%d tracks)\n", playlist.Name, playlist.Tracks.Total)
		}
	}
}

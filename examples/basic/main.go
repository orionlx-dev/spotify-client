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
	_ = godotenv.Load("../../.env")

	// Create configuration
	config := spotify.Config{
		ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
	}

	// Validate configuration
	if err := config.Validate(); err != nil {
		log.Fatalf("Invalid configuration: %v", err)
	}

	// Create Spotify client
	client := spotify.NewClient(config)

	// Example 1: Get a track by ID
	fmt.Println("=== Example 1: Get Track ===")
	trackID := "11dFghVXANMlKmJXsNCbNl"
	track, err := client.GetTrack(trackID)
	if err != nil {
		log.Fatalf("Failed to get track: %v", err)
	}

	fmt.Printf("Track: %s\n", track.Name)
	fmt.Printf("Artists: ")
	for i, artist := range track.Artists {
		if i > 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("%s", artist.Name)
	}
	fmt.Printf("\nAlbum: %s\n", track.Album.Name)
	fmt.Printf("Duration: %d ms\n", track.DurationMs)
	fmt.Printf("Popularity: %d/100\n", track.Popularity)
	fmt.Printf("Explicit: %v\n", track.Explicit)
	if track.PreviewURL != nil {
		fmt.Printf("Preview URL: %s\n", *track.PreviewURL)
	}
	fmt.Println()

	// Example 2: Search for tracks
	fmt.Println("=== Example 2: Search Tracks ===")
	query := "track:Doxy artist:Miles Davis"
	results, err := client.SearchTracks(query, 5)
	if err != nil {
		log.Fatalf("Failed to search tracks: %v", err)
	}

	fmt.Printf("Found %d results for query: %s\n", results.Total, query)
	for i, track := range results.Items {
		fmt.Printf("%d. %s by %s (Album: %s)\n",
			i+1,
			track.Name,
			track.Artists[0].Name,
			track.Album.Name,
		)
	}
	fmt.Println()

	// Example 3: Get a playlist
	fmt.Println("=== Example 3: Get Playlist ===")
	playlistID := "37i9dQZF1DXcBWIGoYBM5M" // Today's Top Hits
	playlist, err := client.GetPlaylist(playlistID)
	if err != nil {
		log.Fatalf("Failed to get playlist: %v", err)
	}

	fmt.Printf("Playlist: %s\n", playlist.Name)
	fmt.Printf("Owner: %s\n", *playlist.Owner.DisplayName)
	fmt.Printf("Total tracks: %d\n", playlist.Tracks.Total)
	if playlist.Description != nil {
		fmt.Printf("Description: %s\n", *playlist.Description)
	}
	fmt.Println()

	// Example 4: Get all tracks from a playlist
	fmt.Println("=== Example 4: Get Playlist Tracks ===")
	tracks, err := client.GetPlaylistTracks(playlistID)
	if err != nil {
		log.Fatalf("Failed to get playlist tracks: %v", err)
	}

	fmt.Printf("Retrieved %d tracks from playlist\n", len(tracks))
	fmt.Println("First 5 tracks:")
	for i, item := range tracks {
		if i >= 5 {
			break
		}
		if item.Track != nil {
			fmt.Printf("%d. %s by %s\n",
				i+1,
				item.Track.Name,
				item.Track.Artists[0].Name,
			)
		}
	}
}

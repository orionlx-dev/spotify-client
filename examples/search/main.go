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

	// Search examples
	searches := []struct {
		name  string
		query string
	}{
		{"Simple search", "Never Gonna Give You Up"},
		{"Track and artist", "track:Bohemian Rhapsody artist:Queen"},
		{"Multiple artists", "artist:Beatles artist:Stones"},
		{"Year filter", "year:2020-2023 genre:rock"},
		{"Album search", "album:Abbey Road artist:Beatles"},
	}

	for _, s := range searches {
		fmt.Printf("\n=== %s ===\n", s.name)
		fmt.Printf("Query: %s\n", s.query)

		results, err := client.SearchTracks(s.query, 3)
		if err != nil {
			log.Printf("Search failed: %v\n", err)
			continue
		}

		fmt.Printf("Found %d total results\n", results.Total)
		for i, track := range results.Items {
			artistNames := ""
			for j, artist := range track.Artists {
				if j > 0 {
					artistNames += ", "
				}
				artistNames += artist.Name
			}
			fmt.Printf("%d. %s - %s (%s, %s)\n",
				i+1,
				track.Name,
				artistNames,
				track.Album.Name,
				track.Album.ReleaseDate,
			)
		}
	}
}

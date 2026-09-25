package spotify

import (
	"fmt"
	"strings"
)

// ParseSpotifyURL extracts the resource type and ID from a Spotify URL or URI.
//
// Supported formats:
//   - URLs: https://open.spotify.com/track/11dFghVXANMlKmJXsNCbNl
//   - URIs: spotify:track:11dFghVXANMlKmJXsNCbNl
//
// Returns the resource type (e.g., "track", "playlist") and the ID.
//
// Example:
//
//	resourceType, id, err := spotify.ParseSpotifyURL("https://open.spotify.com/track/11dFghVXANMlKmJXsNCbNl")
//	// resourceType = "track", id = "11dFghVXANMlKmJXsNCbNl"
func ParseSpotifyURL(input string) (resourceType string, id string, err error) {
	input = strings.TrimSpace(input)

	if input == "" {
		return "", "", fmt.Errorf("empty input")
	}

	// Handle Spotify URI format: spotify:track:11dFghVXANMlKmJXsNCbNl
	if strings.HasPrefix(input, "spotify:") {
		parts := strings.Split(input, ":")
		if len(parts) < 3 {
			return "", "", fmt.Errorf("invalid spotify URI format")
		}
		return parts[1], parts[2], nil
	}

	// Handle URL format: https://open.spotify.com/track/11dFghVXANMlKmJXsNCbNl
	if strings.Contains(input, "open.spotify.com") {
		// Remove query parameters and fragments
		if idx := strings.Index(input, "?"); idx != -1 {
			input = input[:idx]
		}
		if idx := strings.Index(input, "#"); idx != -1 {
			input = input[:idx]
		}

		parts := strings.Split(input, "/")
		if len(parts) < 2 {
			return "", "", fmt.Errorf("invalid spotify URL format")
		}

		// Extract type and ID from the end of URL
		id = parts[len(parts)-1]
		resourceType = parts[len(parts)-2]

		if id == "" || resourceType == "" {
			return "", "", fmt.Errorf("could not extract type and ID from URL")
		}

		return resourceType, id, nil
	}

	// If it's just an ID (alphanumeric string), assume it's a track ID
	// This is a convenience feature
	if isValidSpotifyID(input) {
		return "track", input, nil
	}

	return "", "", fmt.Errorf("unrecognized format: expected Spotify URL, URI, or ID")
}

// ExtractTrackID extracts the track ID from a Spotify track URL, URI, or ID.
// This is a convenience wrapper around ParseSpotifyURL specifically for tracks.
//
// Example:
//
//	id, err := spotify.ExtractTrackID("https://open.spotify.com/track/11dFghVXANMlKmJXsNCbNl?si=xyz")
//	// id = "11dFghVXANMlKmJXsNCbNl"
func ExtractTrackID(input string) (string, error) {
	resourceType, id, err := ParseSpotifyURL(input)
	if err != nil {
		return "", err
	}

	if resourceType != "track" {
		return "", fmt.Errorf("expected track URL/URI, got %s", resourceType)
	}

	return id, nil
}

// ExtractPlaylistID extracts the playlist ID from a Spotify playlist URL, URI, or ID.
// This is a convenience wrapper around ParseSpotifyURL specifically for playlists.
//
// Example:
//
//	id, err := spotify.ExtractPlaylistID("https://open.spotify.com/playlist/37i9dQZF1DXcBWIGoYBM5M")
//	// id = "37i9dQZF1DXcBWIGoYBM5M"
func ExtractPlaylistID(input string) (string, error) {
	resourceType, id, err := ParseSpotifyURL(input)
	if err != nil {
		return "", err
	}

	if resourceType != "playlist" {
		return "", fmt.Errorf("expected playlist URL/URI, got %s", resourceType)
	}

	return id, nil
}

// isValidSpotifyID checks if a string looks like a valid Spotify ID.
// Spotify IDs are typically 22 characters long and contain alphanumeric characters.
func isValidSpotifyID(s string) bool {
	if len(s) != 22 {
		return false
	}

	for _, c := range s {
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')) {
			return false
		}
	}

	return true
}

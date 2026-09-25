// Package spotify provides a client for the Spotify Web API.
//
// This package allows you to interact with the Spotify Web API to retrieve
// information about tracks, playlists, and search for music content.
//
// Example usage:
//
//	cfg := spotify.Config{
//		ClientID:     "your-client-id",
//		ClientSecret: "your-client-secret",
//	}
//	client := spotify.NewClient(cfg)
//
//	// Get a track
//	track, err := client.GetTrack("11dFghVXANMlKmJXsNCbNl")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Track: %s by %s\n", track.Name, track.Artists[0].Name)
//
//	// Search for tracks
//	results, err := client.SearchTracks("Doxy Miles Davis", 10)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Get a playlist
//	playlist, err := client.GetPlaylist("37i9dQZF1DXcBWIGoYBM5M")
//	if err != nil {
//		log.Fatal(err)
//	}
package spotify

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

const (
	spotifyAuthURL = "https://accounts.spotify.com/api/token"
	spotifyAPIURL  = "https://api.spotify.com/v1"
)

// Client represents an HTTP client for the Spotify API.
type Client struct {
	config     Config
	httpClient *http.Client

	// Token management
	tokenMu     sync.RWMutex
	accessToken string
	tokenExpiry time.Time
}

// NewClient creates a new Spotify API client with the provided configuration.
func NewClient(config Config) *Client {
	if err := config.Validate(); err != nil {
		panic(fmt.Sprintf("invalid config: %v", err))
	}

	return &Client{
		config: config,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// NewClientWithHTTPClient creates a new Spotify API client with a custom HTTP client.
func NewClientWithHTTPClient(config Config, httpClient *http.Client) *Client {
	if err := config.Validate(); err != nil {
		panic(fmt.Sprintf("invalid config: %v", err))
	}

	return &Client{
		config:     config,
		httpClient: httpClient,
	}
}

// ============================================================
// Authentication
// ============================================================

// GetAccessToken retrieves a new access token using Client Credentials Flow.
// This method is automatically called when needed, but can be called manually
// to pre-fetch the token.
func (c *Client) GetAccessToken() (*TokenResponse, error) {
	// Prepare credentials
	auth := base64.StdEncoding.EncodeToString(
		[]byte(c.config.ClientID + ":" + c.config.ClientSecret),
	)

	// Prepare request body
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", spotifyAuthURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create auth request: %w", err)
	}

	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute auth request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("auth failed with status %d: %s", resp.StatusCode, string(body))
	}

	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, fmt.Errorf("failed to decode token response: %w", err)
	}

	// Cache token with 5 minute buffer before expiry
	c.tokenMu.Lock()
	c.accessToken = tokenResp.AccessToken
	c.tokenExpiry = time.Now().Add(time.Duration(tokenResp.ExpiresIn-300) * time.Second)
	c.tokenMu.Unlock()

	return &tokenResp, nil
}

// ensureToken ensures we have a valid access token.
func (c *Client) ensureToken() error {
	c.tokenMu.RLock()
	hasValidToken := c.accessToken != "" && time.Now().Before(c.tokenExpiry)
	c.tokenMu.RUnlock()

	if hasValidToken {
		return nil
	}

	_, err := c.GetAccessToken()
	return err
}

// ============================================================
// HTTP Helpers
// ============================================================

// doRequest performs an authenticated HTTP request to Spotify API.
func (c *Client) doRequest(method, endpoint string, body io.Reader) (*http.Response, error) {
	if err := c.ensureToken(); err != nil {
		return nil, fmt.Errorf("failed to ensure token: %w", err)
	}

	url := spotifyAPIURL + endpoint
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	c.tokenMu.RLock()
	req.Header.Set("Authorization", "Bearer "+c.accessToken)
	c.tokenMu.RUnlock()

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	return resp, nil
}

// handleResponse processes API response and handles errors.
func handleResponse(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var errResp ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			body, _ := io.ReadAll(resp.Body)
			return fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(body))
		}
		return fmt.Errorf("API error: %s (status %d)", errResp.Error.Message, errResp.Error.Status)
	}

	if v != nil {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

// ============================================================
// Track Methods
// ============================================================

// GetTrack retrieves track information by Spotify track ID.
//
// Example:
//
//	track, err := client.GetTrack("11dFghVXANMlKmJXsNCbNl")
func (c *Client) GetTrack(trackID string) (*Track, error) {
	endpoint := fmt.Sprintf("/tracks/%s", trackID)
	resp, err := c.doRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var track Track
	if err := handleResponse(resp, &track); err != nil {
		return nil, fmt.Errorf("failed to get track: %w", err)
	}

	return &track, nil
}

// SearchTracks searches for tracks by query string.
// The query can include filters like: "track:Doxy artist:Miles Davis".
//
// The limit parameter specifies the maximum number of results (1-50).
// If limit is 0 or negative, a default of 10 is used.
//
// Example:
//
//	results, err := client.SearchTracks("Doxy Miles Davis", 10)
func (c *Client) SearchTracks(query string, limit int) (*SearchTracksResult, error) {
	if limit <= 0 || limit > 50 {
		limit = 10 // default
	}

	params := url.Values{}
	params.Set("q", query)
	params.Set("type", "track")
	params.Set("limit", fmt.Sprintf("%d", limit))

	endpoint := fmt.Sprintf("/search?%s", params.Encode())
	resp, err := c.doRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var searchResp SearchResponse
	if err := handleResponse(resp, &searchResp); err != nil {
		return nil, fmt.Errorf("failed to search tracks: %w", err)
	}

	if searchResp.Tracks == nil {
		return nil, fmt.Errorf("no tracks found in response")
	}

	return searchResp.Tracks, nil
}

// ============================================================
// Playlist Methods
// ============================================================

// GetPlaylist retrieves playlist information by Spotify playlist ID.
//
// Example:
//
//	playlist, err := client.GetPlaylist("37i9dQZF1DXcBWIGoYBM5M")
func (c *Client) GetPlaylist(playlistID string) (*Playlist, error) {
	endpoint := fmt.Sprintf("/playlists/%s", playlistID)
	resp, err := c.doRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var playlist Playlist
	if err := handleResponse(resp, &playlist); err != nil {
		return nil, fmt.Errorf("failed to get playlist: %w", err)
	}

	return &playlist, nil
}

// GetPlaylistTracks retrieves all tracks from a playlist.
// This method automatically handles pagination and returns all tracks.
//
// Example:
//
//	tracks, err := client.GetPlaylistTracks("37i9dQZF1DXcBWIGoYBM5M")
func (c *Client) GetPlaylistTracks(playlistID string) ([]PlaylistItem, error) {
	var allItems []PlaylistItem
	offset := 0
	limit := 100 // max allowed by Spotify

	for {
		params := url.Values{}
		params.Set("offset", fmt.Sprintf("%d", offset))
		params.Set("limit", fmt.Sprintf("%d", limit))

		endpoint := fmt.Sprintf("/playlists/%s/tracks?%s", playlistID, params.Encode())
		resp, err := c.doRequest("GET", endpoint, nil)
		if err != nil {
			return nil, err
		}

		var tracksResp PlaylistTracks
		if err := handleResponse(resp, &tracksResp); err != nil {
			return nil, fmt.Errorf("failed to get playlist tracks: %w", err)
		}

		allItems = append(allItems, tracksResp.Items...)

		// Check if we need to fetch more
		if tracksResp.Next == nil || len(tracksResp.Items) == 0 {
			break
		}

		offset += limit
	}

	return allItems, nil
}

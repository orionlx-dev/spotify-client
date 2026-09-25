# Spotify Client for Go

A clean, simple, and well-documented Go client library for the Spotify Web API.

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Features

- 🎵 **Track Management** - Get track information, search for tracks
- 📋 **Playlist Support** - Retrieve playlist information and all tracks (with automatic pagination)
- 🔐 **Automatic Authentication** - Handles OAuth2 Client Credentials flow automatically
- ⚡ **Token Management** - Automatic token caching and refresh
- 🔒 **Thread-Safe** - Safe for concurrent use
- 📦 **Zero Dependencies** (except `godotenv` for examples)
- 🎯 **Clean API** - Simple and intuitive interface

## Installation

```bash
go get github.com/your-username/spotify-client
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    
    spotify "github.com/your-username/spotify-client"
)

func main() {
    // Create configuration
    config := spotify.Config{
        ClientID:     "your-client-id",
        ClientSecret: "your-client-secret",
    }
    
    // Create client
    client := spotify.NewClient(config)
    
    // Get a track
    track, err := client.GetTrack("11dFghVXANMlKmJXsNCbNl")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Track: %s by %s\n", track.Name, track.Artists[0].Name)
}
```

## Prerequisites

You need Spotify API credentials:

1. Go to [Spotify Developer Dashboard](https://developer.spotify.com/dashboard)
2. Create an app (it's free!)
3. Get your `Client ID` and `Client Secret`

## Usage

### Configuration

#### Option 1: Direct Configuration

```go
config := spotify.Config{
    ClientID:     "your-client-id",
    ClientSecret: "your-client-secret",
}
client := spotify.NewClient(config)
```

#### Option 2: Environment Variables (Recommended)

Create a `.env` file:

```env
SPOTIFY_CLIENT_ID=your-client-id
SPOTIFY_CLIENT_SECRET=your-client-secret
```

```go
import (
    "os"
    "github.com/joho/godotenv"
    spotify "github.com/your-username/spotify-client"
)

func main() {
    godotenv.Load()
    
    config := spotify.Config{
        ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
        ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
    }
    
    client := spotify.NewClient(config)
}
```

### Get Track by ID

```go
track, err := client.GetTrack("11dFghVXANMlKmJXsNCbNl")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Name: %s\n", track.Name)
fmt.Printf("Artists: %s\n", track.Artists[0].Name)
fmt.Printf("Album: %s\n", track.Album.Name)
fmt.Printf("Duration: %d ms\n", track.DurationMs)
fmt.Printf("Popularity: %d/100\n", track.Popularity)
```

### Search for Tracks

```go
// Simple search
results, err := client.SearchTracks("Never Gonna Give You Up", 10)
if err != nil {
    log.Fatal(err)
}

for _, track := range results.Items {
    fmt.Printf("%s by %s\n", track.Name, track.Artists[0].Name)
}
```

#### Advanced Search with Filters

```go
// Search with filters
queries := []string{
    "track:Doxy artist:Miles Davis",           // Specific track and artist
    "album:Abbey Road artist:Beatles",          // Album search
    "year:2020-2023 genre:rock",                // Year range and genre
    "artist:Queen",                             // Artist only
}

for _, query := range queries {
    results, _ := client.SearchTracks(query, 5)
    // ... process results
}
```

### Get Playlist

```go
playlist, err := client.GetPlaylist("37i9dQZF1DXcBWIGoYBM5M")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Playlist: %s\n", playlist.Name)
fmt.Printf("Owner: %s\n", *playlist.Owner.DisplayName)
fmt.Printf("Total tracks: %d\n", playlist.Tracks.Total)
```

### Get All Tracks from Playlist

The library automatically handles pagination:

```go
// Retrieves ALL tracks, even if playlist has 1000+ tracks
tracks, err := client.GetPlaylistTracks("37i9dQZF1DXcBWIGoYBM5M")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Retrieved %d tracks\n", len(tracks))

for _, item := range tracks {
    if item.Track != nil {
        fmt.Printf("%s by %s\n", 
            item.Track.Name, 
            item.Track.Artists[0].Name,
        )
    }
}
```

### Custom HTTP Client

```go
import (
    "net/http"
    "time"
)

httpClient := &http.Client{
    Timeout: 10 * time.Second,
    // Custom transport, proxies, etc.
}

client := spotify.NewClientWithHTTPClient(config, httpClient)
```

## Extracting IDs from Spotify URLs

Spotify URLs look like:
- Track: `https://open.spotify.com/track/11dFghVXANMlKmJXsNCbNl`
- Playlist: `https://open.spotify.com/playlist/37i9dQZF1DXcBWIGoYBM5M`

The ID is the last part after the last `/`:

```go
import (
    "strings"
)

func extractID(spotifyURL string) string {
    parts := strings.Split(spotifyURL, "/")
    id := parts[len(parts)-1]
    
    // Remove query parameters if present
    if idx := strings.Index(id, "?"); idx != -1 {
        id = id[:idx]
    }
    
    return id
}

// Usage
trackURL := "https://open.spotify.com/track/11dFghVXANMlKmJXsNCbNl?si=xyz"
trackID := extractID(trackURL) // "11dFghVXANMlKmJXsNCbNl"

track, err := client.GetTrack(trackID)
```

## Examples

Check the [`examples/`](./examples) directory for complete working examples:

- [`examples/basic/`](./examples/basic/main.go) - Basic usage of all features
- [`examples/search/`](./examples/search/main.go) - Advanced search examples

To run examples:

```bash
# Create .env file with your credentials
cp .env.example .env
# Edit .env and add your credentials

# Run basic example
go run examples/basic/main.go

# Run search example
go run examples/search/main.go
```

## API Reference

### Types

#### `Config`
Configuration for the Spotify client.

```go
type Config struct {
    ClientID     string
    ClientSecret string
}
```

#### `Client`
Main client for interacting with Spotify API.

**Methods:**
- `NewClient(config Config) *Client`
- `NewClientWithHTTPClient(config Config, httpClient *http.Client) *Client`
- `GetTrack(trackID string) (*Track, error)`
- `SearchTracks(query string, limit int) (*SearchTracksResult, error)`
- `GetPlaylist(playlistID string) (*Playlist, error)`
- `GetPlaylistTracks(playlistID string) ([]PlaylistItem, error)`
- `GetAccessToken() (*TokenResponse, error)` - Usually called automatically

#### `Track`
Represents a Spotify track with full metadata.

```go
type Track struct {
    ID           string
    Name         string
    Artists      []SimplifiedArtist
    Album        SimplifiedAlbum
    DurationMs   int
    Explicit     bool
    Popularity   int
    PreviewURL   *string
    // ... more fields
}
```

#### `Playlist`
Represents a Spotify playlist.

```go
type Playlist struct {
    ID          string
    Name        string
    Description *string
    Owner       User
    Tracks      PlaylistTracks
    Images      []Image
    // ... more fields
}
```

See [`models.go`](./models.go) for complete type definitions.

## Error Handling

The library returns detailed errors:

```go
track, err := client.GetTrack("invalid-id")
if err != nil {
    // Error contains HTTP status and Spotify error message
    log.Printf("Failed to get track: %v", err)
    return
}
```

## Rate Limiting

Spotify API has rate limits (~180 requests per minute). The library doesn't implement rate limiting - handle it in your application:

```go
import (
    "time"
    "golang.org/x/time/rate"
)

limiter := rate.NewLimiter(rate.Every(time.Second), 3) // 3 requests per second

func getTrackWithLimit(client *spotify.Client, id string) (*spotify.Track, error) {
    if err := limiter.Wait(context.Background()); err != nil {
        return nil, err
    }
    return client.GetTrack(id)
}
```

## Thread Safety

The client is thread-safe and can be used concurrently:

```go
var wg sync.WaitGroup
trackIDs := []string{"id1", "id2", "id3"}

for _, id := range trackIDs {
    wg.Add(1)
    go func(trackID string) {
        defer wg.Done()
        track, _ := client.GetTrack(trackID)
        fmt.Println(track.Name)
    }(id)
}

wg.Wait()
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with ❤️ for the Go and Spotify communities
- Uses [Spotify Web API](https://developer.spotify.com/documentation/web-api)
- Inspired by best practices from the Go community

## Support

- 📚 [Spotify API Documentation](https://developer.spotify.com/documentation/web-api)
- 🐛 [Report Issues](https://github.com/your-username/spotify-client/issues)
- 💬 [Discussions](https://github.com/your-username/spotify-client/discussions)

## Roadmap

- [ ] Add support for Albums API
- [ ] Add support for Artists API  
- [ ] Add support for User's library
- [ ] Implement retry logic with exponential backoff
- [ ] Add context support for cancellation
- [ ] Add more comprehensive tests
- [ ] Add CLI tool for quick API exploration

---

Made with 🎵 and Go

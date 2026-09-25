# Spotify Client for Go - Quick Reference

## Installation

```bash
go get github.com/your-username/spotify-client
```

## Basic Usage

```go
import spotify "github.com/your-username/spotify-client"

// Create client
config := spotify.Config{
    ClientID:     "your-client-id",
    ClientSecret: "your-client-secret",
}
client := spotify.NewClient(config)

// Get track
track, _ := client.GetTrack("11dFghVXANMlKmJXsNCbNl")
fmt.Println(track.Name)
```

## Common Operations

### Parse Spotify URLs

```go
// From URL
trackID, _ := spotify.ExtractTrackID("https://open.spotify.com/track/11dFghVXANMlKmJXsNCbNl")

// From URI
playlistID, _ := spotify.ExtractPlaylistID("spotify:playlist:37i9dQZF1DXcBWIGoYBM5M")

// Generic parsing
resourceType, id, _ := spotify.ParseSpotifyURL(input)
```

### Search Tracks

```go
// Simple search
results, _ := client.SearchTracks("Bohemian Rhapsody", 10)

// With filters
results, _ := client.SearchTracks("track:Doxy artist:Miles Davis", 5)
```

### Get Playlist Tracks

```go
// Automatically handles pagination
tracks, _ := client.GetPlaylistTracks("37i9dQZF1DXcBWIGoYBM5M")
for _, item := range tracks {
    fmt.Println(item.Track.Name)
}
```

## Environment Variables

```bash
export SPOTIFY_CLIENT_ID="your-client-id"
export SPOTIFY_CLIENT_SECRET="your-client-secret"
```

Or use `.env` file with `github.com/joho/godotenv`.

## Error Handling

```go
track, err := client.GetTrack("invalid-id")
if err != nil {
    log.Printf("Error: %v", err)
    // Handle error
}
```

## Documentation

- Full documentation: See [README.md](README.md)
- Examples: See [examples/](examples/)
- Contributing: See [CONTRIBUTING.md](CONTRIBUTING.md)
- Changelog: See [CHANGELOG.md](CHANGELOG.md)

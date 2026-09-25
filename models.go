package spotify

// ============================================================
// Authentication Models
// ============================================================

// TokenResponse represents response from Spotify token endpoint.
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"` // seconds
}

// ============================================================
// Track Models
// ============================================================

// Track represents a Spotify track with full metadata.
type Track struct {
	ID           string             `json:"id"`
	Name         string             `json:"name"`
	Artists      []SimplifiedArtist `json:"artists"`
	Album        SimplifiedAlbum    `json:"album"`
	DurationMs   int                `json:"duration_ms"`
	Explicit     bool               `json:"explicit"`
	ExternalIDs  ExternalIDs        `json:"external_ids"`
	ExternalURLs ExternalURLs       `json:"external_urls"`
	Popularity   int                `json:"popularity"`
	PreviewURL   *string            `json:"preview_url"` // nullable
	TrackNumber  int                `json:"track_number"`
	URI          string             `json:"uri"`
	IsLocal      bool               `json:"is_local"`
}

// SimplifiedArtist represents basic artist information.
type SimplifiedArtist struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	URI          string       `json:"uri"`
	ExternalURLs ExternalURLs `json:"external_urls"`
}

// SimplifiedAlbum represents basic album information.
type SimplifiedAlbum struct {
	ID                   string             `json:"id"`
	Name                 string             `json:"name"`
	AlbumType            string             `json:"album_type"` // album, single, compilation
	Artists              []SimplifiedArtist `json:"artists"`
	ReleaseDate          string             `json:"release_date"`
	ReleaseDatePrecision string             `json:"release_date_precision"` // year, month, day
	TotalTracks          int                `json:"total_tracks"`
	Images               []Image            `json:"images"`
	ExternalURLs         ExternalURLs       `json:"external_urls"`
	URI                  string             `json:"uri"`
}

// ExternalIDs represents external IDs like ISRC, EAN, UPC.
type ExternalIDs struct {
	ISRC string `json:"isrc,omitempty"`
	EAN  string `json:"ean,omitempty"`
	UPC  string `json:"upc,omitempty"`
}

// ExternalURLs represents external URLs (Spotify link).
type ExternalURLs struct {
	Spotify string `json:"spotify"`
}

// Image represents album/playlist cover art.
type Image struct {
	URL    string `json:"url"`
	Height *int   `json:"height"` // nullable
	Width  *int   `json:"width"`  // nullable
}

// ============================================================
// Playlist Models
// ============================================================

// Playlist represents a Spotify playlist with full metadata.
type Playlist struct {
	ID            string         `json:"id"`
	Name          string         `json:"name"`
	Description   *string        `json:"description"` // nullable
	Collaborative bool           `json:"collaborative"`
	Public        *bool          `json:"public"` // nullable
	Owner         User           `json:"owner"`
	Tracks        PlaylistTracks `json:"tracks"`
	Images        []Image        `json:"images"`
	ExternalURLs  ExternalURLs   `json:"external_urls"`
	SnapshotID    string         `json:"snapshot_id"`
	URI           string         `json:"uri"`
}

// PlaylistTracks represents paginated playlist tracks.
type PlaylistTracks struct {
	Href     string         `json:"href"`
	Limit    int            `json:"limit"`
	Next     *string        `json:"next"`     // nullable
	Offset   int            `json:"offset"`
	Previous *string        `json:"previous"` // nullable
	Total    int            `json:"total"`
	Items    []PlaylistItem `json:"items"`
}

// PlaylistItem represents a track in a playlist with metadata about when it was added.
type PlaylistItem struct {
	AddedAt string `json:"added_at"` // ISO 8601 datetime
	AddedBy User   `json:"added_by"`
	IsLocal bool   `json:"is_local"`
	Track   *Track `json:"track"` // nullable (for local files or episodes)
}

// User represents a Spotify user.
type User struct {
	ID           string       `json:"id"`
	DisplayName  *string      `json:"display_name"` // nullable
	ExternalURLs ExternalURLs `json:"external_urls"`
	URI          string       `json:"uri"`
}

// ============================================================
// Search Models
// ============================================================

// SearchResponse represents response from search endpoint.
type SearchResponse struct {
	Tracks *SearchTracksResult `json:"tracks,omitempty"`
}

// SearchTracksResult represents paginated search results for tracks.
type SearchTracksResult struct {
	Href     string  `json:"href"`
	Limit    int     `json:"limit"`
	Next     *string `json:"next"`     // nullable
	Offset   int     `json:"offset"`
	Previous *string `json:"previous"` // nullable
	Total    int     `json:"total"`
	Items    []Track `json:"items"`
}

// ============================================================
// Error Models
// ============================================================

// ErrorResponse represents Spotify API error response.
type ErrorResponse struct {
	Error ErrorDetails `json:"error"`
}

// ErrorDetails contains error details.
type ErrorDetails struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

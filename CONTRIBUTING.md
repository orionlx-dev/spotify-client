# Contributing to Spotify Client

Thank you for your interest in contributing! This document provides guidelines and instructions for contributing to this project.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Making Changes](#making-changes)
- [Running Tests](#running-tests)
- [Submitting Changes](#submitting-changes)
- [Style Guidelines](#style-guidelines)

## Code of Conduct

Be respectful, inclusive, and professional in all interactions.

## Getting Started

1. Fork the repository
2. Clone your fork:
   ```bash
   git clone https://github.com/your-username/spotify-client.git
   cd spotify-client
   ```
3. Add upstream remote:
   ```bash
   git remote add upstream https://github.com/your-username/spotify-client.git
   ```

## Development Setup

### Prerequisites

- Go 1.21 or higher
- Spotify Developer Account with API credentials
- Git

### Setup

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Create `.env` file:
   ```bash
   cp .env.example .env
   # Edit .env and add your Spotify credentials
   ```

3. Run examples to verify setup:
   ```bash
   go run examples/basic/main.go
   ```

## Making Changes

1. Create a feature branch:
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. Make your changes following the [Style Guidelines](#style-guidelines)

3. Test your changes:
   ```bash
   go build ./...
   go run examples/basic/main.go
   ```

4. Commit with clear messages:
   ```bash
   git commit -m "Add feature: description of what you added"
   ```

## Running Tests

Currently, the project uses manual testing via examples. Automated tests are planned.

Run all examples:
```bash
go run examples/basic/main.go
go run examples/search/main.go
go run examples/url-parsing/main.go
```

Run CLI tool:
```bash
go run cmd/spotify/main.go
```

## Submitting Changes

1. Push to your fork:
   ```bash
   git push origin feature/your-feature-name
   ```

2. Create a Pull Request on GitHub

3. In your PR description:
   - Describe what you changed and why
   - Reference any related issues
   - Include screenshots/examples if applicable
   - Confirm tests pass

## Style Guidelines

### Go Code Style

Follow standard Go conventions:

- Use `gofmt` to format code
- Use `golint` and `go vet` for code quality
- Write idiomatic Go code
- Keep functions small and focused
- Use meaningful variable names

### Documentation

- Add godoc comments for all exported types, functions, and methods
- Include examples in comments where helpful
- Update README.md if adding new features
- Update CHANGELOG.md

### Example godoc comment:

```go
// GetTrack retrieves track information by Spotify track ID.
//
// The trackID parameter must be a valid Spotify track identifier.
//
// Example:
//
//	track, err := client.GetTrack("11dFghVXANMlKmJXsNCbNl")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Track: %s\n", track.Name)
func (c *Client) GetTrack(trackID string) (*Track, error) {
	// implementation
}
```

### Commit Messages

Use clear, descriptive commit messages:

```
Add feature: support for album endpoints

- Implement GetAlbum method
- Add Album type to models
- Update documentation with examples
```

### Project Structure

```
spotify-client/
├── client.go           # Main client implementation
├── models.go           # Data models
├── config.go           # Configuration
├── utils.go            # Utility functions
├── examples/           # Example code
│   ├── basic/
│   ├── search/
│   └── url-parsing/
├── cmd/                # CLI tools
│   └── spotify/
├── .env.example        # Example environment file
├── README.md           # Main documentation
├── CHANGELOG.md        # Version history
├── CONTRIBUTING.md     # This file
└── LICENSE             # MIT License
```

## Adding New Features

When adding new Spotify API endpoints:

1. Add the response models to `models.go`
2. Add the client method to `client.go`
3. Add godoc comments with examples
4. Create an example in `examples/`
5. Update README.md with usage
6. Update CHANGELOG.md

## Questions?

- Open an issue for bugs or feature requests
- Start a discussion for general questions

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

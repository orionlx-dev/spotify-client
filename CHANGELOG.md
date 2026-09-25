# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Planned
- Add support for Albums API
- Add support for Artists API
- Add support for User's library operations
- Implement retry logic with exponential backoff
- Add context support for cancellation
- Add comprehensive test coverage

## [0.1.0] - 2026-09-25

### Added
- Initial release of Spotify Client library
- OAuth2 Client Credentials Flow authentication
- Automatic token management and refresh
- Get track by ID endpoint
- Search tracks with advanced filters
- Get playlist by ID endpoint
- Get all playlist tracks with automatic pagination
- Thread-safe client implementation
- Comprehensive documentation and examples
- Basic example demonstrating all features
- Search example with advanced query patterns
- CLI tool for testing API functionality

### Features
- Clean, idiomatic Go API
- Zero external dependencies (except godotenv for examples)
- Automatic token caching with 5-minute expiry buffer
- Support for custom HTTP clients
- Detailed error messages from Spotify API
- Well-documented public types and methods

[Unreleased]: https://github.com/your-username/spotify-client/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/your-username/spotify-client/releases/tag/v0.1.0

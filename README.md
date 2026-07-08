# Project Name

[![Build Status](https://img.shields.io/badge/build-passing-brightgreen)]
[![Coverage](https://img.shields.io/badge/coverage-70%25-salmon)]
[![License](https://img.shields.io/badge/license-MIT-yellow)]
[![go version](https://img.shields.io/badge/go-v1.26.4-blue)]
[![Downloads](https://img.shields.io/badge/downloads-0k%2Fmonth-purple)]

**One-line description:** Create and share secrets easily using this API.

[Screenshot or GIF demo here](have to find one)

## Why This Project?

- ✅ **Milestne 1:** liveProject goal is to build a secrets sharing application for a company that frequently shares sensitive data.

## Quick Start

```bash
# Clone repo to local machine
git clone https://github.com/username/repo.git

# cd into repository root directory
cd repo

# Download dependencies
go mod download

# Run
go run ./cmd/api

```

That's it! You're ready to use [Project Name].

## Installation

### Prerequisites

- Golang `1.26.4+`
- git

### Option 1: Build from source

```bash
git clone https://github.com/username/repo.git # Clone repo
cd repo # cd into repository root directory
go mod download # Download and install dependencies
go build -o bin/secret-store ./cmd/api # Build the application
```

## Usage

### Basic Example

```bash
curl -d '{"secret": "super-secret-word"}' http://localhost:8080/

curl -X GET http://localhost:8080/some-hash
```

## Configuration

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `DATA_FILE_PATH` | string | None | Path for file-based storage |

### Environment Variables

```bash
DATA_FILE_PATH=/path/to/json/file.json
```

## API Reference

| Endpoint | HTTP method | Payload expected |
| -- | -- | -- |
| `/` | `POST` | Yes |
| `/{id}` | `GET` | No |
| `/healthcheck` | `GET` | No |

## API Usage

**1.** The following send a create secret request and receive a response containing a `SHA256` hash of the secret.

```bash
PAYLOAD='{"secret": "super-secret"}'
URL='http://localhost:8080'

curl -d "$PAYLOAD" "$URL"/
# output: {"id": "vnerbnvebnernvewcinij34323wq"}
```

**2.** Then use that generated hash to retrieve the secret.

```bash
ID='vnerbnvebnernvewcinij34323wq'
URL='http://localhost:8080'

curl -X GET "$URL"/"$ID"
# output: {"secret": "super-secret"}
```

## Performance

> todo
<!-- | Metric | This Project | Alternative A | Alternative B |
|--------|--------------|---------------|---------------|
| Requests/sec | 47,000 | 23,000 | 31,000 |
| Latency (p99) | 12ms | 45ms | 28ms |
| Memory usage | 45MB | 120MB | 78MB |
| Cold start | 150ms | 890ms | 420ms | -->

<!-- *Benchmarks run on AWS c5.xlarge, Node.js 20, Ubuntu 22.04* -->

## Roadmap

- [x] Milestone 1
- [ ] Milestone 2
- [ ] Milestone 3

See the [project](https://www.manning.com/liveproject/build-a-secrets-sharing-web-application) for more information.

## Troubleshooting

> todo
<!-- 
### Error: "Cannot find module"

```bash
# Clear node_modules and reinstall
rm -rf node_modules package-lock.json
npm install
```

### Error: "Port already in use"

```bash
# Find and kill process on port
lsof -i :3000
kill -9 [PID]
```

### Still stuck?

- Check [existing issues](link)
- Join our [Discord](link)
- Open a [new issue](link) -->

## License

MIT © [Travis](https://github.com/travboz)

## Acknowledgments

- [Manning](https://www.manning.com/) - Producer of live project

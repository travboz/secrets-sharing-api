# Secrets Sharing Web API

![License](https://img.shields.io/badge/license-MIT-yellow)
![go version](https://img.shields.io/badge/go-v1.26.4-blue)
<!-- ![Build Status](https://img.shields.io/badge/build-passing-brightgreen)
![Coverage](https://img.shields.io/badge/coverage-70%25-salmon)
![Downloads](https://img.shields.io/badge/downloads-0k%2Fmonth-purple) -->

> **One-liner:** Easily **create** and **share** secrets with colleagues using this small API.

<figure>
    <img src="knight-gopher.png"
         alt="Hardened Knight Gopher"
         width="350px">
    <figcaption>Our Knight Gopher standing ready to defend against any attempts at brute force attacks!.</figcaption>
</figure>

## Roadmap & Why This Project?

I purchased this project course because I was under the impression *we'd* build something cool - turns out **I** was the one doing the building. This project helped build confidence in testing and problem solving (specifically, working through the steps required to solve a problem).

See the [project](https://www.manning.com/liveproject/build-a-secrets-sharing-web-application) brief on Manning for more information.

- [x] **Milestone 1, Create the Secret Sharing API:** Create a web application which will allow you to create and view secrets.
- [x] **Milestone 2, Testing the API:** Use the Go standard libraries to write tests to verify the functionality of the secret sharing web application.
- [x] **Milestone 3, Encrypt the data at rest:** Implement the encryption of the secrets at rest. Use a symmetric cryptographic algorithm (Advanced Encryption Standard - AES) to encrypt the secret when storing it and then decrypt it back when reading from the file.

### Note: My learning experience

> I feel like the structure of this project may have been overengineered. The saving grace here is that I believe I can easily extend it - for example by swapping out the file-based storage with S3 or MongoDB. It is also easier tested as the interfaces allow for simple mocking. I have also attempted to implement some methods learned in Learn Go with Tests (for example, spying on calls).

## Quick Start

```bash
# Clone repo to local machine
git clone https://github.com/travboz/secrets-sharing-api.git

# cd into repository root directory
cd secrets-sharing-api

# Download dependencies
go mod download

# Change .env.example to .env
mv .env.example .env

# Run
go run ./cmd/api

# Navigate to the URL, for e.g.
curl http://localhost:8080/healthcheck
```

That's it! You're ready to use it.

## Installation

### Prerequisites

- Golang `1.26.4+`
- git

### Picking a milestone

This project consisted of the `3` milestone - which were created in sequence.

To jump into a particular milestone just pick one, navigate to its directory, and explore.

For example:

```bash
# Clone repo
git clone https://github.com/travboz/secrets-sharing-api.git

# Navigate to the final milestone's directory
cd secrets-sharing-api/milestone3-code
```

### Option 1: Build from source

```bash
# Clone repo
git clone https://github.com/travboz/secrets-sharing-api.git

 # Navigate to repository root directory
cd secrets-sharing-api/milestone3-code

# Download and install dependencies
go mod download

# Change .env.example to .env
mv .env.example .env

# Build the application
go build -o bin/secret-store ./cmd/api
```

### Option 2: Build using `Task`

```bash
# Clone the repo
git clone https://github.com/travboz/secrets-sharing-api.git

# Navigate to repository root directory
cd secrets-sharing-api/milestone3-code

# Download and install dependencies
go mod download

# To view available tasks
task -l

# Build the binary and run it
task build && task run-binary
```

## Usage

### Directory structure

To illustrate, here is the tree from the `milestone3-code` directory + those in the `root` of the repository.

The other milestone directories follow a similar structure.

```bash
.
├── README.md # Info and description on the repo.
├── knight-gopher.png # Our courageous defender!
└── milestone3-code # All milestone 3 related code.
    ├── Taskfile.yml # Contains the tasks runnable for this project.
    ├── api-test.http # Used for repeatable manual endpoint testing (see VSCode extension 'REST Client' by Huachao Mao).
    ├── bin # Omitted from repo but this is where the 'build' task places the API's binary.
    │   └── secret-share
    ├── cmd
    │   └── api
    │       ├── create_secret_handler_test.go # Tests related to the POST endpoint used in creating new secrets.
    │       ├── get_secret_handler_test.go # Tests related to the GET endpoint for fetching created secrets.
    │       ├── handlers.go # Contains the health check, get and post handlers for managing secrets.
    │       ├── healthcheck_handler_test.go # Tests pertaining to the GET health check endpoint.
    │       ├── helpers.go # Utils used through the API. e.g. writing json to the response, reading in json from a request, hashing secrets, etc.
    │       ├── main.go # Entrypoint of the API
    │       ├── routes.go # Attach routes to the mux.
    │       ├── routes_test.go # Tests related to testing that our router accepts or rejects the correct URLs (e.g. we do not want PUT requests to our create secrets endpoint).
    │       ├── testutils.go # Contains mocks used for testing. Currently not using Table-Driven tests as goal is to refactor once project (parts 1 and 2) are complete.
    │       └── types.go # Request and response types used in handlers.
    ├── data.json # File created to store secrets in file system - defined by the 'DATA_FILE_NAME` environment variable.
    ├── go.mod
    ├── go.sum
    ├── internal
    │   ├── encryption # Files used for encrypting the secrets at rest.
    │   │   ├── cryptoconfig # Concrete implementation using AES-GCM of Encrypter interface.
    │   │   │   ├── cryptoconfig.go
    │   │   │   └── types.go
    │   │   └── interface.go # Encrypter interface in the event I wanted to use a different algorithm (or mock during testing). 
    │   └── store # Data access for the secrets data storage.
    │       ├── filestore # File-based store.
    │       │   └── filestore.go
    │       ├── interface.go # Store interface used to support easier testing and future storage implementations. 
    │       └── types.go # Storage-associated types (particularly SecretData - used as the base type for the secret storage data).
    └── pkg
        └── testing
            └── assert # Assertion library discovered through following Alex Edwards' work on Let's Go and Let's Go Further - useful and simple assertion package.
                └── assert.go
```

### `.env` file configuration

| Option | Type | Default | Description |
| -------- | ------ | --------- | ------------- |
| `DATA_FILE_PATH` | string | None | Path for file-based storage |
| `PASSWORD` | string | None | A unique password you choose to encrypt and decrypt the secrets. |
| `SALT` | string | None | The salt is some random data used in conjunction with the password for encryption and decryption. |

You'll find a `.env` file included in the repo which contains some default values for these values (as they are **all** ***required***) as these are all used within the store for the application.

Ensure you either **create** your own `.env` file, or rename the sample env file: `mv .env.example .env`, before running the application otherwise you'll encounter an error.

### Basic Example

```bash
curl -X POST -d '{"plain_text":"super-secret-word"}' http://localhost:8080/

ID=vnerbnvebnernvewcinij34323wq
curl -X GET http://localhost:8080/$ID
```

### Environment Variables

```bash
DATA_FILE_PATH=/path/to/json/file.json
PASSWORD=thecolourblue
SALT=pinkseasalt
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
PAYLOAD='{"plain_text":"super-secret"}'
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

> ***NOTE:*** There is **a lot of repetition** in the `cmd/api/XXX_test.go` files. This is intended as the focus was to complete the project. Refactoring will occur later - when application is completed.
<!-- | Metric | This Project | Alternative A | Alternative B |
|--------|--------------|---------------|---------------|
| Requests/sec | 47,000 | 23,000 | 31,000 |
| Latency (p99) | 12ms | 45ms | 28ms |
| Memory usage | 45MB | 120MB | 78MB |
| Cold start | 150ms | 890ms | 420ms | -->

<!-- *Benchmarks run on AWS c5.xlarge, Node.js 20, Ubuntu 22.04* -->

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

- [Manning](https://www.manning.com/) - Producer of liveProject
- [Gopher](https://github.com/egonelbre/gophers) illustration by Egon Elbre (egonelbre/gophers), [CC0 1.0](https://creativecommons.org/publicdomain/zero/1.0/)

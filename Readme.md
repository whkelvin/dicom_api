# Limitations and Areas of Improvement

- Add in proper logging middleware
- Authentication is not implemented
- Currently dicom file and generated png are stored in `assets/uploaded`, relationships between dicom and pngs are not stored anywhere (which pngs are generated from which dicom file).
- Storage Path are hardcoded in the code at the moment. Config should be set up properly to handle config values/env variables.
- Tests are incomplete/incomprehensive, `cmd/test/controller_test.go` only shows that mock testing is easy because layers are loosely coupled and dependent on interfaces.

# Running the app

## Run Server

`go run cmd/server/main.go`

## Run Tests

`go test -v ./cmd/test/`

# API Interface

Open API Specification is included in `/api`. You can see a webview of the spec by pasting the content of the spec [here](https://editor.swagger.io/)

# Architecture

The project structure is based on a mix of Clean Arch/Onion Arch and Vertical Slices Arch. I've combined my favorite features of the two architectures where CA focuses on making horizontal layers loosely coupled and VSA focuses on making features isolated from one another and preventing merge conflicts when multiple devs are working on the same project.

# Limitations and Areas of Improvement

- Add in proper logging middleware
- Authentication is not implemented
- Currently dicom file and generated png are stored in `assets/uploaded`, relationships between dicom and pngs are not stored anywhere (which pngs are generated from which dicom file).
- Tests are incomplete/incomprehensive, `cmd/test/controller_test.go` only shows that mock testing is easy because layers are loosely coupled and dependant on interfaces.

# Running the app

## Run Server

`go run cmd/server/main.go`

## Run Tests

`go test -v ./cmd/test/`

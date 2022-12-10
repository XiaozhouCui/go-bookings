## Initialisation

- Copy all 3 folders from go-web-app
- Run `go mod init bookings`
- In all imports, replace `gowebapp` with `bookings`
- Run `go run cmd/web/*.go`, it will tell you the missing imports
- Get all the missing packages, e.g. `go get github.com/alexedwards/scs/v2`

## Unit tests

- Add _setup_test.go_ to do the initial setup
- Make sure `gcc` is installed in Linux
- Add _blah_test.go_ files inside _./cmd/web/_
- Go to _./cmd/web_, run `go test`
- To show tests verbosely, run `go test -v`
- To show coverage, run `go test -cover`
- To create html report, run `go test -coverprofile=coverage.out && go tool cover -html=coverage.out`

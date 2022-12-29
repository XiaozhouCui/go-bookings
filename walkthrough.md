## Initialisation

- Copy all 3 folders from go-web-app
- Run `go mod init bookings`
- In all imports, replace `gowebapp` with `bookings`
- Run `go run cmd/web/*.go`, it will tell you the missing imports
- Get all the missing packages, e.g. `go get github.com/alexedwards/scs/v2`

## Unit tests

- Go to _./cmd/web_, add `setup_test.go` to do the initial setup
- Make sure `gcc` is installed in Linux
- Add `*_test.go` files inside _./cmd/web/_
- In _./cmd/web_, run `go test`
- To show tests verbosely, run `go test -v`
- To show coverage, run `go test -cover`
- To create html report, run `go test -coverprofile=coverage.out && go tool cover -html=coverage.out`
- To run all tests, go to project root folder, run `go test -v ./...`

## Add run script

- create _run.sh_
- make the script executable: `chmod +x run.sh`
- run `./run.sh`

## Install Soda for DB migration

- Install pop: `go get github.com/gobuffalo/pop/...`
- Install soda CLI: `go install github.com/gobuffalo/pop/v6/soda@latest`
- Update _~/.profile_ (or _~/.zprofile_ on mac), add new line `export PATH="$HOME/go/bin:$PATH"`
- Restart terminal

## Initialisation

- Copy all 3 folders from go-web-app
- Run `go mod init bookings`
- In all imports, replace `gowebapp` with `bookings`
- Run `go run cmd/web/*.go`, it will tell you the missing imports
- Get all the missing packages, e.g. `go get github.com/alexedwards/scs/v2`

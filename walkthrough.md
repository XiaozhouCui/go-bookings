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

## Install Pop and Soda for DB migration

- View the doc here: https://gobuffalo.io/documentation/database/pop/
- Install pop: `go get github.com/gobuffalo/pop/...`
- Install soda CLI: `go install github.com/gobuffalo/pop/v6/soda@latest`
- Update _~/.profile_ (or _~/.zprofile_ on mac), add new line `export PATH="$HOME/go/bin:$PATH"`
- Restart terminal, run `soda -v` to check installation

## Setup database in postgres

- Add docker-compose file and run `docker compose up`
- Connect to db using DBeaver at `localhost:54321`

## Migration

- Create _database.yml_
- Run `soda generate fizz CreateUserTable`
- Update the generated file `*_create_user_table.up.fizz`
- Run `soda migrate` to create the users table
- Update the generated file `*_create_user_table.down.fizz`
- Run `soda migrate down` to drop the users table

## Create other tables

- `soda generate fizz CreateReservationTable`
- `soda generate fizz CreateRoomsTable`
- `soda generate fizz CreateRestrictionsTable`
- `soda generate fizz CreateRoomRestrictionsTable`

## Add foreign key

- Run `soda generate fizz CreateFKForReservationsTable`
- Update the generated migration files
- Run `soda migrate`
- Repeat with `soda generate fizz CreateFKForRoomRestrictions`

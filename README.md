
<h1 align="center">
  <br>
  Payment API
  <br>
</h1>

<h4 align="center">A basic payment backend system in
<a href="https://go.dev" target="_blank">Go</a>, uses in-memory store, mutex and concurrency.</h4>

<p align="center">
  <a href="#key-features">Key Features</a> •
  <a href="#how-to-use">How To Use</a> •
  <a href="#how-to-run-tests">How To run tests</a> •
  <a href="#download">Download</a>
</p>


## Key Features

* Utilizes Go Maps as in-memory stores for accounts & transfers.
* Initializes storage with singleton pattern to avoid any data loss.
* Uses [Gin](https://gin-gonic.com/) web framework to run a web server.
* Multi-layered architecture (Controller, Service, Model, Storage) for easy expansion.
* Starts with an initial dataset of accounts [link](https://git.io/Jm76h), which is processes concurrently into the datastore.
* Utilizes Go Mutex to grantee atomic transfers between accounts.
* Covered with tests to verify the flow.
* Runs on the CLI without any installations effort.


## How To Use

To clone and run this application, you'll need [Git](https://git-scm.com) and [Go](https://go.dev/dl/). From your command line:

```bash
# Clone this repository
$ git clone https://github.com/mohamedsaberibrahim/basic-payment-system

# Go into the repository
$ cd basic-payment-system

# Install dependencies
$ go build

# Run the app
$ go run .
```

## How to run tests
If you have [make]() installed, you can run `make test`. If not, you can you `go test -v ./...`

## Download

You can [download](https://github.com/mohamedsaberibrahim/basic-payment-system/blob/main/bin/payment-api) the latest installable version of Payment API for Windows, macOS and Linux.

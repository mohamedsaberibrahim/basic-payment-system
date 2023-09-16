// Package server is the application's entry point.
// This package is used to start the application.
package server

const (
	// DefaultPort is the default port to listen on.
	DefaultPort = ":8080"
)

func Init() {
	// Initialize the application.
	r := NewRouter()
	r.Run(DefaultPort)
}

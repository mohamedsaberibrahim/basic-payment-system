// Package server is the application's entry point.
// This package is used to start the application.
package server

import "github.com/gin-gonic/gin"

const (
	// DefaultPort is the default port to listen on.
	DefaultPort = ":8080"
)

func Init() {
	// Set the Gin mode to release.
	gin.SetMode(gin.ReleaseMode)

	// Initialize the application.
	r := NewRouter()
	r.Run(DefaultPort)
}

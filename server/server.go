// Package server is the application's entry point.
// This package is used to start the application.
package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	// DefaultPort is the default port to listen on.
	DefaultPort = ":8080"
)

func Init() {
	// Set the Gin mode to release.
	gin.SetMode(gin.ReleaseMode)

	fmt.Println("Starting the application...")
	fmt.Println("Server is running on port", DefaultPort)

	// Initialize the application.
	r := NewRouter()
	r.Run(DefaultPort)
}

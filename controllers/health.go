// Package controllers is a request handler package, through using the services package and return a response.
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (h HealthController) Status(c *gin.Context) {
	c.String(http.StatusOK, "Server is up and running!")
}

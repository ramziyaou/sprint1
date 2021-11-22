package controllers

import (
	"fmt"
	"net/http"
)

// Protected endpoint for token verification
func (c *Controller) ProtectedEndpoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("protectedEndpoint invoked.")
	}
}

// main_test.go
package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func TestHelloWorldHandler(t *testing.T) {
    // Set Gin to Test Mode
    gin.SetMode(gin.TestMode)

    // Setup Routes
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.String(200, "Hello, World!")
    })

    // Create a test context 
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/", nil)
    r.ServeHTTP(w, req)

    // Assert we got the expected response
    assert.Equal(t, 200, w.Code)
    assert.Equal(t, "Hello, World!", w.Body.String())
}

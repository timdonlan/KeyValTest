package router

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
	handler := func(c *gin.Context) {
		c.String(http.StatusOK, "bar")
	}

	router := gin.New()
	router.GET("/foo", handler)

	req, _ := http.NewRequest("GET", "/foo", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	fmt.Print(resp.Body.String())

	assert.Equal(t, resp.Body.String(), "bar")
}



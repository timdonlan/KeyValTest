package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
	"KeyValTest/model"
	"fmt"
)

func TestGetKeyVal(t *testing.T) {

	//mock db
	oldModelGetKeyVal := modelGetKeyVal
	defer func() { modelGetKeyVal = oldModelGetKeyVal }()
	modelGetKeyVal = func(key string) (*model.KeyValData, error) {
		return &model.KeyValData{key, "bar"}, nil
	}

	r := gin.Default()
	r.GET("/key/:key", getKeyVal)
	req, err := http.NewRequest("GET", "/key/foo", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Error("HTTP status expected: 200, got: %d", w.Code)
	}
	fmt.Print(w.Body.String())

}

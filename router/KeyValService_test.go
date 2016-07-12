package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
	"KeyValTest/model"
	"fmt"
)

type mockedKeyValDAL struct{}

func (m *mockedKeyValDAL) GetAll() ([]*model.KeyValData, error){
	return nil,nil
}
func (m *mockedKeyValDAL) GetKeyVal(key string) (*model.KeyValData, error){
	return nil,nil
}
func (m *mockedKeyValDAL) CreateKeyVal(key string, value string) (*model.KeyValData, error){
	return nil,nil
}
func (m *mockedKeyValDAL) UpdateKeyVal(key string, newValue string) (*model.KeyValData, error){
	return nil,nil
}
func (m *mockedKeyValDAL) DeleteKeyVal(key string) (bool, error){
	return false,nil
}

func TestGetAll(t *testing.T) {
	//Mock DAL for testing service
	keyValDAL = new(mockedKeyValDAL)

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

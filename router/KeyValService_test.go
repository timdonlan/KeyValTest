package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
	"KeyValTest/model"
	_"bytes"
	"fmt"
	"bytes"
	"encoding/json"
)

/**
* Mocked Data Access Layer
 */

type mockedKeyValDAL struct{}

func (m *mockedKeyValDAL) GetAll() ([]*model.KeyValData, error){

	var keyValSlice []*model.KeyValData
	keyValSlice = append(keyValSlice,&model.KeyValData{"hello","world"})

	return keyValSlice,nil
}
func (m *mockedKeyValDAL) GetKeyVal(key string) (*model.KeyValData, error){
	return &model.KeyValData{"hello","world"},nil
}
func (m *mockedKeyValDAL) CreateKeyVal(key string, value string) (*model.KeyValData, error){
	return &model.KeyValData{"hello","world"},nil
}
func (m *mockedKeyValDAL) UpdateKeyVal(key string, newValue string) (*model.KeyValData, error){
	return &model.KeyValData{"hello","world"},nil
}
func (m *mockedKeyValDAL) DeleteKeyVal(key string) (bool, error){
	return true,nil
}

//-----------------------------

func TestCreate(t *testing.T) {
	KeyValDAL = new(mockedKeyValDAL)

	r := gin.Default()
	r.POST("/key", createKeyVal)

	jsonData := &model.KeyValData{"hello","world"}
	jsonStr,_ := json.Marshal(jsonData)

	req, err := http.NewRequest("POST", "/key", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	fmt.Print(w.Body.String())
	if w.Code != 200 {
		t.Error("HTTP status expected: 200, got: %d", w.Code)
	}
}

func TestUpdate(t *testing.T) {
	KeyValDAL = new(mockedKeyValDAL)

	r := gin.Default()
	r.PUT("/key/:key", updateKeyVal)

	jsonData := &model.KeyValData{"hello","world"}
	jsonStr,_ := json.Marshal(jsonData)

	req, err := http.NewRequest("PUT", "/key/hello", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	fmt.Print(w.Body.String())
	if w.Code != 200 {
		t.Error("HTTP status expected: 200, got: %d", w.Code)
	}
}

func TestGetKeyVal(t *testing.T) {
	//Mock DAL for testing service
	KeyValDAL = new(mockedKeyValDAL)

	r := gin.Default()
	r.GET("/key/:key", getKeyVal)
	req, err := http.NewRequest("GET", "/key/hello", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Error("HTTP status expected: 200, got: %d", w.Code)
	}
}

func TestGetAll(t *testing.T) {
	KeyValDAL = new(mockedKeyValDAL)
	r := gin.Default()
	r.GET("/keys", getKeyVal)
	req, err := http.NewRequest("GET", "/keys", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Error("HTTP status expected: 200, got: %d", w.Code)
	}
}

func TestDelete(t *testing.T) {
	KeyValDAL = new(mockedKeyValDAL)
	r := gin.Default()
	r.DELETE("/key/:key", deleteKeyVal)
	req, err := http.NewRequest("DELETE", "/key/hello", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Error("HTTP status expected: 200, got: %d", w.Code)
	}
}

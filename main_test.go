package main

import (
	"net/http"
	"testing"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"KeyValTest/router"
	"KeyValTest/model"
	"bytes"
	"encoding/json"
	"strings"
)

func testPOST(t *testing.T, url string, body []byte, respStatus int, rspBody []byte){
	resp, err := http.Post(url,"application/json", bytes.NewBuffer(body))
	defer resp.Body.Close()
	assert.NoError(t, err)
	body, ioerr := ioutil.ReadAll(resp.Body)
	assert.NoError(t, ioerr)
	assert.Equal(t, strings.TrimSpace(string(rspBody)),strings.TrimSpace(string(body)), "resp body should match")
	assert.Equal(t,respStatus, resp.StatusCode, "StatusCodes did not match")
}

func testGET(t *testing.T, url string, respStatus int, rspBody []byte){
	resp, err := http.Get(url)
	defer resp.Body.Close()
	assert.NoError(t, err)
	body, ioerr := ioutil.ReadAll(resp.Body)
	assert.NoError(t, ioerr)
	assert.Equal(t, strings.TrimSpace(string(rspBody)), strings.TrimSpace(string(body)), "resp body should match")
	assert.Equal(t,respStatus, resp.StatusCode, "StatusCodes did not match")
}

func testPUT(t *testing.T, url string, body []byte, respStatus int, rspBody []byte){
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp,err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	assert.NoError(t, err)
	body, ioerr := ioutil.ReadAll(resp.Body)
	assert.NoError(t, ioerr)
	assert.Equal(t, strings.TrimSpace(string(rspBody)), strings.TrimSpace(string(body)), "resp body should match")
	assert.Equal(t,respStatus, resp.StatusCode, "StatusCodes did not match")
}



func testDELETE(t *testing.T, url string, respStatus int, rspBody string){
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp,err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	assert.NoError(t, err)
	body, ioerr := ioutil.ReadAll(resp.Body)
	assert.NoError(t, ioerr)
	assert.Equal(t, strings.TrimSpace(string(rspBody)), strings.TrimSpace(string(body)), "resp body should match")
	assert.Equal(t,respStatus, resp.StatusCode, "StatusCodes did not match")
}


func TestHttpServer(t *testing.T) {

	dbName := "default.db"

	//Setup backend db
	model.ResetDB(dbName)
	defer model.CleanDB(dbName)
	router.KeyValDAL = model.OpenDB(dbName)

	//Start test server
	r := router.CreateServiceHandlers()
	ts := httptest.NewServer(r)
	defer ts.Close()

	jsonData := &model.KeyValData{"hello","world"}
	jsonStr,_ := json.Marshal(jsonData)
	testPOST(t,ts.URL+"/key",jsonStr,200,jsonStr)

	testGET(t,ts.URL+"/key/hello",200,jsonStr)

	jsonData = &model.KeyValData{"hello","world2"}
	jsonStr,_ = json.Marshal(jsonData)
	testPUT(t,ts.URL+"/key/hello",jsonStr,200,jsonStr)

	testGET(t,ts.URL+"/key/hello",200,jsonStr)

	testDELETE(t,ts.URL+"/key/hello",200,"true")

	testGET(t,ts.URL+"/key/hello",500, []byte("{\"error\":\"sql: no rows in result set\"}"))
}
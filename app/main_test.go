package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleRoot(t *testing.T) {
	ts := httptest.NewServer(routes())
	defer ts.Close()

	client := new(http.Client)

	req, _ := http.NewRequest("GET", ts.URL, nil)

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Hello, World", string(respBody))
}

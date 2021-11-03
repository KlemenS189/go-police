package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func doRequest(method, uri string, body *bytes.Buffer, handle httprouter.Handle) (*httptest.ResponseRecorder, error) {
	resp := httptest.NewRecorder()
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return nil, err
	}

	router := httprouter.New()
	router.Handle(method, uri, handle)
	router.ServeHTTP(resp, req)
	return resp, nil
}

func Test_canPassThrough(t *testing.T) {
	passThroughPercent = 1.0
	if canPassThrough() != true {
		t.Error("If passThroughPercent is 1, all requests should be logged")
	}
}

func TestViolationHandlerNoContent(t *testing.T) {
	passThroughPercent = 0
	endpoint = "/test/"
	router := initRouter()

	req := httptest.NewRequest("POST", endpoint, nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	status := rr.Code
	if status != http.StatusNoContent {
		t.Error("Should be 204")
	}
}

func TestViolationHandlerBadBody(t *testing.T) {
	passThroughPercent = 1.0
	endpoint = "/test/"

	body := bytes.NewBufferString("bad body")
	response, _ := doRequest("POST", endpoint, body, violationHandler)
	if response.Code != http.StatusBadRequest {
		t.Error("Should be bad status code")
	}
}

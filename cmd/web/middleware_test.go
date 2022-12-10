package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	// mocked handler type from setup_test.go
	var myH myHandler

	h := NoSurf(&myH)

	// make assertion on the return type of NoSurf, should return "http.Handler"
	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("type is not http.Handler, but is %T", v)
	}
}

func TestSessionLoad(t *testing.T) {
	// mocked handler type from setup_test.go
	var myH myHandler

	h := SessionLoad(&myH)

	// test if the return type of SessionLoad is "http.Handler"
	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("type is not http.Handler, but is %T", v)
	}
}

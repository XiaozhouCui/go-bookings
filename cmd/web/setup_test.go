package main

import (
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

type myHandler struct{}

// ServeHTTP method has a pointer receiver *myHandler
func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

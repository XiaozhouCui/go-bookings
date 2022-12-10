package main

import (
	"bookings/internal/config"
	"testing"

	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	// routes() should return a pointer type *chi.Mux
	mux := routes(&app)

	// make assertion on the return type
	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing, test passed
	default:
		t.Errorf("type is not *chi.Mux, type is %T", v)
	}
}

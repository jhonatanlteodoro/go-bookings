package main

import (
	"testing"

	"github.com/go-chi/chi"
	"github.com/jhonatanlteodoro/go-bookings/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing :)
	default:
		t.Errorf("Type is not *chi.Mux. Current type is %T", v)
	}
}

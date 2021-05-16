package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/joyohartonowork/bookings/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		//nothing
	default:
		t.Error(fmt.Sprintf("type not *chi.mux, %T", v))
	}

}

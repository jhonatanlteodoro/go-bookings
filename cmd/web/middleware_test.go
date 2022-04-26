package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH myHandler
	h := NoSurf(&myH)

	switch h.(type) {
	case http.Handler:
		// Great :)
	default:
		t.Error("type is not http.handler")
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler
	h := SessionLoad(&myH)

	switch h.(type) {
	case http.Handler:
		// Great :)
	default:
		t.Error("type is not http.handler")
	}
}

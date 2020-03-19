package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var Message string = "Code.education Rocks!"

func TestServer(t *testing.T) {
	response := httptest.NewRecorder()

	// codigo de retorno da pagina
	rsCode := response.Code

	if rsCode != http.StatusOK {
		t.Fatalf("Expected status code %v, but received: %v", "200", rsCode)
	}
}

func TestCrash(t *testing.T) {
	if strings.Contains(Crash(), Message) == false {
		t.Fatalf("Expected: %v, but received: %v", Message, Crash())
	}
}

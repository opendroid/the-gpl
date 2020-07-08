package wwwexamples

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test_rootHandler tests the server
//  Example: https://blog.questionable.services/article/testing-http-handlers-go/
//  go test -v -run Test_rootHandler
func Test_rootHandler(t *testing.T) {
	// Test root
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Declare a recorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rootHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("rootHandler: received status: %v expected: %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "Hello ASERV\n"
	if rr.Body.String() != expected {
		t.Errorf("rootHandler: received body: %v expected: %v",
			rr.Body.String(), expected)
	}
}

// Test_incrHandler
//  go test -v -run Test_rootHandler
func Test_incrHandler(t *testing.T) {
	// Test incr
	req, err := http.NewRequest("GET", "/incr", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Declare a recorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(incrHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("incrHandler: received status: %v expected: %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "URL: \"/incr\"\n"
	if rr.Body.String() != expected {
		t.Errorf("incrHandler: received body: %v expected: %v", rr.Body.String(), expected)
	}
}

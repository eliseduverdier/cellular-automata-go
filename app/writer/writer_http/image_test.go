package writer_http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRenderImagePage(t *testing.T) {
	req, _ := http.NewRequest("GET", "/text", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RenderImagePage)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect
	if actual := rr.Code; actual != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			actual, http.StatusOK)
	}

	// Check the response body is an image
	expected := "image/png"
	if actual := rr.Header().Get("Content-Type"); actual != expected {
		t.Errorf("handler returned unexpected body: expected %s, got %s",
			expected, actual)
	}
}

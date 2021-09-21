package writer_http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRenderImagePage(t *testing.T) {
	req, err := http.NewRequest("GET", "/text", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RenderImagePage)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "image/png"
	actual := rr.Header().Get("Content-Type")
	//Check the response body is an image
	if actual != expected {
		t.Errorf("handler returned unexpected body: expected %s, got %s",
			expected, actual)
	}
}

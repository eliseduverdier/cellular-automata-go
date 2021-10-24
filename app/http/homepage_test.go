package http

import (
	"testing"
)

func TestRenderHomePage(t *testing.T) {
	// req, _ := http.NewRequest("GET", "/", nil)
	// rr := httptest.NewRecorder()
	// handler := http.HandlerFunc(RenderHomePage)

	// handler.ServeHTTP(rr, req)

	// // ------------ Check status code
	// if actual := rr.Code; actual != http.StatusOK {
	// 	t.Errorf("handler returned wrong status code: got %v want %v",
	// 		actual, http.StatusOK)
	// }

	// // ------------- compare content
	// expected := "......."
	// if rr.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: expected <%s>, got <%s>",
	// 		expected, rr.Body.String())
	// }
}

package handlercomp

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNetHTTP(t *testing.T) {
	t.Parallel()
	s := NewHTTPServer(":8080")

	req := httptest.NewRequest(http.MethodGet, "/user/1234?name=MyName", nil)
	rec := httptest.NewRecorder()

	s.Handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	expectedBody := "ID: 1234, Name: MyName"
	if rec.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rec.Body.String())
	}
}

func BenchmarkNetHTTP(b *testing.B) {
	s := NewHTTPServer(":8080")
	req := httptest.NewRequest(http.MethodGet, "/user/1234?name=MyName", nil)
	rec := httptest.NewRecorder()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Handler.ServeHTTP(rec, req)
	}
}

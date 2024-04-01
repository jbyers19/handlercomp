package handlercomp

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEcho(t *testing.T) {
	t.Parallel()
	e := NewEchoServer()

	req := httptest.NewRequest(http.MethodGet, "/user/1234?name=MyName", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1234")

	err := echoHandler(c)
	if err != nil {
		t.Errorf("handler returned error: %q", err.Error())
	}

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	expectedBody := "ID: 1234, Name: MyName"
	if rec.Body.String() != expectedBody {
		t.Errorf("expected body %q but got %q", expectedBody, rec.Body.String())
	}
}

func BenchmarkEcho(b *testing.B) {
	e := NewEchoServer()
	req := httptest.NewRequest(http.MethodGet, "/user/1234?name=MyName", nil)
	rec := httptest.NewRecorder()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1234")
		_ = echoHandler(c)
	}
}

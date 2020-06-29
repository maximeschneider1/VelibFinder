package handler

import (
  "github.com/julienschmidt/httprouter"
  "net/http"
  "net/http/httptest"
  "testing"
)

// TestHealthCheckHandler send get request to the server to check if it's still up and running
func TestHealthCheckHandler(t *testing.T) {
  req, err := http.NewRequest("GET", "/healthcheck", nil)
  if err != nil {
    t.Fatal(err)
  }
  s := server{
    router : httprouter.New(),
  }

  rr := httptest.NewRecorder()
  handler := s.HealthCheckHandler()

  s.router.GET("/healthcheck", handler)

  s.router.ServeHTTP(rr, req)

  // Check if the status code and response body are the one expected.
  if status := rr.Code; status != http.StatusOK {
    t.Errorf("handler returned wrong status code: got %v want %v",
      status, http.StatusOK)
  }
  expected := `{"alive": true}`
  if rr.Body.String() != expected {
    t.Errorf("handler returned unexpected body: got %v want %v",
      rr.Body.String(), expected)
  }
}

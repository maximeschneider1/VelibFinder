package handler

import (
  "github.com/julienschmidt/httprouter"
  "io"
  "net/http"
)

// HealthCheckHandler can be useful in production
func (s *server)  HealthCheckHandler() httprouter.Handle {
  return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

    w.WriteHeader(http.StatusOK)

    io.WriteString(w, `{"alive": true}`)
  }
}

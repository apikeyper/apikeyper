package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) HealthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.Db.Health())
	_, _ = w.Write(jsonResp)
}

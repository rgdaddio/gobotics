package server

import (
	"encoding/json"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, req *http.Request) {
	msg := ServerMsg{Message: "live"}
	json.NewEncoder(w).Encode(msg)
}

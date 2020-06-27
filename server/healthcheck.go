package server

import (
	"encoding/json"
	"net/http"
)

func healthcheck(w http.ResponseWriter, req *http.Request) {
	msg := Msg{Message: "live"}
	json.NewEncoder(w).Encode(msg)
}
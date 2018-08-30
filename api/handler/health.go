package handler

import (
	"encoding/json"
	"net/http"
)

//HealthChecker retorno de status da api
func HealthChecker(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	response, _ := json.Marshal(map[string]interface{}{
		"status": "up",
	})
	w.Write(response)

	return
}

package untils

import (
	"encoding/json"
	"net/http"
)

func JSONRespon(w http.ResponseWriter, status int, data any, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]any{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

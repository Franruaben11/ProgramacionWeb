package response

import (
	"encoding/json"
	"net/http"
)

// JSON escribe `data` como JSON con el codigo de estado indicado.
func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// Error escribe una respuesta de error en el formato {"error": message}.
func Error(w http.ResponseWriter, status int, message string) {
	JSON(w, status, map[string]string{"error": message})
}

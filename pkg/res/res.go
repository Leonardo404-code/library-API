package res

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data == nil {
		if err := json.NewEncoder(w).Encode(`{"error": "data for response not found"}`); err != nil {
			log.Fatal(err)
		}
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

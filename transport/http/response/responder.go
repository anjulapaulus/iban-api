package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Send sets all required fields and write the response.
func Send(w http.ResponseWriter, code int, payload interface{}) {
	write(w, code, payload)
}

// toJSON converts the payload to JSON
func toJSON(payload interface{}) []byte {
	msg, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("JSON Marshalling Error: %v", err)
	}
	return msg
}

// write sets all required fields and write the response.
func write(w http.ResponseWriter, code int, payload interface{}) {

	// set headers
	w.Header().Set("Content-Type", "application/json")

	// set response code
	w.WriteHeader(code)

	// set payload
	w.Write(toJSON(payload))
}

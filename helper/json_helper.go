package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Decoder JSON
func DecoderJson(r *http.Request, target interface{}) error {

	if r.Header.Get("Content-Type") != "application/json" {
		log.Println("Error: invalid content-type", r.Header.Get("Content-Type"))
		return fmt.Errorf("unsupported media type: application/json")
	}

	r.Body = http.MaxBytesReader(nil, r.Body, 1<<20)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(target)
	if err != nil {
		log.Println("cannot decode json")
		return fmt.Errorf("cannot parse json %w", err)
	}

	return nil
}
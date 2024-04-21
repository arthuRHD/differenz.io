package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type Pageable struct {
	Size   int `json:"size"`
	Number int `json:"number"`
}

type Endpoint func(w http.ResponseWriter, r *http.Request) (int, error)

func makeHTTPHandler(endpoint Endpoint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if status, err := endpoint(w, r); err != nil {
			log.Println(err.Error())
			if err := JSONEncode(w, status, err); err != nil {
				panic(err)
			}
		}
	}
}

func JSONEncode(w http.ResponseWriter, httpStatus int, body any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("%+v was not encoded", body)
		return err
	}
	return nil
}

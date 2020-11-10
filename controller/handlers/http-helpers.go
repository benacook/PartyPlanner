package handlers

import (
	"encoding/json"
	"net/http"
)

var (
	headers = map[string]string{
		"Content-Type": "application/json",
		"Cache-Control":"no-store, no-cache, must-revalidate, post-check=0, pre-check=0",
		"Pragma": "no-cache",
		"X-Content-Type-Options": "nosniff",
	}
)

//======================================================================================

//encodeResponseAsJSON takes a blank interface,
//converts it to json and writes it to the io.writer w.
func encodeResponseAsJSON(data interface{}, responseCode int, w http.ResponseWriter) {
	httpAddHeaders(headers, w)
	w.WriteHeader(responseCode)
	enc := json.NewEncoder(w)
	enc.Encode(data)
}

//======================================================================================

//httpAddHeaders takes a http response writer and adds the headers provided.
func httpAddHeaders(headers map[string]string, w http.ResponseWriter)  {
	for k, v := range headers {
		w.Header().Set(k, v)
	}
}

//======================================================================================

func respondWith(key, value string, responseCode int, w http.ResponseWriter)  {
	response := map[string]string{key: value}
	encodeResponseAsJSON(response, responseCode, w)
}

//======================================================================================

func respondWithMessage(message string, responseCode int, w http.ResponseWriter)  {
	respondWith("message", message, responseCode, w)
}

//======================================================================================

func respondWithError(err error, responseCode int, w http.ResponseWriter)  {
	respondWith("error", err.Error(), responseCode, w)
}
package handlers

import "net/http"

//lint:file-ignore U1000 StaticCheck ignore unused code, it's an interface.
type httpHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
	Put(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}


package router

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

var urls = []string{"http://localhost:8081", "http://localhost:8082"}
var i = 0
var n = 2

func getNextUrl() string {
	url := urls[i]
	i += 1
	i %= n
	return url
}

// Handler for the /getData route
func GetLBHandler(w http.ResponseWriter, r *http.Request) {
	// Define the path to the static file

	n_url := getNextUrl()
	url_obj, _ := url.Parse(n_url)
	print(n_url)
	proxy := httputil.NewSingleHostReverseProxy(url_obj)
	proxy.ServeHTTP(w, r)

}

// Handler for the /getData route
func GetDataHandlerA(w http.ResponseWriter, r *http.Request) {
	// Define the path to the static file

	data := []byte("howdy whatsup")
	// Send the file content as the response
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// Handler for the /getData route
func GetDataHandlerB(w http.ResponseWriter, r *http.Request) {
	// Define the path to the static file

	data := []byte("yoooooooo yoooooo")
	// Send the file content as the response
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

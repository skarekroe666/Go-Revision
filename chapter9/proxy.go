package chapter9

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Proxy() {
	targetUrl, _ := url.Parse("https://jsonplaceholder.typicode.com/")
	proxy := httputil.NewSingleHostReverseProxy(targetUrl)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Proxying request for: %s%s", r.Host, r.URL.Path)
		proxy.ServeHTTP(w, r)
	})

	log.Println("Proxy server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

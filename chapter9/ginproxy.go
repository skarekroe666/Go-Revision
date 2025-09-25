package chapter9

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func GinProxy() {
	r := gin.Default()
	defer r.Run()

	r.Any("/proxyPath", proxy)
}

func proxy(c *gin.Context) {
	remote, err := url.Parse("https://example.com")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	proxy.Director = func(r *http.Request) {
		r.Header = c.Request.Header
		r.Host = remote.Host
		r.URL.Scheme = remote.Scheme
		r.URL.Host = remote.Host
		r.URL.Path = c.Param("proxyPath")
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}

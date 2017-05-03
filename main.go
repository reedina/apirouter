package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Prox struct
type Prox struct {
	target *url.URL
	proxy  *httputil.ReverseProxy
}

// New func
func New(target string) *Prox {
	url, _ := url.Parse(target)

	return &Prox{target: url, proxy: httputil.NewSingleHostReverseProxy(url)}
}

func (p *Prox) handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-GoProxy", "GoProxy")

	p.proxy.ServeHTTP(w, r)

}

func main() {
	const (
		defaultPort        = ":4040"
		defaultPortUsage   = "default server port, ':4040'"
		defaultTarget      = "http://127.0.0.1:8080"
		defaultTargetUsage = "default redirect url, 'http://127.0.0.1:8080'"
	)

	// flags
	port := flag.String("port", defaultPort, defaultPortUsage)
	url := flag.String("url", defaultTarget, defaultTargetUsage)

	flag.Parse()

	fmt.Printf("server will run on: %s\n", *port)
	fmt.Printf("redirecting to :%s\n", *url)
	// proxy
	proxy := New(*url)

	// server
	http.HandleFunc("/", proxy.handle)
	http.ListenAndServe(*port, nil)
}

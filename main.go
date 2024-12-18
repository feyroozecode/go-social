// define the package
package main

// import necessary package
import (
	"log"
	"net/http"
)

// define struct for server
type api struct {
	addr string
}

// ServeHTTP implements http.Handler.
func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("Index "))
		case "/users":
			w.Write([]byte("users page"))
		}
	default:
		w.Write([]byte("404 page"))
	}
}

func main() {
	s := &api{addr: ":8080"}
	if err := http.ListenAndServe(s.addr, s); err != nil {
		log.Fatal((err))
	}
}

package pkg

import (
	"log"
	"net/http"
	"time"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		//log.Println(r.Header.Get("Content-Type"))
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}

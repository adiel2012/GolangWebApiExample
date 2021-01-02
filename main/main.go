package main

import (
	"gowebapi/config"
	"log"
	"net/http"
	"os"
)

func main() {
	routes := make(map[string]func(http.ResponseWriter, *http.Request))
	config.Routes(&routes)
	for k, v := range routes {
		//http.HandleFunc(k, v)
		//http.HandleFunc(k, LogMiddleware(v))
		http.HandleFunc(k, MultipleMiddleware(v,
			RequireAuthMiddleware,
			SomeOtherMiddleware,
			LogMiddleware))
	}
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

// https://medium.com/@chrisgregory_83433/chaining-middleware-in-go-918cfbc5644d
type Middleware func(http.HandlerFunc) http.HandlerFunc

func MultipleMiddleware(h http.HandlerFunc, m ...Middleware) http.HandlerFunc {

	if len(m) < 1 {
		return h
	}

	wrapped := h

	// loop in reverse to preserve middleware order
	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}

	return wrapped

}

func LogMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.SetOutput(os.Stdout) // logs go to Stderr by default
		log.Println(r.Method, r.URL)
		h.ServeHTTP(w, r) // call ServeHTTP on the original handler

	})
}

func RequireAuthMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r) // call ServeHTTP on the original handler
	})
}

func SomeOtherMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r) // call ServeHTTP on the original handler
	})
}

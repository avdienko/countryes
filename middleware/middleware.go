package middleware

import (
	"log"
	"net/http"
)

type middleware struct{}

func New() *middleware {
	return &middleware{}
}

func (md *middleware) PanicWrap(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if parnic := recover(); parnic != nil {
				log.Println("PANIC: " + parnic.(string))
			}
		}()

		handler(w, r)
	}
}

package middleware

import (
	"countryes/common/log"
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
				log.Info("PANIC: " + parnic.(string))
			}
		}()

		handler(w, r)
	}
}

func (md *middleware) Log(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Info("Call api:" + r.RequestURI)

		handler(w, r)
	}
}

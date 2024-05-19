package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

func middlewareLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		requestID := middleware.GetReqID(r.Context())
		method := r.Method
		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}
		url := fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		defer func() {
			log.Info().
				Dur("duration", time.Now().Sub(start)).
				Int("response_length", ww.BytesWritten()).
				Int("status", ww.Status()).
				Str("method", method).
				Str("request_id", requestID).
				Str("url", url).
				Send()
		}()

		next.ServeHTTP(ww, r)
	})
}

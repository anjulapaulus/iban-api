package middleware

import (
	"fmt"
	"net/http"

	"github.com/anjulapaulus/iban-api/transport/http/middleware/errors"
	"github.com/anjulapaulus/iban-api/transport/http/response"
)

// RequestCheckerMiddleware validates the request header.
type RequestCheckerMiddleware struct {
	omittedRoutes []string
}

// NewRequestCheckerMiddleware returns a new instance of RequestCheckerMiddleware.
func NewRequestCheckerMiddleware() *RequestCheckerMiddleware {
	return &RequestCheckerMiddleware{
		omittedRoutes: []string{
			"/favicon.ico",
		},
	}
}

// Middleware executes middleware rules of RequestCheckerMiddleware.
func (m *RequestCheckerMiddleware) Middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requestURI := r.RequestURI
		contentType := r.Header.Get("Content-Type")

		// skip omitted routes
		for _, v := range m.omittedRoutes {
			if v == requestURI {
				next.ServeHTTP(w, r)
				return
			}
		}

		// check content type
		if contentType != "application/json" {
			err := errors.NewMiddlewareError("100",
				fmt.Sprintf("API only accepts JSON as Content-Type, '%s' is given", contentType),
				nil,
			)
			response.Error(w, err)
			return
		}

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

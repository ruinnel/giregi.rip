package middleware

import (
	"net/http"
	"net/url"
)

func decodeSlice(slice []string) []string {
	result := make([]string, len(slice))
	for idx, value := range slice {
		unescaped, err := url.QueryUnescape(value)
		if err == nil {
			result[idx] = unescaped
		}
	}
	return result
}

func UrlDecodeMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			for key, values := range request.URL.Query() {
				request.URL.Query()[key] = decodeSlice(values)
			}

			next.ServeHTTP(writer, request)
		})
	}
}

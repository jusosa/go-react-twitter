package middlew

import (
	"github.com/jusosa/go-react-twitter/routers"
	"net/http"
)

func Validate(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		_, _, _, err := routers.ProcessToken(request.Header.Get("Authorization"))
		if err != nil {
			http.Error(writer, "Error validating token: "+err.Error(), http.StatusBadRequest)
		}
		next.ServeHTTP(writer, request)
	}
}

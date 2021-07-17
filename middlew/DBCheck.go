package middlew

import (
	"github.com/jusosa/go-react-twitter/bd"
	"net/http"
)

func DBCheck(next http.HandlerFunc) http.HandlerFunc{
	return func(writer http.ResponseWriter, request *http.Request) {
		if bd.CheckConnection()==0 {
			http.Error(writer, "Db connection refused", 500)
			return
		}
		next.ServeHTTP(writer, request)
	}
}
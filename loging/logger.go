package loging

import "net/http"

type Logger interface {
	Info(request *http.Request)
}

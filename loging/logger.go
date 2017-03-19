package loging

import "net/http"

// Logger type for logging incoming http requests
type Logger interface {
	Info(request *http.Request)
}

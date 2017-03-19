package loging

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
)

// OutputLogger implementation of logger for console logging
type OutputLogger struct {
}

// Info logging requests to the command line output
func (l *OutputLogger) Info(request *http.Request) {
	log.WithFields(log.Fields{
		"Method": request.Method,
		"Path":   request.URL.Path,
		"Query":  request.URL.RawQuery,
	}).Infof("Forwarging to: %s", request.URL.Host)
}

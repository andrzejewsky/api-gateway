package event

import (
	"net/http"

	"github.com/andrzejewsky/api-gatewey/loging"
)

// Notifier notify about new requests
type Notifier struct {
	requests chan *http.Request
	logger   loging.Logger
}

// NewNotifier it creates a new notifier
func NewNotifier(requests chan *http.Request, logger loging.Logger) *Notifier {
	return &Notifier{requests, logger}
}

// Notify it sends new requests to the channel
func (l *Notifier) Notify(request *http.Request) {
	l.requests <- request
}

// Start it's starting logging requests
func (l *Notifier) Start() {
	go func() {
		for request := range l.requests {
			l.logger.Info(request)
		}
	}()
}

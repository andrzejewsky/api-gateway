package event

import (
	"net/http"

	"github.com/andrzejewsky/api-gatewey/loging"
)

type Notifier struct {
	requests chan *http.Request
	logger   loging.Logger
}

func NewNotifier(requests chan *http.Request, logger loging.Logger) *Notifier {
	return &Notifier{requests, logger}
}

func (l *Notifier) Notify(request *http.Request) {
	l.requests <- request
}

func (l *Notifier) Start() {
	go func() {
		for request := range l.requests {
			l.logger.Info(request)
		}
	}()
}

package server

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/andrzejewsky/api-gatewey/api/event"
	"github.com/andrzejewsky/api-gatewey/configuration"
	"github.com/gorilla/mux"
)

// HTTPServer it is a api gateway server
type HTTPServer struct {
	router   *mux.Router
	config   configuration.Config
	notifier *event.Notifier
}

// NewHTTPServer it creates a new one
func NewHTTPServer(router *mux.Router, config configuration.Config, notifier *event.Notifier) *HTTPServer {

	return &HTTPServer{router, config, notifier}
}

// Listen it's starting a listening
func (s *HTTPServer) Listen() {

	for endpoint, destination := range s.config.Get() {

		destinationUrl, _ := url.Parse(destination)

		s.router.HandleFunc(
			endpoint+"{rest:.*}",
			s.redirectToDestinationApi(
				httputil.NewSingleHostReverseProxy(destinationUrl),
				endpoint,
			),
		)

	}
}

func (s *HTTPServer) redirectToDestinationApi(proxy *httputil.ReverseProxy, endpoint string) func(
	writer http.ResponseWriter,
	request *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {
		request.URL.Path = request.URL.Path[len(endpoint):]
		s.notifier.Notify(request)
		proxy.ServeHTTP(writer, request)
	}
}

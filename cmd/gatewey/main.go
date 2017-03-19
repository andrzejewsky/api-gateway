package main

import (
	"flag"
	"net/http"

	"github.com/andrzejewsky/api-gatewey/api/event"
	"github.com/andrzejewsky/api-gatewey/api/server"
	"github.com/andrzejewsky/api-gatewey/configuration"
	"github.com/andrzejewsky/api-gatewey/loging"
	"github.com/gorilla/mux"
)

var endpoints configuration.ArrayFlags
var destinations configuration.ArrayFlags
var listenOn string

func init() {
	flag.StringVar(&listenOn, "listen", "127.0.0.1:8080", "Specify host and port")
	flag.Var(&endpoints, "endpoint", "Endpont for forwarding in the gateway")
	flag.Var(&destinations, "dest", "Destintion URL")
}

func main() {
	flag.Parse()

	notifier := event.NewNotifier(make(chan *http.Request), &loging.OutputLogger{})
	notifier.Start()

	router := mux.NewRouter()

	server.NewHTTPServer(
		router,
		configuration.CreateFromCliParams(endpoints, destinations),
		notifier,
	).Listen()

	http.ListenAndServe(listenOn, router)
}

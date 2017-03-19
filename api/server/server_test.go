package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andrzejewsky/api-gatewey/api/event"
	"github.com/andrzejewsky/api-gatewey/configuration"
	"github.com/andrzejewsky/api-gatewey/loging"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var (
	endpoints     = []string{}
	destinations  = []string{}
	gatewayServer *httptest.Server
	lastResponse  string
)

func TestListeningServer(t *testing.T) {

	givenConfigurationOfApiGatewayRoute("/s1", "service1")
	givenConfigurationOfApiGatewayRoute("/s2", "service2")

	givenApiGatewayWasStarted()

	whenImGoToTheRoute("/s1/test/?a=1")
	thenServiceResponseWillBeEqualsTo("service1/test/a=1", t)

	whenImGoToTheRoute("/s2/test/?b=1")
	thenServiceResponseWillBeEqualsTo("service2/test/b=1", t)
}

func givenConfigurationOfApiGatewayRoute(endpoint, serviceName string) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, serviceName+r.URL.Path+r.URL.RawQuery)
	}))

	endpoints = append(endpoints, endpoint)
	destinations = append(destinations, ts.URL)
}

func givenApiGatewayWasStarted() {
	router := mux.NewRouter()

	notifier := event.NewNotifier(make(chan *http.Request), &loging.OutputLogger{})
	notifier.Start()

	NewHTTPServer(
		router,
		configuration.CreateFromCliParams(endpoints, destinations),
		notifier,
	).Listen()

	gatewayServer = httptest.NewServer(router)
}

func whenImGoToTheRoute(route string) {
	res, _ := http.Get(gatewayServer.URL + route)

	greeting, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	lastResponse = fmt.Sprintf("%s", greeting)
}

func thenServiceResponseWillBeEqualsTo(serviceResponse string, t *testing.T) {
	assert.Equal(t, serviceResponse, lastResponse)
}

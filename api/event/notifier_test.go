package event

import (
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

type LoggerMock struct {
	mock.Mock
}

func (l *LoggerMock) Info(request *http.Request) {
	l.Called(request)
}

var (
	logger   *LoggerMock
	request  *http.Request
	notifier *Notifier
)

func TestNotify(t *testing.T) {

	givenRequest(new(http.Request))
	givenLogger(new(LoggerMock))
	givenNotifier()
	thenCountOfLoggedRequestShouldBeEqualTo(3)
	whenIAmNotifyingAboutTheRequest(request)
	whenIAmNotifyingAboutTheRequest(request)
	whenIAmNotifyingAboutTheRequest(request)

	logger.AssertExpectations(t)
}

func givenRequest(requestMock *http.Request) {
	request = requestMock
}

func givenLogger(loggerMock *LoggerMock) {
	logger = loggerMock
}

func givenNotifier() {
	notifier = NewNotifier(make(chan *http.Request), logger)
	notifier.Start()
}

func whenIAmNotifyingAboutTheRequest(request *http.Request) {
	notifier.Notify(request)
}

func thenCountOfLoggedRequestShouldBeEqualTo(count int) {
	logger.On("Info", request).Times(count)
}

package main

import (
	"os"

	"github.com/trakerr-io/trakerr-go/src/trakerr"
)

func main() {
	var client *trakerr.TrakerrClient
	if len(os.Args) >= 1 {
		client = trakerr.NewTrakerrClientWithDefaults(
			os.Args[1],
			"1.0",
			"development")
	} else {
		client = trakerr.NewTrakerrClientWithDefaults(
			"API Key here",
			"1.0",
			"development")
	}
	
	//Option-1: Global error handling

	appEvent := client.NewErrorEvent("Error")
	// set any custom data on appEvent
	appEvent.CustomProperties.StringData.CustomData1 = "foo"
	appEvent.CustomProperties.StringData.CustomData2 = "bar"
	appEvent.EventUser = "john@trakerr.io"
	appEvent.EventSession = "12"

	ts := TestSession{client, appEvent}
	buf := []int{1, 2, 3}
	te := TestError{}
	te.BufferOverflowError(buf, 4, ts)
	


	// Option-2: send error
	/*err := errors.New("Something bad happened here")
	client.SendError(err)

	// Option-3: send error with custom properties
	appEventWithErr := client.CreateAppEventFromError(err)

	// set any custom data on appEvent
	appEventWithErr.CustomProperties.StringData.CustomData1 = "foo"
	appEventWithErr.CustomProperties.StringData.CustomData2 = "bar"

	client.SendEvent(appEventWithErr)

	// Option-4: send event manually
	appEvent := client.NewAppEvent("Info", "SomeType", "SomeMessage")

	// set any custom data on appEvent
	appEvent.CustomProperties.StringData.CustomData1 = "foo"
	appEvent.CustomProperties.StringData.CustomData2 = "bar"

	client.SendEvent(appEvent)*/
<<<<<<< HEAD

	//Option 4: Global error handling

	appEvent := client.NewErrorEvent()
	// set any custom data on appEvent
	appEvent.CustomProperties.StringData.CustomData1 = "foo"
	appEvent.CustomProperties.StringData.CustomData2 = "bar"
	appEvent.EventUser = "john@user.com"
	appEvent.EventSession = "12"

	ts := TestSession{client, appEvent}
	buf := []int{1, 2, 3}
	te := TestError{}
	te.BufferOverflowError(buf, 4, ts)
=======
>>>>>>> e3707b8d8ba4a699ff681f3fe3ac98e05f3c75d3
}

type TestSession struct {
	client   *trakerr.TrakerrClient
	appEvent *trakerr.AppEvent
}

type TestError struct {
}

//BufferOverflowError ...
func (testError *TestError) BufferOverflowError(buf []int, i int, session TestSession) (x int) {
	//defer client.Recover()
	defer session.client.RecoverWithAppEvent("Error", session.appEvent)

	x = buf[i]
	return x
}

package stream

import (
	"log"
	"net/http"
)

func ExampleStream() {
	// Stream objects implement the http.Handler interface,
	// allowing to use them with the net/http package like so:
	strm := NewStream()
	http.Handle("/camera", strm)
	// Then push new JPEG frames to the connected clients using stream.UpdateJPEG().
	go log.Fatal(http.ListenAndServe(":8080", nil))

	//UpdateJPEG(somethingImage)
}

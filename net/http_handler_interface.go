package main

// START OMIT
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

// END OMIT

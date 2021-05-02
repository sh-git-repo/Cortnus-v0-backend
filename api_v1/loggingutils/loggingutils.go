package loggingutils

import "fmt"

func LogResponse(route string, status int) {
	/*
		Example for a logger
	*/
	fmt.Printf("Response for %s - [%d]\n", route, status)
}

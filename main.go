package main

import "net/http"

func main() {
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		logPrintln("Hello Mahesh")
	})
	http.ListenerAndServe(":8080", nil)
}

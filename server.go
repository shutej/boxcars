package boxcars

import (
	"fmt"
	"net/http"
)

func Listen(port int) {
	debug("Starting at %d", port)
	http.HandleFunc("/", OnRequest)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		debug("Fatal: %v", err)
	}
}

func ListenTLS(port int, certFile, keyFile string) {
	debug("Starting at %d", port)
	http.HandleFunc("/", OnRequest)
	err := http.ListenAndServeTLS(fmt.Sprintf(":%d", port), certFile, keyFile, nil)
	if err != nil {
		debug("Fatal: %v", err)
	}
}

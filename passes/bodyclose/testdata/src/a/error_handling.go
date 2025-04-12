package a

import (
	"fmt"
	"io"
	"net/http"
)

func errorHandlingClose() {
	resp, err := http.Get("http://example.com/") // OK
	if err != nil {
		// handle error
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(resp.Body)
}

func errorHandlingCloseInline() {
	resp, err := http.Get("http://example.com/") // OK
	if err != nil {
		// handle error
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}()
}

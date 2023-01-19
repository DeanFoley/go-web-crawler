package app

import (
	"fmt"
	"io"
	"net/http"
)

func GrabWebpage(client http.Client, url string, dataChan chan []byte, doneChan chan struct{}, errorChan chan error) {
	resp, err := client.Get(url)
	if err != nil {
		errorChan <- fmt.Errorf("problem grabbing page: %v", err)
		return
	}
	defer resp.Body.Close()

	html, err := io.ReadAll(resp.Body)
	if err != nil {
		errorChan <- fmt.Errorf("problem reading result from url: %v", err)
		return
	}
	doneChan <- struct{}{}
	dataChan <- html
}

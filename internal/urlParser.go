package app

import "net/url"

func ValidateUrl(requestedUrl string, doneChan chan struct{}, errorChan chan error) {
	_, err := url.ParseRequestURI(requestedUrl)
	if err != nil {
		errorChan <- err
		return
	}
	doneChan <- struct{}{}
}

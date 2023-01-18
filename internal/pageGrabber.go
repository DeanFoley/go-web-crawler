package app

import (
	"fmt"
	"io"
	"net/http"
)

func GrabWebpage(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("problem grabbing page: %v", err)
	}
	defer resp.Body.Close()

	html, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("problem reading result from url: %v", err)
	}
	return html, nil
}
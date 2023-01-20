package app

import (
	"io"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

var wg sync.WaitGroup

func ExtractAnchors(webpage []byte, anchorsChan chan string, doneChan chan struct{}, errorChan chan error) {
	stringifiedWebpage := string(webpage)

	tkn := html.NewTokenizer(strings.NewReader(stringifiedWebpage))

	anchorsQueue := make(chan string)

	go func(anchorsQueue chan string) {
		for anchor := range anchorsQueue {
			anchorsChan <- anchor
		}
	}(anchorsQueue)

	for {
		tt := tkn.Next()

		switch {
		case tt == html.ErrorToken:
			if tkn.Err() == io.EOF {
				wg.Wait()
				close(anchorsChan)
				doneChan <- struct{}{}
				return
			} else {
				errorChan <- tkn.Err()
				return
			}

		case tt == html.StartTagToken:
			t := tkn.Token()
			if t.Data == "a" {
				for _, value := range t.Attr {
					if value.Key == "href" {
						wg.Add(1)
						anchorsQueue <- value.Val
					}
				}
			}
		}
	}
}

func FormatAnchors(anchorStringsChan chan string, bytesChannel chan []byte, doneChannel chan struct{}) {
	for url := range anchorStringsChan {
		bytesChannel <- []byte(url + "\n")
		wg.Done()
	}
	close(bytesChannel)
	doneChannel <- struct{}{}
}

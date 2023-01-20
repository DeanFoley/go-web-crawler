package app

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func StartWorkflow(uri string, outputToConsole bool) {
	doneChan := make(chan struct{})
	errorChan := make(chan error)

	go func() {
		err := <-errorChan
		fmt.Println(err)
		os.Exit(0)
	}()

	go ValidateUrl(uri, doneChan, errorChan)
	<-doneChan

	tr := &http.Transport{}
	defer tr.CloseIdleConnections()
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 10,
	}

	dataChan := make(chan []byte, 1)
	go GrabWebpage(*client, uri, dataChan, doneChan, errorChan)
	go func() {
		for {
			select {
			case <-doneChan:
				fmt.Printf("\nSuccessfully downloaded page!\n")
				return
			default:
				for _, r := range `-\|/` {
					fmt.Printf("\rWaiting for page... %c", r)
					time.Sleep(100000000)
				}
			}
		}
	}()
	webpage := <-dataChan

	anchorsChan := make(chan string)
	go ExtractAnchors(webpage, anchorsChan, doneChan, errorChan)

	anchorBytes := make(chan []byte)
	go FormatAnchors(anchorsChan, anchorBytes, doneChan)

	go PrintResults(anchorBytes, outputToConsole, doneChan, errorChan)

	<-doneChan
	<-doneChan
	<-doneChan

	fmt.Println("Thanks for using my cool tool!")
}

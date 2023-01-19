package app

import (
	"fmt"
	"os"
	"time"
)

func PrintResults(anchorBytes chan []byte, doneChan chan struct{}, errorChan chan error) {
	workingDirectory, err := os.Getwd()
	if err != nil {
		errorChan <- err
		return
	}
	filename := fmt.Sprintf("%s/anchors-%s.txt", workingDirectory, time.Now().Format(time.RFC3339))
	file, err := os.Create(filename)
	if err != nil {
		errorChan <- err
		return
	}
	defer file.Close()
	printerSignal := make(chan struct{})
	go func() {
		for url := range anchorBytes {
			_, err = file.Write(url)
			if err != nil {
				errorChan <- err
				return
			}
		}
		printerSignal <- struct{}{}
	}()
	<-printerSignal
	fmt.Printf("\nAnchors printed to %s!\n", filename)
	doneChan <- struct{}{}

}

package app

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var output *os.File
var filename string

func PrintResults(anchorBytes chan []byte, outputToConsole bool, doneChan chan struct{}, errorChan chan error) {
	if outputToConsole {
		output = os.Stdout
		fmt.Print("\n-----------------------------\nURLs:\n-----------------------------\n")
	} else {
		workingDirectory, err := os.Getwd()
		if err != nil {
			errorChan <- err
			return
		}
		filename = fmt.Sprintf("%s/anchors-%s.txt", workingDirectory, time.Now().Format(time.RFC3339))
		output, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			errorChan <- err
			return
		}
	}

	writer := bufio.NewWriter(output)

	printerSignal := make(chan struct{})
	go func() {
		for url := range anchorBytes {
			_, err := writer.Write(url)
			if err != nil {
				errorChan <- err
				return
			}
		}
		writer.Flush()
		printerSignal <- struct{}{}
	}()
	<-printerSignal
	if !outputToConsole {
		fmt.Printf("\nAnchors printed to %s!\n", filename)
	} else {
		fmt.Println("-----------------------------")
	}
	doneChan <- struct{}{}

}

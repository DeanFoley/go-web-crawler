package app

import (
	"os"
	"reflect"
	"testing"
)

func TestExtractAnchors(t *testing.T) {
	data, err := os.ReadFile("../test/data/test_page.html")
	if err != nil {
		t.Fatal()
	}

	expected := []string{
		"https://www.google.com",
		"https://www.golang.org",
		"https://www.github.com",
	}

	anchorsChan := make(chan string)
	doneChan := make(chan struct{})
	errorChan := make(chan error)
	actual := []string{}

	go ExtractAnchors(data, anchorsChan, doneChan, errorChan)
	for {
		go func() {
			for val := range anchorsChan {
				actual = append(actual, val)
			}
		}()

		select {
		case <-errorChan:
			t.Fatal()
		case <-doneChan:
			if len(actual) != 3 {
				t.Fatal()
			}

			if !reflect.DeepEqual(expected, actual) {
				t.Fatal()
			}
			return
		}
	}
}

func TestFormatAnchors(t *testing.T) {
	expected := []byte("biscuits\ncoffee\nporkpies\n")
	rawData := []string{
		"biscuits",
		"coffee",
		"porkpies",
	}

	doneChan := make(chan struct{})
	bytesChan := make(chan []byte)

	stringsChan := make(chan string)

	actual := make([]byte, 0)

	go func() {
		for url := range bytesChan {
			actual = append(actual, url...)
		}
	}()

	go FormatAnchors(stringsChan, bytesChan, doneChan)
	go func() {
		for _, val := range rawData {
			stringsChan <- val
		}
		close(stringsChan)
	}()
	<-doneChan

	if !reflect.DeepEqual(expected, actual) {
		t.Fatal()
	}
}

func BenchmarkExtractAnchors(b *testing.B) {
	data, err := os.ReadFile("../test/data/test_page.html")
	if err != nil {
		b.Fatal()
	}

	anchorsChan := make(chan string)
	doneChan := make(chan struct{})
	errorChan := make(chan error)

	go func() {
		for {
			<-doneChan
		}
	}()

	go func() {
		for {
			<-anchorsChan
		}
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go ExtractAnchors(data, anchorsChan, doneChan, errorChan)
	}
}

func BenchmarkFormatAnchors(b *testing.B) {
	doneChan := make(chan struct{})
	bytesChan := make(chan []byte)
	stringsChan := make(chan string)

	go func() {
		for {
			stringsChan <- "Hello"
		}
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go FormatAnchors(stringsChan, bytesChan, doneChan)
	}
}

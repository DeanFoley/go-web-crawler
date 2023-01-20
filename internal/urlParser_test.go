package app

import (
	"testing"
)

func TestValidURL(t *testing.T) {
	uri := "https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package"

	doneChan := make(chan struct{})
	errorChan := make(chan error)

	go ValidateUrl(uri, doneChan, errorChan)
	for {
		select {
		case <-errorChan:
			t.Fatal()
		case <-doneChan:
			return
		}
	}
}

func TestInvalidURL(t *testing.T) {
	uri := "https://digitalocean^A/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package"

	doneChan := make(chan struct{})
	errorChan := make(chan error)

	go ValidateUrl(uri, doneChan, errorChan)
	for {
		select {
		case <-errorChan:
			return
		case <-doneChan:
			t.Fatal()
		}
	}
}

func BenchmarkValidURL(b *testing.B) {
	uri := "https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package"

	doneChan := make(chan struct{})
	errorChan := make(chan error)

	go func() {
		for {
			<-doneChan
		}
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ValidateUrl(uri, doneChan, errorChan)
	}
}

func BenchmarkInvalidURL(b *testing.B) {
	uri := "https://digitalocean^A/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package"

	doneChan := make(chan struct{})
	errorChan := make(chan error)

	go func() {
		for {
			<-errorChan
		}
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ValidateUrl(uri, doneChan, errorChan)
	}
}

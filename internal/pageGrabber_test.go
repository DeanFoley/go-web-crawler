package app

import (
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

var vortexPage = "../test/data/vortex.html"

func Test_GrabWebpage(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page, err := os.ReadFile(vortexPage)
		if err != nil {
			panic(err)
		}
		w.Write([]byte(page))
	}))
	defer svr.Close()

	expected, err := os.ReadFile(vortexPage)
	if err != nil {
		t.Fatal()
	}

	tr := &http.Transport{}
	defer tr.CloseIdleConnections()
	cl := &http.Client{
		Transport: tr,
	}

	bytesChan := make(chan []byte)
	doneChan := make(chan struct{})
	errorChan := make(chan error)

	go GrabWebpage(*cl, svr.URL, bytesChan, doneChan, errorChan)
	for {
		select {
		case <-errorChan:
			t.Fatal()
		case <-doneChan:
			actual := <-bytesChan

			if result := reflect.DeepEqual(expected, actual); !result {
				t.Fatal()
			}
			return
		}
	}
}

func Test_GrabWebpageInvalid(t *testing.T) {
	bytesChan := make(chan []byte)
	doneChan := make(chan struct{})
	errorChan := make(chan error)

	tr := &http.Transport{}
	defer tr.CloseIdleConnections()
	cl := &http.Client{
		Transport: tr,
	}

	go GrabWebpage(*cl, "www.vortex.com", bytesChan, doneChan, errorChan)
	for {
		select {
		case <-errorChan:
			return
		case <-doneChan:
			t.Fatal()
		}
	}

}

func Benchmark_GrabWebpage(b *testing.B) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page, err := os.ReadFile(vortexPage)
		if err != nil {
			panic(err)
		}
		w.Write([]byte(page))
	}))
	defer svr.Close()

	tr := &http.Transport{}
	defer tr.CloseIdleConnections()
	cl := &http.Client{
		Transport: tr,
	}

	bytesChan := make(chan []byte, 1)
	doneChan := make(chan struct{}, 1)
	errorChan := make(chan error)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go GrabWebpage(*cl, svr.URL, bytesChan, doneChan, errorChan)
		<-doneChan
		b.StopTimer()
		b.StartTimer()
	}
}

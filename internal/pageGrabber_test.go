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

	actual, err := GrabWebpage(svr.URL)
	if err != nil {
		t.Fatal()
	}

	if result := reflect.DeepEqual(expected, actual); !result {
		t.Fatal()
	}
}

func Test_GrabWebpageInvalid(t *testing.T) {
	_, err := GrabWebpage("www.vortex.com")
	if err == nil {
		t.Fatal()
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
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GrabWebpage(svr.URL)
	}
}

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

	actual, err := ExtractAnchors(data)

	if err != nil {
		t.Fatal()
	}

	if len(actual) != 3 {
		t.Fatal()
	}

	if !reflect.DeepEqual(expected, actual) { 
		t.Fatal()
	}
}

func TestFormatAnchors(t *testing.T) {
	expected := []byte("biscuits\ncoffee\nporkpies\n")
	rawData := []string{
		"biscuits",
		"coffee",
		"porkpies",
	}
	actual := FormatAnchors(rawData)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatal()
	}
}

func BenchmarkExtractAnchors(b *testing.B) {
	data, err := os.ReadFile("../test/data/test_page.html")
	if err != nil {
		b.Fatal()
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ExtractAnchors(data)
	}
}

func BenchmarkFormatAnchors(b *testing.B) {
	rawData := []string{
		"biscuits",
		"coffee",
		"porkpies",
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FormatAnchors(rawData)
	}
}
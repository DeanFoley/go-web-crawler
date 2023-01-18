package app

import (
	"testing"
)

func TestValidURL(t *testing.T) {
	uri := "https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package"

	_, error := ValidateUrl(uri)
	if error != nil {
		t.Fatal();
	}
}

func TestInvalidURL(t *testing.T) {
	uri := "https://digitalocean^A/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package"

	_, error := ValidateUrl(uri)
	if error == nil {
		t.Fatal();
	}
}

func BenchmarkValidURL(b *testing.B) {
	uri := "https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package"

	b.ResetTimer()
	for i := 0; i < b.N; i ++ {
		ValidateUrl(uri)
	}
}

func BenchmarkInvalidURL(b *testing.B) {
	uri := "https://digitalocean^A/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package"

	b.ResetTimer()
	for i := 0; i < b.N; i ++ {
		ValidateUrl(uri)
	}
}
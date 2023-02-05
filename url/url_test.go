package url_test

import (
	"testing"

	"github.com/hueristiq/hqgoutils/url"
)

func TestDefaultScheme(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{input: "localhost", output: "http://localhost"},
		{input: "example.com", output: "http://example.com"},
		{input: "https://example.com", output: "https://example.com"},
		{input: "://example.com", output: "http://example.com"},
		{input: "//example.com", output: "http://example.com"},
	}

	for index := range tests {
		test := tests[index]

		URL := url.DefaultScheme(test.input, "http")

		if URL != test.output {
			t.Errorf(`"%s": got "%s", want "%v"`, test.input, URL, test.output)
		}
	}
}

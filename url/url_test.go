package url_test

import (
	"testing"

	"github.com/hueristiq/hqgoutils/url"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input  string
		output url.URL
	}{
		{
			input: "https://sub.example.com:8080",
			output: url.URL{
				ETLDPlusOne: "example.com",
				Subdomain:   "sub",
				TLD:         "com",
				Port:        "8080",
				Extension:   "",
			},
		},
		{
			input: "https://sub.example.com:8080/path/to/file.txt",
			output: url.URL{
				ETLDPlusOne: "example.com",
				Subdomain:   "sub",
				TLD:         "com",
				Port:        "8080",
				Extension:   "txt",
			},
		},
	}

	for index := range tests {
		test := tests[index]

		URL, err := url.Parse(test.input)
		if err != nil {
			t.Error(err)
		}

		if URL.ETLDPlusOne != test.output.ETLDPlusOne {
			t.Errorf(`"%s": got "%s", want "%v"`, test.input, URL.ETLDPlusOne, test.output.ETLDPlusOne)
		}

		if URL.Subdomain != test.output.Subdomain {
			t.Errorf(`"%s": got "%s", want "%v"`, test.input, URL.Subdomain, test.output.Subdomain)
		}

		if URL.TLD != test.output.TLD {
			t.Errorf(`"%s": got "%s", want "%v"`, test.input, URL.TLD, test.output.TLD)
		}

		if URL.Port != test.output.Port {
			t.Errorf(`"%s": got "%s", want "%v"`, test.input, URL.Port, test.output.Port)
		}
	}
}

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

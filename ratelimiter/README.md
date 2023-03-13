# hqgoutils/ratelimiter

A [Go(Golang)](https://golang.org/) package for handling rate limiting.

## Resources

* [Features](#features)
* [Usage](#usage)

## Installation

```
go get -v -u github.com/hueristiq/hqgoutils/ratelimiter
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/hueristiq/hqgoutils/ratelimiter"
)

func main() {
	options := &ratelimiter.Options{
		RequestsPerMinute: 40,
		MinimumDelayInSeconds: 2,
	}

	limiter := ratelimiter.New(options)

	// Make 10 requests and ensure that they are rate limited.
	for i := 1; i <= 10; i++ {
		limiter.Wait()
		fmt.Printf("Request %d made at %v\n", i, time.Now())
	}
}
```
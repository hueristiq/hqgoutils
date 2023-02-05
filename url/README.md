# Hueristiq's Go URLs Utility Package

A [Golang](http://golang.org/) package for handling URLs.

## Features

1. [Parsing URL](#parsing-url)

### Parsing URL


#### Usage

```go
import "github.com/hueristiq/hqgoutils/url"

func main() {
    url, _ := url.Parse(url.Options{URL: "example.com", DefaultScheme: "http"})
    // url.Scheme == "http"
    // url.Host == "example.com"

    fmt.Print(url)
    // Prints http://example.com
}
```

#### Difference between `github.com/hueristiq/hqgoutils/url` and `net/url`

<table>
<thead>
<tr>
<th><a href="https://godoc.org/github.com/hueristiq/hqgoutils/url#Parse">github.com/hueristiq/hqgoutils/url</a></th>
<th><a href="https://golang.org/pkg/net/url/#Parse">net/url</a></th>
</tr>
</thead>
<tr>
<td>
<pre>
url.Parse("example.com")

&url.URL{
   Scheme:  "http",
   Host:    "example.com",
   Path:    "",
}
</pre>
</td>
<td>
<pre>
url.Parse("example.com")

&url.URL{
   Scheme:  "",
   Host:    "",
   Path:    "example.com",
}
</pre>
</td>
</tr>
<tr>
<td>
<pre>
url.Parse("localhost:8080")

&url.URL{
   Scheme:  "http",
   Host:    "localhost:8080",
   Path:    "",
   Opaque:  "",
}
</pre>
</td>
<td>
<pre>
url.Parse("localhost:8080")

&url.URL{
   Scheme:  "localhost",
   Host:    "",
   Path:    "",
   Opaque:  "8080",
}
</pre>
</td>
</tr>
<tr>
<td>
<pre>
url.Parse("user.local:8000/path")

&url.URL{
   Scheme:  "http",
   Host:    "user.local:8000",
   Path:    "/path",
   Opaque:  "",
}
</pre>
</td>
<td>
<pre>
url.Parse("user.local:8000/path")

&url.URL{
   Scheme:  "user.local",
   Host:    "",
   Path:    "",
   Opaque:  "8000/path",
}
</pre>
</td>
</tr>
</table>

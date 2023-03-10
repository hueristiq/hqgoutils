package url

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/publicsuffix"
)

// A URL represents a parsed URL (technically, a URI reference).
//
// The general form represented is:
//
//	[scheme:][//[userinfo@]host][/]path[?query][#fragment]
//
// URLs that do not start with a slash after the scheme are interpreted as:
//
//	scheme:opaque[?query][#fragment]
//
// https://sub.example.com:8080
type URL struct {
	*url.URL
	// Scheme      string    // e.g https
	// Opaque      string    // encoded opaque data
	// User        *Userinfo // username and password information
	// Host        string    // e.g. sub.example.com, sub.example.com:8080
	// Path        string    // path (relative paths may omit leading slash)
	// RawPath     string    // encoded path hint (see EscapedPath method)
	// OmitHost    bool      // do not emit empty host (authority)
	// ForceQuery  bool      // append a query ('?') even if RawQuery is empty
	// RawQuery    string    // encoded query values, without '?'
	// Fragment    string    // fragment for references, without '#'
	// RawFragment string    // encoded fragment hint (see EscapedFragment method)
	Domain      string // e.g. sub.example.com
	ETLDPlusOne string // e.g. example.com
	Subdomain   string // e.g. sub
	RootDomain  string // e.g. example
	TLD         string // e.g. com
	Port        string // e.g. 8080
}

// Parse parses a raw url into a URL structure.
//
// It uses the  `net/url`'s Parse() internally, but it slightly changes its behavior:
//  1. It forces the default scheme if the url doesnt have a scheme and port to http
//  2. It favors absolute paths over relative ones, thus "example.com"
//     is parsed into url.Host instead of url.Path.
//  3. It lowercases the Host (not only the Scheme).
func Parse(rawURL string) (parsedURL *URL, err error) {
	var (
		defaultScheme string = "http"
	)

	rawURL = DefaultScheme(rawURL, defaultScheme)
	parsedURL = &URL{}

	parsedURL.URL, err = url.Parse(rawURL)
	if err != nil {
		err = fmt.Errorf("[hqgoutils/url] %s", err)

		return
	}

	// Host = Domain + Port
	for i := len(parsedURL.URL.Host) - 1; i >= 0; i-- {
		if parsedURL.URL.Host[i] == ':' {
			parsedURL.Domain = parsedURL.URL.Host[:i]
			parsedURL.Port = parsedURL.URL.Host[i+1:]
			break
		} else if parsedURL.URL.Host[i] < '0' || parsedURL.URL.Host[i] > '9' {
			parsedURL.Domain = parsedURL.URL.Host
		}
	}

	// ETLDPlusOne
	parsedURL.ETLDPlusOne, err = publicsuffix.EffectiveTLDPlusOne(parsedURL.Domain)
	if err != nil {
		err = fmt.Errorf("[hqgoutils/url] %s", err)

		return
	}

	// RootDomain + TLD
	i := strings.Index(parsedURL.ETLDPlusOne, ".")
	parsedURL.RootDomain = parsedURL.ETLDPlusOne[0:i]
	parsedURL.TLD = parsedURL.ETLDPlusOne[i+1:]

	// Subdomain
	if rest := strings.TrimSuffix(parsedURL.Domain, "."+parsedURL.ETLDPlusOne); rest != parsedURL.Domain {
		parsedURL.Subdomain = rest
	}

	return
}

// DefaultScheme forces default scheme to `http` scheme, so net/url.Parse() doesn't
// put both host and path into the (relative) path.
func DefaultScheme(URL, scheme string) (URLWithScheme string) {
	URLWithScheme = URL

	// e.g //example.com
	if strings.Index(URLWithScheme, "//") == 0 {
		URLWithScheme = scheme + ":" + URLWithScheme
	}

	// e.g ://example.com
	if strings.Contains(URLWithScheme, "://") && !strings.HasPrefix(URLWithScheme, "http") {
		URLWithScheme = scheme + URLWithScheme
	}

	// e.g example.com, localhost
	if !strings.Contains(URLWithScheme, "://") {
		URLWithScheme = scheme + "://" + URLWithScheme
	}

	return
}

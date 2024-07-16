package client

import "net/http"

// wrapTransport is a simple warpper of the provided transport to add more headers into the http request.
type wrapTransport struct {
	http.RoundTripper
	headers map[string]string
}

// RoundTrip .
func (p *wrapTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if len(p.headers) != 0 {
		for k, v := range p.headers {
			req.Header.Set(k, v)
		}
	}
	return p.RoundTripper.RoundTrip(req)
}

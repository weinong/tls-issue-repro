package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
)

var (
	url                 string
	withTLS12MaxVersion bool
)

func init() {
	flag.StringVar(&url, "url", "", "url to http get in https")
	flag.BoolVar(&withTLS12MaxVersion, "withTLS12MaxVersion", false, "enable TLS 1.2 as max version")
}

func main() {
	flag.Parse()

	tlsClientConfig := &tls.Config{
		Renegotiation: tls.RenegotiateFreelyAsClient,
		MinVersion:    tls.VersionTLS12,
	}
	if withTLS12MaxVersion {
		fmt.Println("setting TLS max version to be 1.2")
		tlsClientConfig.MaxVersion = tls.VersionTLS12
	} else {
		fmt.Println("not setting tls max version. this will include TLS 1.2 and 1.3")
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSNextProto:    make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
			TLSClientConfig: tlsClientConfig,
		},
	}

	resp, err := client.Get("https://" + url)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println("http response status code", resp.StatusCode)
	fmt.Println("http response header:", resp.Header)
}

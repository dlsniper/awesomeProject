package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type Pricer func(client *http.Client, ticker string) float64

func CurrentPrice(client *http.Client, ticker string) float64 {
	resp, err := client.Get("http://stock-service.com/currentprice/" + ticker)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return parsePriceFromBody(body)
}

func parsePriceFromBody(bytes []byte) float64 {
	return 0
}

func analyze(se Pricer, client *http.Client, ticker string, maxTradePrice float64) (bool, error) {
	currentPrice := se(client, ticker)
	var hasTraded bool
	var err error
	if currentPrice <= maxTradePrice {
		err = doTrade(ticker, currentPrice)
		if err == nil {
			hasTraded = true
		}
	}
	return hasTraded, err
}

func doTrade(s string, f float64) error {
	return nil
}

type fakeTransport struct{}

func (*fakeTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	// Exercise: make this a dynamic return, based on the request
	body := bytes.NewBufferString(`{"hello": "world"}`)

	w := &http.Response{
		Status:           "200 OK",
		StatusCode:       200,
		Proto:            "HTTP/1.0",
		ProtoMajor:       1,
		ProtoMinor:       0,
		Header:           map[string][]string{"Content-Type": {"application/json"}},
		Body:             ioutil.NopCloser(body),
		ContentLength:    0,
		TransferEncoding: nil,
		Close:            false,
		Uncompressed:     true,
		Trailer:          nil,
		Request:          r,
		TLS:              nil,
	}
	return w, nil
}

func main() {
	// The real transport
	// tr := &http.Transport{}

	tr := &fakeTransport{}

	client := &http.Client{
		Transport: tr,
	}

	analyze(CurrentPrice, client,"GOOGL", 10)
}

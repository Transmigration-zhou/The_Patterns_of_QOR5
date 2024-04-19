package main

import (
	"encoding/json"
	"fmt"
	"github.com/ysyesilyurt/go-restclient/restclient"
	"net/http"
	"net/url"
	"time"
)

func main() {
	var response = json.RawMessage{}
	req, reqErr := restclient.RequestBuilder().
		Scheme("https").
		Host("httpbin.org").
		PathElements([]string{"anything"}).
		QueryParams(&url.Values{"foo": []string{"bar"}}).
		Header(&http.Header{"X-Custom-Header": []string{"value1"}}).
		Auth(&restclient.BasicAuthenticator{Username: "username", Password: "password"}).
		Timeout(30 * time.Second).
		ResponseReference(&response).
		Build()
	if reqErr != nil {
		panic(reqErr)
	}
	reqErr = req.Get()
	if reqErr != nil {
		panic(reqErr)
	}
	fmt.Println(string(response))
}

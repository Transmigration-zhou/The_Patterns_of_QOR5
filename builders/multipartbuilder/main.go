package main

import (
	"encoding/json"
	"fmt"
	"github.com/mxmCherry/multipartbuilder"
	"github.com/ysyesilyurt/go-restclient/restclient"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	builder := multipartbuilder.New()
	builder.AddField("field", "value").
		AddReader("reader", "some.txt", strings.NewReader("Some reader")).
		AddFile("file", "./main.go")
	contentType, bodyReader := builder.Build()
	defer bodyReader.Close()

	var response = json.RawMessage{}
	req, reqErr := restclient.RequestBuilder().
		Scheme("https").
		Host("httpbin.org").
		PathElements([]string{"anything"}).
		QueryParams(&url.Values{"foo": []string{"bar"}}).
		Header(&http.Header{"X-Custom-Header": []string{"value1"}}).
		Auth(&restclient.BasicAuthenticator{Username: "username", Password: "password"}).
		Timeout(30 * time.Second).
		Header(&http.Header{"Content-Type": []string{contentType}}).
		Body(bodyReader).
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

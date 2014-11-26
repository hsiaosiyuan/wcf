package main

import (
	"net/http"
	"io/ioutil"
	"net/url"
)

type Forwarder struct {
	app        wcfApp
	request    http.Request
	respWriter http.ResponseWriter
}

func NewForwarder(w http.ResponseWriter, r *http.Request) *Forwarder {
	nf := new(Forwarder)
	nf.request = r
	nf.respWriter = w

	return nf
}

func (f *Forwarder) Do() {
	r := f.request
	respWriter := f.respWriter

	proxyClient := new(http.Client)
	proxyRequest := new(http.Request)

	proxyRequest.URL = url.Parse(f.app.ServiceUrl);
	proxyRequest.Header = r.Header
	proxyRequest.Body = r.Body

	for _, v := range r.Cookies() {
		proxyRequest.AddCookie(v)
	}

	if proxyResp, err := proxyClient.Do(proxyRequest); err == nil {
		for k, v := range proxyResp.Header {
			for _, vv := range v {
				respWriter.Header().Add(k, vv)
			}
		}

		resp, _ := ioutil.ReadAll(proxyResp.Body)
		respWriter.Write(resp)
	}
}

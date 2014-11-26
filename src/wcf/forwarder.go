package wcf

import (
	"net/http"
	"io/ioutil"
	"net/url"
	"bytes"
)

type Forwarder struct {
	app        wcfApp
	request    http.Request
	respWriter http.ResponseWriter
}

func NewForwarder(w http.ResponseWriter, r *http.Request) *Forwarder {
	nf := new(Forwarder)
	nf.request = *r
	nf.respWriter = w

	return nf
}

func (f *Forwarder) Do() {
	r := f.request
	respWriter := f.respWriter

	var err error
	var proxyResp *http.Response
	proxyClient := new(http.Client)
	proxyRequest := new(http.Request)

	if err = r.ParseForm(); err != nil {
		Log("failed to parse request")
		return
	}

	appId := r.Form.Get("appId");
	if (appId == "") {
		Log("invalid app id: " + appId)
		return
	}

	if app, ok := Config.Apps[appId]; ok {
		f.app = app
	}else {
		Log("app does not exists: " + appId)
		return;
	}

	serviceUrl := f.app.ServiceUrl + "?" + r.URL.RawQuery

	if proxyRequest.URL, err = url.Parse(serviceUrl); err != nil {
		Log("invalid service url: " + f.app.ServiceUrl)
		return
	}

	requestBody, _ := ioutil.ReadAll(r.Body)

	proxyRequest.Method = r.Method
	proxyRequest.Header = r.Header
	proxyRequest.ContentLength = int64(len(requestBody));
	proxyRequest.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

	Log("wechat request header: ", r.Header)
	Log("wechat request body: ", string(requestBody))

	for _, v := range r.Cookies() {
		proxyRequest.AddCookie(v)
	}

	Log("proxy to app: " + proxyRequest.URL.String())
	Log("proxy to app header: ", proxyRequest.Header)
	Log("proxy to app body: ", string(requestBody))

	if proxyResp, err = proxyClient.Do(proxyRequest); err != nil {
		Log("proxy request error: ", err)
		return
	}

	for k, v := range proxyResp.Header {
		for _, vv := range v {
			respWriter.Header().Add(k, vv)
		}
	}

	resp, _ := ioutil.ReadAll(proxyResp.Body)
	Log("app anwser: " + string(resp))
	respWriter.Write(resp)
}

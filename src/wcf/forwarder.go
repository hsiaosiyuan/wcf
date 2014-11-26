package wcf

import (
	"net/http"
	"io/ioutil"
	"net/url"
	"bytes"
	"log"
)

type Forwarder struct {
	integrationMode bool
	request         http.Request
	requestBody     []byte
	respWriter      http.ResponseWriter
	matched         bool
}

func NewForwarder(w http.ResponseWriter, r *http.Request) *Forwarder {
	nf := new(Forwarder)
	nf.request = *r
	nf.respWriter = w

	return nf
}

func (f *Forwarder) do(app wcfApp) {
	log.SetPrefix(app.AppId + ": ")

	r := f.request
	respWriter := f.respWriter

	var err error
	var proxyResp *http.Response

	proxyClient := new(http.Client)
	proxyRequest := new(http.Request)

	serviceUrl := app.ServiceUrl + "?" + r.URL.RawQuery

	if proxyRequest.URL, err = url.Parse(serviceUrl); err != nil {
		Log("invalid service url: " + serviceUrl)
		return
	}

	proxyRequest.Method = r.Method
	proxyRequest.Header = r.Header
	proxyRequest.ContentLength = int64(len(f.requestBody));
	proxyRequest.Body = ioutil.NopCloser(bytes.NewBuffer(f.requestBody))

	Log("wechat request header: ", r.Header)
	Log("wechat request string: ", r.URL.RawQuery)
	Log("wechat request body: ", string(f.requestBody))

	if proxyResp, err = proxyClient.Do(proxyRequest); err != nil {
		Log("proxy request error: ", err)
		return
	}

	for k, v := range proxyResp.Header {
		for _, vv := range v {
			respWriter.Header().Set(k, vv)
		}
	}

	resp, _ := ioutil.ReadAll(proxyResp.Body)
	respLen := len(resp)
	Log("app anwser: "+string(resp), " [ length:", respLen, "]")

	if respLen == 0 {
		Log("app anwser length is 0, skipped")
		return
	}

	respWriter.Write(resp)
	f.matched = true
}

func (f *Forwarder) do4Integration() {
	r := f.request

	var err error

	if err = r.ParseForm(); err != nil {
		Log("failed to parse request")
		return
	}

	appId := r.Form.Get("appId");
	if app, ok := Config.Apps[appId]; ok {
		f.do(app)
	}else {
		Log("app does not exists: " + appId)
		return;
	}
}

func (f *Forwarder) Do() {
	if f.integrationMode {
		f.do4Integration()
	}else {
		for _, app := range Config.Apps {
			if f.matched {
				break
			}

			f.do(app)
		}
	}
}

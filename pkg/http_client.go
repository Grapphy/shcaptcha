// MIT License

// Copyright (c) 2022 Grapphy

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// opies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package shcaptcha

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const DefaultUserAgent string = "Mozilla/5.0 (Windows NT 10.0; rv:91.0) Gecko/20100101 Firefox/91.0"

type HttpClient struct {
	Client       *http.Client
	AccessCookie string
	UserAgent    string
	Proxy        string
}

func NewHttpClient(accessibilityCookie string) *HttpClient {
	return &HttpClient{Client: &http.Client{}, UserAgent: DefaultUserAgent, AccessCookie: accessibilityCookie}
}

func (hc *HttpClient) SetupProxy(proxy string) {
	p, _ := url.Parse(proxy)
	hc.Client.Transport = &http.Transport{Proxy: http.ProxyURL(p)}
}

func (hc *HttpClient) GetHost(uri string) (string, error) {
	up, err := url.Parse(uri)
	if err != nil {
		return "", err
	}
	return up.Host, nil
}

func (hc *HttpClient) Request(route *Route, params *url.Values, data string) (string, error) {
	uri := route.GetPath()
	bufferData := strings.NewReader(data)

	if params != nil {
		uri = uri + "?" + params.Encode()
	}

	req, err := http.NewRequest(route.Method, uri, bufferData)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", hc.UserAgent)
	req.Header.Add("Origin", "https://newassets.hcaptcha.com")
	req.Header.Add("Referer", "https://newassets.hcaptcha.com/")
	req.Header.Add("Cookie", "hc_accessibility="+hc.AccessCookie)

	if data != "" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	res, err := hc.Client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), err
}

func (hc *HttpClient) GetChallenge(sitekey string, host string) (*Checksiteconfig, error) {
	params := url.Values{}
	params.Set("host", host)
	params.Set("sitekey", sitekey)
	params.Set("sc", "1")
	params.Set("swa", "1")

	res, err := hc.Request(&Route{Method: "GET", Path: "/checksiteconfig"}, &params, "")
	if err != nil {
		return &Checksiteconfig{}, err
	}

	var siteConfig Checksiteconfig
	err = JsonToStruct(res, &siteConfig)
	if err != nil {
		return &Checksiteconfig{}, err
	}

	return &siteConfig, nil
}

func (hc *HttpClient) GetSolution(sitekey string, host string, nhash string, motionData string, cData string) (*CaptchaSolution, error) {
	data := url.Values{}
	data.Set("sitekey", sitekey)
	data.Set("v", "fd971c2")
	data.Set("host", host)
	data.Set("n", nhash)
	data.Set("motiondata", motionData)
	data.Set("hl", "en")
	data.Set("c", cData)

	params := url.Values{}
	params.Add("s", sitekey)

	res, err := hc.Request(&Route{Method: "POST", Path: "/getcaptcha"}, &params, data.Encode())
	if err != nil {
		return &CaptchaSolution{}, err
	}

	if !strings.Contains(res, "generated_pass_UUID") {
		return &CaptchaSolution{}, errors.New("failed to bypass captcha. check your cookies/rate limit or try again")
	}

	var cresponse CaptchaSolution
	err = JsonToStruct(res, &cresponse)
	if err != nil {
		return &CaptchaSolution{}, err
	}

	return &cresponse, nil
}

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

// Client that handles all functions from this package
// so you only need to check this file.
package shcaptcha

type Client struct {
	Http *HttpClient
}

// Creates a new shcaptcha client.
func NewClient(hcAccessibility string) *Client {
	return &Client{Http: NewHttpClient(hcAccessibility)}
}

// Sets a different user agent for the internal http client.
// Default is Firefox 91.0
func (c *Client) SetUserAgent(userAgent string) {
	c.Http.UserAgent = userAgent
}

// Sets a proxy for the internal http client.
func (c *Client) SetProxy(proxy string) {
	c.Http.SetupProxy(proxy)
}

// If an accessibility cookie is set, hcaptcha will return you
// the solution string. A sitekey and url is required.
// You can find the sitekey at the html source code of the
// target site.
func (c *Client) BypassCaptcha(sitekey string, url string) (string, error) {
	host, err := c.Http.GetHost(url)
	if err != nil {
		return "", err
	}

	challenge, err := c.Http.GetChallenge(sitekey, host)
	if err != nil {
		return "", err
	}

	nhash, err := ResolveHslRequest(challenge.RCaptcha.Req)
	if err != nil {
		return "", err
	}

	motionData := NewMotionData(url, c.Http.UserAgent)
	motionDataStr, err := StructToJsonString(motionData)
	if err != nil {
		return "", err
	}

	challenge.RCaptcha.Type = "hsl"
	cDataStr, err := StructToJsonString(challenge.RCaptcha)
	if err != nil {
		return "", err
	}

	solution, err := c.Http.GetSolution(sitekey, host, nhash, string(motionDataStr), string(cDataStr))
	if err != nil {
		return "", err
	}

	return solution.Result, nil
}

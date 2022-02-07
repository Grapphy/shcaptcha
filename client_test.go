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

import "testing"

func TestClient(t *testing.T) {
	var accessibility_cookie string = "Some accessibility cookie"
	var shcpt *Client = NewClient(accessibility_cookie)
	shcpt.SetUserAgent("Mozilla/5.0 (Windows NT 10.0; rv:91.0) Gecko/20100101 Firefox/91.0")

	var sitekey string = "3ceb8624-1970-4e6b-91d5-70317b70b651"
	var url string = "https://2captcha.com/demo/hcaptcha?difficulty=always-on"

	solution, err := shcpt.BypassCaptcha(sitekey, url)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if solution[:2] != "P0" {
		t.Errorf("Invalid response: %s", solution)
	}
}

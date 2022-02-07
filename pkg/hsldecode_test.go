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
	"strings"
	"testing"
)

func TestResolveHslRequest(t *testing.T) {
	sample := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9." +
		"eyJzIjoxNiwidCI6InciLCJkIjoiYUtKM2IwQytadHVlUnB2STBOMEdBK3EyaDBlbXdBSiszYj" +
		"RGUGowY2x5aE1SQnR4MzRNdFpEYWowdENmWTNMVXJRWnpXb2xhMTBwNk1OcUVTYU9MM3MrdTdI" +
		"c1B3SEFZK0ZPdGhzTDQ5dlhvS2s2YVZmMlVpdENUK0drL0czeFQvL1ZQWDBQb0NjODR0ZndwUj" +
		"Q3N3lLWDMzREppSit4V0ZKdC9BTHowTGpKL3RCWjQ0RWtpSWp6WVEzMD1tSjNvWlJDVkVSWVFI" +
		"elYrIiwibCI6Imh0dHBzOi8vbmV3YXNzZXRzLmhjYXB0Y2hhLmNvbS9jLzk0NDlmZGE4IiwiZS" +
		"I6MTY0NDA5NzE3Nn0.Oa8-18uR07zaAcWt6KELF53TdhKAsJqW5IWPcfOwGcY"

	want := "1:16:20220206174934:aKJ3b0C+ZtueRpvI0N0GA+q2h0emwAJ+3b4FPj0clyhMRBtx34" +
		"MtZDaj0tCfY3LUrQZzWola10p6MNqESaOL3s+u7HsPwHAY+FOthsL49vXoKk6aVf2UitCT+Gk/" +
		"G3xT//VPX0PoCc84tfwpR477yKX33DJiJ+xWFJt/ALz0LjJ/tBZ44EkiIjzYQ30=mJ3oZRCVER" +
		"YQHzV+::2qZ"

	got, _ := ResolveHslRequest(sample)

	sgot := strings.Split(got, ":")
	swant := strings.Split(want, ":")

	if sgot[3] != swant[3] {
		t.Errorf("Got %s, expected %s", sgot[3], swant[3])
	}

	if sgot[4] != swant[4] {
		t.Errorf("Got %s, expected %s", sgot[4], swant[4])
	}
}

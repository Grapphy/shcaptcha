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

// This module contains all stuff related to json data.
package shcaptcha

import (
	"encoding/json"
	"time"
)

type RCaptcha struct {
	Type string `json:"type"`
	Req  string `json:"req"`
}

type Checksiteconfig struct {
	Pass     bool     `json:"pass"`
	RCaptcha RCaptcha `json:"c"`
}

type CaptchaSolution struct {
	Result     string `json:"generated_pass_UUID"`
	Expiration int    `json:"expiration"`
}

type NvData struct {
	Permissions         map[int]int `json:"permissions"`
	Donottrack          string      `json:"doNotTrack"`
	Maxtouchpoints      int         `json:"maxTouchPoints"`
	Mediacapabilities   map[int]int `json:"mediaCapabilities"`
	Oscpu               string      `json:"oscpu"`
	Vendor              string      `json:"vendor"`
	Vendorsub           string      `json:"vendorSub"`
	Productsub          string      `json:"productSub"`
	Cookieenabled       bool        `json:"cookieEnabled"`
	Buildid             string      `json:"buildID"`
	Activevrdisplays    []int       `json:"activeVRDisplays"`
	Mediadevices        map[int]int `json:"mediaDevices"`
	Serviceworker       map[int]int `json:"serviceWorker"`
	Credentials         map[int]int `json:"credentials"`
	Clipboard           map[int]int `json:"clipboard"`
	Mediasession        map[int]int `json:"mediaSession"`
	Webdriver           bool        `json:"webdriver"`
	Hardwareconcurrency int         `json:"hardwareConcurrency"`
	Geolocation         map[int]int `json:"geolocation"`
	Appcodename         string      `json:"appCodeName"`
	Appname             string      `json:"appName"`
	Appversion          string      `json:"appVersion"`
	Platform            string      `json:"platform"`
	Useragent           string      `json:"userAgent"`
	Product             string      `json:"product"`
	Language            string      `json:"language"`
	Languages           []string    `json:"languages"`
	Online              bool        `json:"onLine"`
	Storage             map[int]int `json:"storage"`
	Plugins             []int       `json:"plugins"`
}

type ScData struct {
	AvailWidth       int    `json:"availWidth"`
	AvailHeight      int    `json:"availHeight"`
	Width            int    `json:"width"`
	Height           int    `json:"height"`
	ColorDepth       int    `json:"colorDepth"`
	PixelDepth       int    `json:"pixelDepth"`
	Top              int    `json:"top"`
	Left             int    `json:"left"`
	Availtop         int    `json:"availTop"`
	Availleft        int    `json:"availLeft"`
	MozOrientation   string `json:"mozOrientation"`
	Onmozorientation *int   `json:"onmozorientationchange"`
}

type TopLevelData struct {
	St   int64     `json:"st"`
	Sc   ScData    `json:"sc"`
	Nv   NvData    `json:"nv"`
	Dr   string    `json:"dr"`
	Inv  bool      `json:"inv"`
	Exec bool      `json:"exec"`
	Wn   [][]int64 `json:"wn"`
	Wnmp int       `json:"wn-mp"`
	Xy   [][]int64 `json:"xy"`
	Xymp int       `json:"xy-mp"`
	Mm   [][]int64 `json:"mm"`
	Mmmp int       `json:"mm-mp"`
}

type PrevData struct {
	Escaped  bool `json:"escaped"`
	Passed   bool `json:"passed"`
	Expiredc bool `json:"expiredChallenge"`
	Expiredr bool `json:"expiredResponse"`
}

type MotionData struct {
	St         int64        `json:"st"`
	Mm         [][]int64    `json:"mm"`
	Mmmp       float32      `json:"mm-mp"`
	Md         [][]int64    `json:"md"`
	Mdmp       int          `json:"md-mp"`
	Mu         [][]int64    `json:"mu"`
	Mump       int          `json:"mu-mp"`
	V          int          `json:"v"`
	TopLevel   TopLevelData `json:"topLevel"`
	Session    []int        `json:"session"`
	WidgetList []string     `json:"widgetList"`
	WidgetId   string       `json:"widgetId"`
	Href       string       `json:"href"`
	Prev       PrevData     `json:"prev"`
}

// This function generates tracking data. It is set by
// default to use data given by a Firefox User-Agent on Windows 10
// but can be changed according to your needs.
func NewMotionData(uri string, userAgent string) *MotionData {
	return &MotionData{
		St: time.Now().Unix() * 1000,
		Mm: [][]int64{{41, 26, time.Now().Unix() * 1000},
			{40, 26, time.Now().Unix() * 1000},
			{40, 27, time.Now().Unix() * 1000},
			{39, 28, time.Now().Unix() * 1000}},
		Mmmp: 16.25,
		Md:   [][]int64{{40, 28, time.Now().Unix() * 1000}},
		Mdmp: 0,
		Mu:   [][]int64{{39, 28, time.Now().Unix() * 1000}},
		Mump: 0,
		V:    1,
		TopLevel: TopLevelData{
			St: time.Now().Unix() * 1000,
			Sc: ScData{
				AvailWidth:       1920,
				AvailHeight:      1040,
				Width:            1920,
				Height:           1080,
				ColorDepth:       24,
				PixelDepth:       24,
				Top:              0,
				Left:             0,
				Availtop:         0,
				Availleft:        0,
				MozOrientation:   "landscape-primary",
				Onmozorientation: nil,
			},
			Nv: NvData{
				Permissions:         make(map[int]int),
				Donottrack:          "unspecified",
				Maxtouchpoints:      0,
				Mediacapabilities:   make(map[int]int),
				Oscpu:               "Windows NT 10.0; Win64; x64",
				Vendor:              "",
				Vendorsub:           "",
				Productsub:          "20100101",
				Cookieenabled:       true,
				Buildid:             "20181001000000",
				Activevrdisplays:    []int{},
				Mediadevices:        make(map[int]int),
				Serviceworker:       make(map[int]int),
				Credentials:         make(map[int]int),
				Clipboard:           make(map[int]int),
				Mediasession:        make(map[int]int),
				Webdriver:           false,
				Hardwareconcurrency: 8,
				Geolocation:         make(map[int]int),
				Appcodename:         "Mozilla",
				Appname:             "Netscape",
				Appversion:          "5.0 (Windows)",
				Platform:            "Win32",
				Useragent:           userAgent,
				Product:             "Gecko",
				Language:            "en-US",
				Languages:           []string{"en-US", "en"},
				Online:              true,
				Storage:             make(map[int]int),
				Plugins:             []int{},
			},
			Dr:   "",
			Inv:  false,
			Exec: false,
			Wn:   [][]int64{{1920, 955, 1, time.Now().Unix() * 1000}},
			Wnmp: 0,
			Xy:   [][]int64{{0, 0, 1, time.Now().Unix() * 1000}},
			Xymp: 0,
			Mm: [][]int64{{735, 507, time.Now().Unix() * 1000},
				{718, 520, time.Now().Unix() * 1000},
				{700, 532, time.Now().Unix() * 1000},
				{697, 535, time.Now().Unix() * 1000},
				{695, 536, time.Now().Unix() * 1000},
				{691, 539, time.Now().Unix() * 1000}},
			Mmmp: 16},
		Session:    []int{},
		WidgetList: []string{"069inmsnzgz"},
		WidgetId:   "069inmsnzgz",
		Href:       uri,
		Prev: PrevData{
			Escaped:  false,
			Passed:   false,
			Expiredc: false,
			Expiredr: false,
		},
	}
}

func JsonToStruct(response string, sjson interface{}) error {
	return json.Unmarshal([]byte(response), sjson)
}

func StructToJsonString(sjson interface{}) ([]byte, error) {
	return json.Marshal(sjson)
}

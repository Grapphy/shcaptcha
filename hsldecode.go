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

// Module provides a way to decode captcha challenge into
// a hash that validates the request. Functions are a
// reimplementation of newassets.hcaptcha.com/c/165f682e/hsl.js
package shcaptcha

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

const lettersPool string = "0123456789/:abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type HslPayload struct {
	SPayload int    `json:"s"`
	DPayload string `json:"d"`
}

// Checks if literal 1 is in a given array.
// If true, returns the value at index 1 else -1
func oneInArray(nArray []int) int {
	for _, i := range nArray {
		if i == 1 {
			return nArray[1]
		}
	}
	return -1
}

// Increments values from a give array of zeroes.
// Max number is length of lettersPool.
// Once it reaches max, it is set to zero.
func incrementZero(zeroArray []int) bool {
	for i := len(zeroArray) - 1; i >= 0; i-- {
		if zeroArray[i] < len(lettersPool)-1 {
			zeroArray[i]++
			return true
		}
		zeroArray[i] = 0
	}
	return false
}

// Returns a string using characters from lettersPool.
// Uses an array containing the indices to use.
func number_to_string(indices []int) string {
	var result string
	for _, i := range indices {
		result = result + string(lettersPool[i])
	}
	return result
}

// Checks if ustring matches a valid hsl hash
func is_complete(spayload int, ustring string) bool {
	hasher := sha1.New()
	hasher.Write([]byte(ustring))
	ustringSha1 := hasher.Sum(nil)
	byteData := []int{}

	for value := 0; value < 160; value++ {
		byteresult := int(ustringSha1[value/8]) >> (value % 8) & 1
		byteData = append(byteData, byteresult)
	}

	chunk := byteData[:spayload]
	if (chunk[0] == 0) &&
		(oneInArray(chunk) >= spayload-1) ||
		(oneInArray(chunk) == -1) {
		return true
	}

	return false
}

// Tries multiple random values until a valid
// signature is found for the hsl hash.
func unhash(spayload int, dpayload string) string {
	for it := 0; it < 25; it++ {
		zeroes := []int{}
		for i := 0; i < it; i++ {
			zeroes = append(zeroes, 0)
		}
		for incrementZero(zeroes) {
			ustring := dpayload + "::" + number_to_string(zeroes)
			if is_complete(spayload, ustring) {
				return number_to_string(zeroes)
			}
		}
	}
	return ""
}

// Transform a base64 captcha request into a valid signed hsl string
func ResolveHslRequest(hslRequest string) (string, error) {
	b64payload := strings.Split(hslRequest, ".")[1]

	decodedPayload, err := base64.RawStdEncoding.DecodeString(b64payload)
	if err != nil {
		return "", err
	}

	var payload HslPayload
	json.Unmarshal(decodedPayload, &payload)
	sign := unhash(payload.SPayload, payload.DPayload)

	hsl := strings.Join([]string{
		"1",
		strconv.Itoa(payload.SPayload),
		time.Now().Format("20060102150405"),
		payload.DPayload,
		"",
		sign,
	}, ":")

	return hsl, nil
}

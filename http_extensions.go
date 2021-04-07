/**
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package attest

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
)

const defaultURL = "http://example.com"

// NewRecorder retreives an httptest.ResponseRecorder and http.Request pointer
// pair. This function is variadic. That is, it accepts a varying number of
// arguments.
// If no arguments are passed, it uses the method "GET" and the default URL,
// with an empty body. This works for testing the index path.
// The first-priority parameter is the URL. It can begin  with an actual URL,
// or a literal "/" and the default URL will be used with the given path.
// If two paramters are passed, they should be in the order (method, URL). An
// empty body will be used.
// If three paramters are passed, they are forwarded to the httptest.NewRequest
// function, with the following modifications:
//  - The default URL is prepended to a URL which starts with "/"
//  - The body is converted from a string with bytes.NewBufferString.
func (t *Test) NewRecorder(params ...string) (*httptest.ResponseRecorder, *http.Request) {
	switch len(params) {
	case 0:
		return t.NewRecorder("GET", defaultURL+"/")
	case 1:
		return t.NewRecorder("GET", params[0])
	case 2:
		url := params[1]
		if url[0] == '/' {
			url = defaultURL + url
		}
		return httptest.NewRecorder(), httptest.NewRequest(params[0], url, nil)
	case 3:
		url := params[1]
		if url[0] == '/' {
			url = defaultURL + url
		}
		return httptest.NewRecorder(), httptest.NewRequest(
			params[0],
			url,
			bytes.NewBufferString(params[2]),
		)
	}
	t.Fatalf(
		"Unable to build ResponseRecorder/Request pair for parameters: %v",
		params,
	)
	return nil, nil
}

// ResponseOK passes the test if the status code of the given response is less
// than 400
func (t *Test) ResponseOK(response *http.Response, msgAndFmt ...interface{}) {
	var message string
	switch len(msgAndFmt) {
	case 0:
		message = ""
	case 1:
		message = msgAndFmt[0].(string)
	default:
		message = fmt.Sprintf(msgAndFmt[0].(string), msgAndFmt[1:len(msgAndFmt)-1]...)
	}
	if response.StatusCode > 400 {
		t.Errorf(
			"Got status %d: %s.\n%s",
			response.StatusCode,
			response.Status,
			message,
		)
	}
}

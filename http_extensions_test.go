package attest

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_NewRecorder(t *testing.T) {
	test := Test{t}
	rec, req := test.NewRecorder()
	test.Equals("GET", req.Method)
	test.Equals("/", req.URL.Path)
	test.Equals("example.com", req.URL.Host)
	test.Equals("http", req.URL.Scheme)
	test.TypeIs("http.noBody", req.Body)
	test.TypeIs("*httptest.ResponseRecorder", rec)
	rec, req = test.NewRecorder("http://different-url.and/path")
	test.Equals("GET", req.Method)
	test.Equals("/path", req.URL.Path)
	test.Equals("different-url.and", req.URL.Host)
	test.Equals("http", req.URL.Scheme)
	test.TypeIs("http.noBody", req.Body)
	test.TypeIs("*httptest.ResponseRecorder", rec)
	rec, req = test.NewRecorder(
		"GET",
		"http://a-url.and/path",
		"A string to go in the body",
	)
	test.Equals("GET", req.Method)
	test.Equals("/path", req.URL.Path)
	test.Equals("a-url.and", req.URL.Host)
	test.Equals("http", req.URL.Scheme)
	readBack, err := ioutil.ReadAll(req.Body)
	req.Body.Close()
	test.Handle(err)
	test.Equals("A string to go in the body", string(readBack))
	test.TypeIs("*httptest.ResponseRecorder", rec)
	rec, req = test.NewRecorder("/just/the/path")
	test.Equals("GET", req.Method)
	test.Equals("/just/the/path", req.URL.Path)
	test.Equals("example.com", req.URL.Host)
	test.Equals("http", req.URL.Scheme)
	test.TypeIs("http.noBody", req.Body)
	test.TypeIs("*httptest.ResponseRecorder", rec)
}
func Test_ResponseOK(t *testing.T) {
	test := Test{t}
	rec, req := test.NewRecorder()
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}(rec, req)
	res := rec.Result()
	test.ResponseOK(res)
}

package gocookie_string_reader

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

var readCookiesTests = []struct {
	Header http.Header
}{
	{
		http.Header{"Cookie": {"Cookie-1=v$1", "c2=v2"}},
	},
	{
		http.Header{"Cookie": {"Cookie-1=v$1", "c2=v2"}},
	},
	{
		http.Header{"Cookie": {"Cookie-1=v$1; c2=v2"}},
	},
	{
		http.Header{"Cookie": {"Cookie-1=v$1; c2=v2"}},
	},
	{
		http.Header{"Cookie": {`Cookie-1="v$1"; c2="v2"`}},
	},
}

func TestReadCookies(t *testing.T) {
	for i, tt := range readCookiesTests {
		for n := 0; n < 2; n++ { // to verify readCookies doesn't mutate its input
			c := ParseToCookie(tt.Header.Get("Cookie"))
			rawRequest := fmt.Sprintf("GET / HTTP/1.0\r\nCookie: %s\r\n\r\n", tt.Header.Get("Cookie"))

			req, _ := http.ReadRequest(bufio.NewReader(strings.NewReader(rawRequest)))

			if !reflect.DeepEqual(c, req.Cookies()) {
				t.Errorf("#%d readCookies:\nhave: %s\nwant: %s\n", i, toJSON(c), toJSON(req.Cookies()))
				continue
			}
		}
	}
}

func toJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("%#v", v)
	}
	return string(b)
}

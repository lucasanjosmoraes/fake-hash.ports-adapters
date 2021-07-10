package nethttpadapter

import (
	"io/ioutil"
	"net/http"
)

type Request struct {
	NetHTTPRequest *http.Request
}

func (r Request) BodyBytes() []byte {
	data, err := ioutil.ReadAll(r.NetHTTPRequest.Body)
	if err != nil {
		return make([]byte, 0)
	}

	defer func() {
		_ = r.NetHTTPRequest.Body.Close()
	}()

	return data
}

func (r Request) Url() string {
	return r.NetHTTPRequest.URL.String()
}

func (r Request) Header(name string) string {
	return r.NetHTTPRequest.Header.Get(name)
}

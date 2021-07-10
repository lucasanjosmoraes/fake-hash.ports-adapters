package toolkit

type MockHTTPRequest struct {
	Body    []byte
	URL     string
	Headers map[string]string
}

func (m MockHTTPRequest) BodyBytes() []byte {
	return m.Body
}

func (m MockHTTPRequest) Url() string {
	return m.URL
}

func (m MockHTTPRequest) Header(name string) string {
	return m.Headers[name]
}

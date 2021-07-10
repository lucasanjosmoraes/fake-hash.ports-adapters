package nethttpadapter

import "net/http"

type Response struct {
	ResponseWriter http.ResponseWriter
}

func (r Response) Write(statusCode int, body []byte) error {
	return writeCommon(r.ResponseWriter, statusCode, body)
}

func (r Response) WriteCreated(body []byte) error {
	return writeCommon(r.ResponseWriter, http.StatusCreated, body)
}

func (r Response) WriteBadRequest(body []byte) error {
	return writeCommon(r.ResponseWriter, http.StatusBadRequest, body)
}

func (r Response) WriteInternalError(body []byte) error {
	return writeCommon(r.ResponseWriter, http.StatusInternalServerError, body)
}

func (r Response) JsonResponse() {
	r.ResponseWriter.Header().Set("Content-Type", "application/json")
}

func writeCommon(writer http.ResponseWriter, status int, body []byte) error {
	writer.WriteHeader(status)
	_, err := writer.Write(body)

	return err
}

package middleware

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/context"
)

func CopyBody(rw http.ResponseWriter, r *http.Request) {
	b := bytes.NewBuffer([]byte{})
	reader := io.TeeReader(r.Body, b)
	ioutil.ReadAll(reader)
	r.Body = ioutil.NopCloser(b)

	context.Set(r, "body", b.Bytes())
}

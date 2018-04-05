package status

import "net/http"

func status(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("alive"))
}

package headers

import (
	"net/http"

	"github.com/stormentt/zpass-server/db"
)

func GetDevice(r *http.Request) (*db.Device, bool) {
	sel := r.Header.Get("Device-Selector")
}

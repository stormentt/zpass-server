package request

import (
	"net/http"

	"github.com/stormentt/zpass-lib/canister"
	"github.com/stormentt/zpass-server/db"
)

type Request struct {
	can           canister.Canister
	Authenticated bool
	Device        *db.Device
	signature     []byte
}

func Parse(r *http.Request) {

}

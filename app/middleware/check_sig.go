package middleware

import (
	"encoding/hex"
	"net/http"

	"github.com/gorilla/context"
	log "github.com/sirupsen/logrus"
	"github.com/stormentt/zpass-lib/zcrypto"
	"github.com/stormentt/zpass-server/db"
)

func CheckSig(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	dev := context.Get(r, "device").(*db.Device)
	if dev == nil {
		log.Warning("authorization attempted without device")
		return
	}

	body := context.Get(r, "body").([]byte)

	signature := r.Header.Get("Content-Signature")
	decoded := make([]byte, zcrypto.AuthSigSize)
	n, err := hex.Decode(decoded, []byte(signature))
	if err != nil || n != zcrypto.AuthSigSize {
		log.Warning("authorization attempted with bad-length Content-Signature")
		return
	}

	validSig := dev.CheckSig(body, []byte(signature))
	if !validSig {
		log.Warning("authorization attempted with invlaid Content-Signature")
		return
	}

	next(rw, r)
}

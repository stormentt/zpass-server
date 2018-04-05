package middleware

import (
	"net/http"

	"github.com/gorilla/context"
	log "github.com/sirupsen/logrus"
	"github.com/stormentt/zpass-server/db"
)

func FindDevice(rw http.ResponseWriter, r *http.Request) {
	sel := r.Header.Get("Device-Selector")
	device, err := db.FindDevice(sel)
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err,
			"selector": sel,
		}).Error("unable to find device")
		return
	}

	if device != nil {
		context.Set(r, "device", device)
	}
}

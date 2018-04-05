package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stormentt/zpass-server/app/controllers/status"
	"github.com/stormentt/zpass-server/app/controllers/users"
	"github.com/stormentt/zpass-server/app/routing"
)

func Serve() {
	host := viper.GetString("host")
	port := viper.GetInt("port")
	address := fmt.Sprintf("%s:%d", host, port)

	cLog := log.WithFields(log.Fields{
		"host":    host,
		"port":    port,
		"address": address,
	})

	cLog.Info("starting server")

	r := newRouter()
	err := http.ListenAndServe(address, r)
	if err != nil {

		cLog.WithFields(log.Fields{
			"error": err,
		}).Fatal("unable to start server")
	}
}

func newRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	userSubrouter := r.PathPrefix("/users").Subrouter()
	routing.RegisterRoutes(userSubrouter, users.Routes)

	statusSubrouter := r.PathPrefix("/status").Subrouter()
	routing.RegisterRoutes(statusSubrouter, status.Routes)

	return r
}

package users

import "github.com/stormentt/zpass-server/app/routing"

var Routes = []routing.Route{
	routing.Route{
		"UsersStore",
		"POST",
		"/",
		store,
	},
}

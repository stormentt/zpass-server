package status

import "github.com/stormentt/zpass-server/app/routing"

var Routes = []routing.Route{
	routing.Route{
		"CheckStatus",
		"GET",
		"/",
		status,
	},
}

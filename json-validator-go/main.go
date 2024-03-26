package main

//https://stackoverflow.com/questions/76448408/cors-policy-response-to-preflight-request-doesnt-pass-access-control-check-no
//https://stackoverflow.com/questions/76448408/cors-policy-response-to-preflight-request-doesnt-pass-access-control-check-no
//https://en.wikipedia.org/wiki/Cross-origin_resource_sharing

import (
	"json-validator/api"
)

func main() {
	apiServer := api.ApiServer()
	apiServer.Run("localhost:8080")
}

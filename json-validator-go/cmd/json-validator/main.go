package main

import "json-validator/internal/routes"

//https://stackoverflow.com/questions/76448408/cors-policy-response-to-preflight-request-doesnt-pass-access-control-check-no
//https://stackoverflow.com/questions/76448408/cors-policy-response-to-preflight-request-doesnt-pass-access-control-check-no
//https://en.wikipedia.org/wiki/Cross-origin_resource_sharing

func main() {
	router := routes.NewRouter()
	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}

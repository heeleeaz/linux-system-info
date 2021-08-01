package main

import (
	"net/http"
)

func main() {
	exposeSystemInformationEndpoints()
	http.ListenAndServe(":8080", nil)
}

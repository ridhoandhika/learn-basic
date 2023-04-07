package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", authcontroller.index)
}

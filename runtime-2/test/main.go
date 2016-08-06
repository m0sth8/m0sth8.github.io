package main

import (
	"net/http"
)

func tmp() {
	http.DefaultClient.Do(nil) // req, err
}

package main

import "net/http"

func SessionsLoad(next http.Handler) http.Handler {
	return sessions.LoadAndSave(next)
}
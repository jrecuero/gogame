package server

import (
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pomg"))
}

func Server() {
	http.HandleFunc("/ping", ping)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

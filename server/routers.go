package server

import "net/http"

func TasksRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		get(w, r)
		return
	case "POST":
		post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}

func TaskRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getTask(w, r)
		return
	case "PUT":
		updateTask(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}

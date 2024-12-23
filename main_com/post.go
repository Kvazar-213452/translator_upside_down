package main_com

import "net/http"

func Post_test(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		w.Header().Set("Content-Type", "application/json")
		w.Write(nil)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

package handle

import "net/http"

func ApiDemo(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Get data success api"))
}

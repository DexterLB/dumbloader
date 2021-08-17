package static

import "net/http"

func Serve(f []byte, rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Write(f)
}

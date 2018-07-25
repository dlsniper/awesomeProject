package subpkgdemo

type MyHandler struct{}

func (h MyHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	if h.isChrome(r) {
		w.Write
	}

	w.Write([]byte("Hello from indexHandler"))
}

func (h MyHandler) isChrome(r *http.Request) bool {
	return strings.ToLower(r.UserAgent()) == "chrome"
}

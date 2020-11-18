package handlers

import "net/http"

type hello struct {

}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	log.Println("Hello World!")
	d, _ := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "error", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hello %s\n", d)
}
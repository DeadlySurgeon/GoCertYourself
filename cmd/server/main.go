package main

import (
	"log"
	"net/http"
)

const (
	prefix = "../../resources/"
	crt    = prefix + "localhost.crt"
	key    = prefix + "localhost.key"
)

func helloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
	// fmt.Fprintf(w, "This is an example server.\n")
	// io.WriteString(w, "This is an example server.\n")
}

func main() {
	http.HandleFunc("/hello", helloServer)
	err := http.ListenAndServeTLS(":443", crt, key, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

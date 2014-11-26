package wcf

import (
	"net/http"
	"fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
	f := NewForwarder(w, r);
	f.Do()
}

func RunServer(port string) {
	listenAddr := ":" + port

	fmt.Println("wcf running on: " + listenAddr)
	http.HandleFunc("/", handler)
	http.ListenAndServe(listenAddr, nil)
}

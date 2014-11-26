package wcf

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func RunServer(port string, integrationMode bool) {
	listenAddr := ":" + port
	fmt.Println("wcf running on: " + listenAddr)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			f := NewForwarder(w, r);
			f.integrationMode = integrationMode
			f.requestBody, _ = ioutil.ReadAll(r.Body)

			f.Do()
		})

	http.ListenAndServe(listenAddr, nil)
}

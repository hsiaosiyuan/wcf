package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	f := NewForwarder(w, r);
	f.Do()
}

func main() {
	config := `
		{
			"Port" : 8888,
			"Apps": {
				"app1" : {
					"ServiceUrl" : "url"
				}
			}
		}
	`
	if Config.Init(config) == false {
		fmt.Println("invalid config strings")
	}

	listenAddr := ":" + string(Config.Port)

	http.HandleFunc("/", handler)
	http.ListenAndServe(listenAddr, nil)
}

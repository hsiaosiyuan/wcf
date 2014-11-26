package main

import (
	"fmt"
	"wcf"
	"flag"
)

func main() {
	config := `
		{
			"Debug" : true,
			"LogFile" : "debug.log",
			"Port" : "80",
			"Apps": {
				"txmch" : {
					"ServiceUrl" : "http://shop.plmt-soft.com/index.php/mobile/connection/weChat"
				},
				"UTouch" : {
					"ServiceUrl" : "http://180.97.81.101/youtouchx/index.php/weichat/connection/handle"
				}
			}
		}
	`
	if wcf.Config.Init(config) == false {
		fmt.Println("invalid config strings")
		return
	}

	var integrationMode bool
	flag.BoolVar(&integrationMode, "i", false, "integration mode")

	wcf.RunServer(wcf.Config.Port, integrationMode)
}

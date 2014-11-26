package main

import (
	"fmt"
	"wcf"
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
				}
			}
		}
	`
	if wcf.Config.Init(config) == false {
		fmt.Println("invalid config strings")
		return
	}

	wcf.RunServer(wcf.Config.Port)
}

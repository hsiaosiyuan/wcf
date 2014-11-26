package wcf

import (
	"testing"
)

func TestConfigInit(t *testing.T) {
	json := `
		{
			"Port" : "8888",
			"Apps": {
				"app1" : {
					"ServiceUrl" : "url"
				}
			}
		}
	`

	if Config.Init(json) == false {
		t.Fatal("Invalid json")
	}

	if len(Config.Apps) <= 0 {
		t.Fatal("Failed to init")
	}

	if Config.Apps["app1"].ServiceUrl != "url" {
		t.Fatal("Failed to parse apps url")
	}

	if Config.Apps["app1"].AppId != "app1" {
		t.Fatal("Failed to parse apps id")
	}

	if Config.Port != 8888 {
		t.Fatal("Failed to parse port")
	}
}

package wcf

import (
	JSON "encoding/json"
)

type config struct {
	isInitialized bool
	json          string
	Debug         bool 	`json:"Debug"`
	LogFile       string `json:"LogFile"`
	Port          string `json:"Port"`
	Apps map[string]wcfApp `json:"Apps"`
}

type wcfApp struct {
	AppId           string `json:"AppId"`
	ServiceUrl      string `json:"ServiceUrl"`
}

var Config = new(config)

func (c config) Init(json string) bool {
	if Config.isInitialized {
		return true
	}

	if json == "" {
		return false
	}

	Config.json = json

	if err := JSON.Unmarshal([]byte(json), Config); err != nil {
		return false
	}

	for appId, app := range Config.Apps {
		app.AppId = appId;
		Config.Apps[appId] = app
	}

	return true
}

package wcf

import (
	"log"
	"fmt"
	"reflect"
	"os"
)

func Log(v ...interface{}) {
	if !Config.Debug {
		return
	}

	if (Config.LogFile == "") {
		fmt.Println("missing log file")
		return
	}

	f, err := os.OpenFile(Config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		return
	}

	defer f.Close()

	log.SetOutput(f)

	args := make([]reflect.Value, len(v))
	for i := 0; i < len(v); i++ {
		args[i] = reflect.ValueOf(v[i])
	}

	l := reflect.ValueOf(log.Println)
	l.Call(args)
}

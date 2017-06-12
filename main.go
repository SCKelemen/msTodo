package main

import (
	"github.com/SCKelemen/msTodo/web"
	"github.com/SCKelemen/msTodo/logging"
	"os"
	"io/ioutil"
)

func main() {
	logging.Initialize(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	logging.Trace.Log("Loggers initialized")

	web.ApiRouter().Run("127.0.0.1:6969")
	logging.Trace.Log("Listening on 127.0.0.1:6969")
}

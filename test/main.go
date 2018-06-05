package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/intoyun/gologger/log"
	"github.com/intoyun/gologger/errors"
)

var fluentEnable = true
var logFile      = "gologger_test.log"

func main() {
	fluentHost := "192.168.1.100"
	fluentPort := 24225
	appName    := "intoyun-gologger"
	if fluentEnable {
		log.InitFluent(fluentHost, fluentPort, appName)
	}

	file, err := exec.LookPath(os.Args[0])
	fmt.Println("==> exec bin file:", file)

	logDir := filepath.Join(filepath.Dir(file), "logs")
	fmt.Println("==> logDir:", logDir)

	os.MkdirAll(logDir, os.ModePerm)
	w, err := log.NewRollingFile(logDir+"/"+logFile, log.DailyRolling)
	if err != nil {
		fmt.Println("==> new rolling file err: ", err)
		log.PanicErrorf(err, "==> Error: open log file %s failed", logFile)
	} else {
		log.StdLog = log.New(w, "")
	}

	// Set lowwer level, otherwise CAN'T generate log file.
	log.SetLevel(log.LevelInfo)
	log.Infof("==> gologger start...")

	logLevel := "DEBUG"
	log.SetLevelString(logLevel)
	log.Debugf("==> this is a debug log.")
	log.Infof("==> yep! this is a info log.")

	errTest := errors.New("error log.")
	log.Error(errTest, "==> hm! this is an error log.")
}

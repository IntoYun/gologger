package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/intoyun/gologger/log"
	"github.com/intoyun/gologger/errors"
)

var fluentEnable = false
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

	log.SetLevelString("DEBUG")
	log.Print("\n")
	log.Println("==> set level to DEBUG.")
	log.Debug("==> this is a debug log.")
	log.Info("==> this is a info log.")
	log.Warn("==> this is a warn log.")
	log.Error("==> this is a error log with stack info automatically.")

	log.SetLevelString("INFO")
	log.Print("\n")
	log.Println("==> set level to INFO.")
	log.Debugf("==> this will be ignored.")
	log.Infof("==> this is a Info %s log.", "formatted")
	log.Warnf("==> this is a Warn %s log.", "formatted")
	log.Errorf("==> this is a Error %s log with stack info automatically.", "formatted")

	gerr := errors.New("error testing")

	log.SetLevelString("WARN")
	log.Printf("%s", "\n")
	log.Println("==> set level to WARN.")
	log.DebugError(gerr, "==> this will be ignored.")
	log.InfoError(gerr, "==> this will be ignored.")
	log.WarnError(gerr, "==> this is a warn log with stack.")
	log.ErrorError(gerr, "==> this is a Error log with stack info by purpose.")

	log.SetLevelString("ERROR")
	log.Printf("%s", "\n")
	log.Println("==> set level to ERROR.")
	log.DebugErrorf(gerr, "==> this will be ignored.")
	log.InfoErrorf(gerr, "==> this will be ignored.")
	log.WarnErrorf(gerr, "==> this will be ignored.")
	log.ErrorErrorf(gerr, "==> this is a Error %s log with stack info by purpose.", "formatted")

	log.SetLevelString("PANIC")
	log.Printf("%s", "\n")
	log.Println("==> set level to PANIC.")
	log.ErrorError(gerr, "==> Error message will ignore the log level.")
	log.PanicErrorf(gerr, "==> this is a Panic %s log with stack. will exit!!!", "formatted")
	log.Println("==> this won't be executed.")
}

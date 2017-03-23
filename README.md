# gologger

use fluent to colllect logs. log is based on CodisLabs log package.

## usage

you must use it with fluentd, and make sure that connection to fluentd well.

````golang 
func main() {
	logLevel := "INFO" 
	logFile := "test.log" 
    fluentHost := 192.168.1.100
    fluentPort:=24225
    appname:=intoyun-kfkworkers
    fluent.InitSetting(fluentHost, fluentPort, appname)
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.PanicErrorf(err, "Error in exec.LookPath")
	}
	logdir := filepath.Join(filepath.Dir(file), "log")
	os.MkdirAll(logdir, os.ModePerm)
	w, err := log.NewRollingFile(logdir+"/"+logFile, log.DailyRolling)
	if err != nil {
		log.PanicErrorf(err, "open log file %s failed", logFile)
	} else {
		log.StdLog = log.New(w, "")
	}
	log.SetLevelString(logLevel)
	// test
	log.Debugf("debug log.")
	log.Infof("info log.")
	errTest := errors.New("error log.")
	log.Error(errTest, "error log.")
}
````

package fluent

import (
	"fmt"
	"time"

	"github.com/fluent/fluent-logger-golang/fluent"
)

var (
	f         *fluent.Fluent
	FluentTag string
)

func Post(message interface{}) {
	f.Post(FluentTag, message)
}

func PostWithTime(tm time.Time, message interface{}) {
	f.PostWithTime(FluentTag, tm, message)
}

func EncodeAndPostData(tm time.Time, data interface{}) {
	f.EncodeAndPostData(FluentTag, tm, data)
}

func InitSetting(fluentHost string, fluentPort int, appname string) {
	var err error
	f, err = fluent.New(fluent.Config{FluentHost: fluentHost, FluentPort: fluentPort})
	if err != nil {
		fmt.Println(err)
	}
	var data = map[string]string{
		"txt": "fluent started.",
	}
	f.Post(appname, data)
}

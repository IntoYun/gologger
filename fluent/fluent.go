package fluent

import (
	"fmt"
	"time"

	"github.com/fluent/fluent-logger-golang/fluent"
)

var (
	f *fluent.Fluent
	FluentTag string
	Connected bool
)

func Post(message interface{}) {
	if Connected {
		f.Post(FluentTag, message)
	}
}

func PostWithTime(tm time.Time, message interface{}) {
	if Connected {
		f.PostWithTime(FluentTag, tm, message)
	}
}

func EncodeAndPostData(tm time.Time, data interface{}) {
	if Connected {
		f.EncodeAndPostData(FluentTag, tm, data)
	}
}

func New(fluentHost string, fluentPort int, appName string) bool {
	var err error
	f, err = fluent.New(fluent.Config{FluentHost: fluentHost, FluentPort: fluentPort})
	if err != nil {
		fmt.Println("==> Connect to fluent Error: ", err)
		Connected = false
	} else {
		Connected = true
		FluentTag = appName
	}

	return Connected
}

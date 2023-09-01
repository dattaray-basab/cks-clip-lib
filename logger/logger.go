package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

const LOG_LEVEL = logrus.DebugLevel

var (
	Log *logrus.Logger // share will all packages
)

func init() {
	if _, err := os.Stat("logs/application.log"); os.IsNotExist(err) {
		os.Mkdir("logs", 0755)
	}
	f, err := os.OpenFile("logs/application.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}
	Log = logrus.New()
	Log.SetFormatter(&logrus.TextFormatter{
		// ForceColors: true,
		// DisableColors: false,
		PadLevelText: true,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyMsg:   "@message",
			logrus.FieldKeyLevel: "@level",
			logrus.FieldKeyFunc:  "@caller",
		},
	})
	Log.SetReportCaller(true)
	Log.SetLevel(LOG_LEVEL)

	mw := io.MultiWriter(os.Stdout, f)
	Log.SetOutput(mw)
}

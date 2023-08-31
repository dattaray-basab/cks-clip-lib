package logger

import (
	"io"
	// logging "log"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	Log *logrus.Logger // share will all packages
)

func init() {
	// the file needs to exist prior
	if _, err := os.Stat("logs/application.log"); os.IsNotExist(err) {
		os.Mkdir("logs", 0755)
	}
	f, err := os.OpenFile("logs/application.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		// use go's logger, while we configure Logrus
		panic(err)
	}

	// configure Logrus
	Log = logrus.New()
	Log.SetFormatter(&logrus.TextFormatter{
      // ForceColors: true,
		// DisableColors: false,
      PadLevelText: true,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "@timestamp",
			logrus.FieldKeyMsg: "@message",
			logrus.FieldKeyLevel: "@level",
         logrus.FieldKeyFunc: "@caller",
		},
	})
	// Log.Formatter = &logrus.JSONFormatter{}
	Log.SetReportCaller(true)

	mw := io.MultiWriter(os.Stdout, f)
	Log.SetOutput(mw)
}

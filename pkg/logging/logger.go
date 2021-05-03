package logging

import "github.com/sirupsen/logrus"

func GetLogger(lvl string) *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		log.SetLevel(logrus.TraceLevel)
		log.Infof("failed to parse log level %s or it was empty, defaulting to debug", lvl)
	} else {
		log.SetLevel(level)
	}
	return log
}

package goutil

import (
	"fmt"
	"runtime/debug"

	"github.com/sirupsen/logrus"
)

func RecoverAndLog(log logrus.FieldLogger) {
	if r := recover(); r != nil {
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("recovered %#v", r)
		}
		log.WithField("trace", string(debug.Stack())).WithError(err).Errorln("Panic!")
	}
}

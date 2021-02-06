package main

import (
	"errors"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestDefault(t *testing.T) {
	SetConfig("", "", "", true)
	Init()

	// 写入日志
	Debug("hello, %s", "world")
	WithError(errors.New("test err")).Debugf("hello, %s", "all")
	WithField("name", "yuanzp").Info("testing")
	WithFields(logrus.Fields{"name": "test1", "value": "jalg"}).Info("test2")
}

package task

import (
	"github.com/Sirupsen/logrus"
	"k8s.io/client-go/informers"
)

// WatchTaskHelper ...
type WatchTaskHelper interface {
	Run(informers informers.SharedInformerFactory, stop <-chan struct{})
}

// DefaultWatchTaskHelper ...
type DefaultWatchTaskHelper struct {
}

var registry = make(map[string]WatchTaskHelper)

// Register add different authenticators to registry map.
func Register(name string, h WatchTaskHelper) {
	if _, dup := registry[name]; dup {
		logrus.Infof("task: %s has been registered, skip", name)
		return
	}
	registry[name] = h
	logrus.Infof("Registered task : %s", name)
}

// RunTask ...
func RunTask(informers informers.SharedInformerFactory, stop <-chan struct{}) {
	logrus.Info("tasks start ...")
	for name, taskItem := range registry {
		logrus.Infof("start watchTask: %s", name)
		go taskItem.Run(informers, stop)
	}
}

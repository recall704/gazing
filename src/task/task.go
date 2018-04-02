package task

import (
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/informers"
)

// WatchHelper ...
type WatchHelper interface {
	Run(informers informers.SharedInformerFactory, stop <-chan struct{})
}

// DefaultWatchHelper ...
type DefaultWatchHelper struct {
}

// RunWatchTasks   run all watch task
func RunWatchTasks(informers informers.SharedInformerFactory, stop <-chan struct{}) {
	for name, item := range registry {
		logrus.Infof("start watch task : %s", name)
		item.Run(informers, stop)
	}
}

var registry = make(map[string]WatchHelper)

// Register add different authenticators to registry map.
func Register(name string, w WatchHelper) {
	if _, dup := registry[name]; dup {
		logrus.Infof("watch task: %s has been registered,skip", name)
	}
	registry[name] = w
	logrus.Infof("Registered watch task: %s", name)
}

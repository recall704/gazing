package task

import (
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/informers"
)

// WatchHelper ...
type WatchHelper interface {
	Run(stop <-chan struct{})
}

// DefaultWatchHelper ...
type DefaultWatchHelper struct {
}

// RunWatchTasks   run all watch task
func RunWatchTasks(informers informers.SharedInformerFactory) {

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

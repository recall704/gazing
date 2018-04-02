package pod

import (
	"github.com/recall704/gazing/src/task"
	"k8s.io/client-go/informers"
)

// PodWatchTask ...
type PodWatchTask struct {
	task.DefaultWatchHelper
}

// Run ...
func (task *PodWatchTask) Run(informers informers.SharedInformerFactory, stop <-chan struct{}) {
	podInformer := informers.Core().V1().Pods().Informer()
	podInformer.Run(stop)
}

func init() {
	task.Register("pod", &PodWatchTask{})
}

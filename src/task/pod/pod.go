package pod

import (
	"github.com/recall704/gazing/src/task"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
)

type podWatchTask struct {
	task.DefaultWatchTaskHelper
}

// Run ...
func (podTask *podWatchTask) Run(informers informers.SharedInformerFactory, stop <-chan struct{}) {
	podInformer := informers.Core().V1().Pods().Informer()

	podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    processPodAdd,
		UpdateFunc: processPodUpdate,
		DeleteFunc: processPodDelete,
	})

	podInformer.Run(stop)
}

func init() {
	task.Register("pod", &podWatchTask{})
}

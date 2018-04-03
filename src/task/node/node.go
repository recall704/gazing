package node

import (
	"github.com/recall704/gazing/src/task"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
)

type nodeWatchTask struct {
	task.DefaultWatchTaskHelper
}

// Run ...
func (nodeTask *nodeWatchTask) Run(informers informers.SharedInformerFactory, stop <-chan struct{}) {
	nodeInformer := informers.Core().V1().Nodes().Informer()

	nodeInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    processNodeAdd,
		UpdateFunc: processNodeUpdate,
		DeleteFunc: processNodeDelete,
	})

	nodeInformer.Run(stop)
}

func init() {
	task.Register("node", &nodeWatchTask{})
}

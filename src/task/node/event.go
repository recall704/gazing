package node

import (
	"github.com/Sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
)

func processNodeAdd(obj interface{}) {
	node := obj.(*v1.Node)
	logrus.Info(node.Name)
}

func processNodeUpdate(oldObj, newObj interface{}) {
	oldNode := oldObj.(*v1.Node)
	newNode := newObj.(*v1.Node)
	logrus.Info(oldNode.Name)
	logrus.Info(newNode.Name)
}

func processNodeDelete(obj interface{}) {
	node := obj.(*v1.Node)
	logrus.Info(node.Name)
}

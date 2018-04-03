package pod

import (
	"github.com/Sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
)

func processPodAdd(obj interface{}) {
	pod := obj.(*v1.Pod)
	logrus.Infof("pod add :%s", pod.Name)
}

func processPodUpdate(oldObj, newObj interface{}) {
	oldPod := oldObj.(*v1.Pod)
	newPod := newObj.(*v1.Pod)
	logrus.Info(oldPod.Name)

	logrus.Infof("pod update %s", newPod.Name)
}

func processPodDelete(obj interface{}) {
	pod := obj.(*v1.Pod)
	logrus.Infof("pod delete :%s", pod.Name)
}

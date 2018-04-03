package main

import (
	"flag"
	"os"

	"github.com/Sirupsen/logrus"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/recall704/gazing/src/log"
	"github.com/recall704/gazing/src/task"
	_ "github.com/recall704/gazing/src/task/node"
	_ "github.com/recall704/gazing/src/task/pod"
)

var (
	kubeconfig string
)

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "config", "kubeconfig file path")
	flag.Parse()

	if kubeconfig == "" {
		logrus.Errorln("kubeconfig file must be required")
		os.Exit(0)
	}

	if _, err := os.Stat(kubeconfig); os.IsNotExist(err) {
		logrus.Errorln("kubeconfig file does not exists")
		os.Exit(0)
	}

	log.InitLogrus("DEBUG", true)

}

func main() {

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	sharedInformers := informers.NewSharedInformerFactory(clientset, 0)
	sharedInformers.WaitForCacheSync(wait.NeverStop)

	task.RunTask(sharedInformers, wait.NeverStop)

	select {}
}

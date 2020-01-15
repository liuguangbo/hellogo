package main

import (
	"k8s.io/client-go/kubernetes"
	// "path/filepath"
	//"k8s.io/apimachinery/pkg/util/intstr"
)

var ClientSet *kubernetes.Clientset

func main() {
	ClientSet = InitClientSet()
	router := InitRouter()
	router.Run(":8888")
}

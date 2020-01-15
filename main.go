package main

import (
	"k8s.io/client-go/kubernetes"
)

var ClientSet *kubernetes.Clientset

func main() {
	ClientSet = InitClientSet()
	router := InitRouter()
	router.Run(":8888")
}

package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	clientset "github.com/hidevopsio/mioclient/examples/samplecontroller/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	masterURL  string
	kubeconfig string
)

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "/Users/mac/.kube/config", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}

func main() {
	flag.Parse()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		glog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	exampleClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building example clientset: %s", err.Error())
	}

	fooList, err := exampleClient.SamplecontrollerV1alpha1().Foos("defaule").List(metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error %v", err)
		return
	}
	fmt.Println("fooList ", fooList)

}

package main

import (
	"context"
	"flag"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "C:/Users/wv3cxq/.kube/config", "location to kubeconfig file")

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("Error while reading kubeconfig, %s. Will try for incluster.", err.Error())

		config, err = rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// list all pods from kube-system namespace
	ListAllPodsFromNamespace(client, "kube-system")

	// list all pods from kube-system namespace with label tier=control-plane
	ListAllPodsFromNamespaceWithLabel(client, "kube-system", "tier=control-plane")

	// list all pods which have priority more than 0
	ListAllPodsFromNamespaceWithField(client, "kube-system", "status.phase=Running")
}

func ListAllPodsFromNamespace(client *kubernetes.Clientset, namespace string) {
	fmt.Printf("Listing all pods from %s namespace\n", namespace)
	podList, err := client.CoreV1().Pods(namespace).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for i, pod := range podList.Items {
		fmt.Printf("%d: %s, %s\n", i, pod.Name, pod.Status.QOSClass)
	}
}

func ListAllPodsFromNamespaceWithLabel(client *kubernetes.Clientset, namespace string, label string) {
	fmt.Printf("Listing all pods from %s namespace with label %s\n", namespace, label)
	podList, err := client.CoreV1().Pods(namespace).List(context.TODO(), v1.ListOptions{LabelSelector: label})
	if err != nil {
		panic(err.Error())
	}

	for i, pod := range podList.Items {
		fmt.Printf("%d: %s, %v\n", i, pod.Name, pod.ObjectMeta.GetLabels())
	}
}

func ListAllPodsFromNamespaceWithField(client *kubernetes.Clientset, namespace string, field string) {
	fmt.Printf("Listing all pods from %s namespace with field %s\n", namespace, field)
	podList, err := client.CoreV1().Pods(namespace).List(context.TODO(), v1.ListOptions{FieldSelector: field})
	if err != nil {
		panic(err.Error())
	}

	for i, pod := range podList.Items {
		fmt.Printf("%d: %s, %s\n", i, pod.Name, pod.Status.Phase)
	}
}

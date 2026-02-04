package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func getKubeClient() *kubernetes.Clientset {
	// Try in-cluster config first
	config, err := rest.InClusterConfig()
	if err != nil {
		kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			panic(err)
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return clientset
}

func main() {
	var edgeOnline bool
	flag.BoolVar(&edgeOnline, "edge-online", true, "Simulate edge connectivity")
	flag.Parse()

	client := getKubeClient()

	factory := informers.NewSharedInformerFactory(client, 0)
	deployInformer := factory.Apps().V1().Deployments().Informer()

	deployInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			deploy := obj.(*appsv1.Deployment)
			handleDeployment(client, deploy, edgeOnline)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			deploy := newObj.(*appsv1.Deployment)
			handleDeployment(client, deploy, edgeOnline)
		},
	})

	stop := make(chan struct{})
	fmt.Println("Mini Edge Agent running...")
	factory.Start(stop)

	cache.WaitForCacheSync(stop, deployInformer.HasSynced)
	select {}
}

func handleDeployment(client *kubernetes.Clientset, deploy *appsv1.Deployment, edgeOnline bool) {
	ctx := context.TODO()

	if !edgeOnline && *deploy.Spec.Replicas != 0 {
		fmt.Printf("Edge offline → scaling %s to 0\n", deploy.Name)
		var zero int32 = 0
		deploy.Spec.Replicas = &zero
		client.AppsV1().
			Deployments(deploy.Namespace).
			Update(ctx, deploy, metav1.UpdateOptions{})
	}
}

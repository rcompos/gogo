package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	//"strings"
	//"regexp"
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// This program lists the pods in a cluster
//
func main() {
	var ns string
	var nsall bool
	flag.StringVar(&ns, "namespace", "default", "K8s namespace")
	flag.BoolVar(&nsall, "all-namespaces", false, "All namespaces")
	flag.Parse()
	fmt.Println("namespace : ", ns)
	//fmt.Println("all-namespaces : ", nsall)

	// Bootstrap k8s configuration from local 	Kubernetes config file
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	log.Println("Using kubeconfig file: ", kubeconfig)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	// Create an rest client not targeting specific API version
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	if nsall == true {
		// Get all namespaces in cluster
		namespaces, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
		if err != nil {
			log.Fatalln("failed to get namespaces:", err)
		}
		for _, namespace := range namespaces.Items {

			fmt.Println("Namespace: ", namespace.Name)
			getpod(namespace.Name)
		}
	} else {
		fmt.Println("Namespace: ", ns)
		getpod(ns)
	}

}

func getpod(ns string) {

	//fmt.Println(ns)

    // Bootstrap k8s configuration from local   Kubernetes config file
    kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
    //log.Println("Using kubeconfig file: ", kubeconfig)
    config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
    if err != nil {
        log.Fatal(err)
    }

	// Create an rest client not targeting specific API version
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

		//fmt.Println("####  ", ns, "  ####")
		pods, err := clientset.CoreV1().Pods(string(ns)).List(metav1.ListOptions{})
		if err != nil {
			log.Fatalln("failed to get pods:", err)
		}
		for _, pod := range pods.Items {

			//b := pod.Status.ContainerStatuses
		    //out, err := json.Marshal(b)
			//if err != nil {
			//	panic (err)
			//}
			//fmt.Println(string(out)) 

			c := pod.Status.Phase
		    outc, err := json.Marshal(c)
			if err != nil {
				panic (err)
			}
			fmt.Println("OUTC: ", string(outc)) 

			//containsy := strings.Contains(string(outc), "\"lastState\":{\"terminated\"")
			////containsy := strings.Contains(string(out), "\"lastState\":{},\"ready\"")
			//if containsy == true {
			//	fmt.Println(string(out)) 
			//}
		}
}

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	_ "strings"
	_ "regexp"
	"encoding/json"
	//"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

//
// This program lists specified pods in a Kubernetes cluster
//
func main() {
	var ns string
	var nsall bool
	flag.StringVar(&ns, "n", "default", "Defined namespace")
	flag.BoolVar(&nsall, "a", false, "All namespaces")
	flag.Parse()

	// Bootstrap k8s configuration from local Kubernetes config file
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	//log.Println("# Using kubeconfig file: ", kubeconfig)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	// Create an rest client not targeting specific API version
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// Run for a single namespace or loop over all namespaces
	if nsall == true {
		// Get all namespaces in cluster
		namespaces, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
		if err != nil {
			log.Fatalln("failed to get namespaces:", err)
		}
		//fmt.Println("# Type of namespaces: ", reflect.TypeOf(namespaces))
		//fmt.Println("# Value of namespaces: ", reflect.ValueOf(namespaces))
		for _, namespace := range namespaces.Items {
			//fmt.Println("# Namespace: ", namespace.Name)
			getpod(namespace.Name, clientset)
		}
	} else {
		//fmt.Println("# Value of namespaces: ", reflect.ValueOf(ns))
		getpod(ns, clientset)
	}

}

func getpod(ns string, c *kubernetes.Clientset) {

	pods, err := c.CoreV1().Pods(string(ns)).List(metav1.ListOptions{
      //LabelSelector: "app=nginx-ingress",
      //FieldSelector: "status.phase!=Running",
      //FieldSelector: "status.phase=Running",
    })
	if err != nil {
		log.Fatalln("failed to get pods:", err)
	}

    //reggie := regexp.MustCompile(`.*Terminated.*`)
//    reggie, _ := regexp.Compile(`.*"terminated".*`)
//
	for _, pod := range pods.Items {

//        cs := pod.Status.ContainerStatuses
//		outc, err := json.Marshal(cs)
        pname := pod.Name
        //pname := pod.Status.Phase
		outc, err := json.Marshal(pname)
		if err != nil {
			panic(err)
		}
//
//		//fmt.Println("###", string(outc))
//		//fmt.Printf("%v\n", reggie.FindString(string(outc)))
//
	fmt.Printf("%v\n", string(outc))
	}


}

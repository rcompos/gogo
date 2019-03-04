package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	_ "regexp"
	_ "strings"
	//"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

//
// Delete non-running Kubernetes pods
//
func main() {
	var ns string
	var nsall bool
	flag.StringVar(&ns, "n", "default", "Defined namespace")
	flag.BoolVar(&nsall, "a", false, "All namespaces")
	flag.Parse()

	// Bootstrap k8s configuration from local Kubernetes config file
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
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
		for _, namespace := range namespaces.Items {
			DeletePodsNotRunning(namespace.Name, clientset)
		}
	} else {
		DeletePodsNotRunning(ns, clientset)
	}

} // func main

func DeletePod(ns string, pod string, c *kubernetes.Clientset) {
	err := c.CoreV1().Pods(ns).Delete(pod, &metav1.DeleteOptions{})
	if err != nil {
		log.Printf("Error deleting pod %s, %s", pod, err)
	}
	fmt.Println("Deleted! ", ns, pod)
}

func DeletePodsNotRunning(ns string, c *kubernetes.Clientset) []string {

	pods, err := c.CoreV1().Pods(string(ns)).List(metav1.ListOptions{
		FieldSelector: "status.phase!=Running",
	})
	if err != nil {
		log.Fatalln("failed to get pods:", err)
	}

	var plist []string
	for _, pod := range pods.Items {

		pname := pod.Name
		outc, err := json.Marshal(pname)
		if err != nil {
			panic(err)
		}
		TargetPod := trimQuote(string(outc))
		plist = append(plist, TargetPod)
		DeletePod(ns, TargetPod, c)
	}
	return plist
}

func trimQuote(s string) string {
	if len(s) >= 2 {
		if s[0] == '"' && s[len(s)-1] == '"' {
			return s[1 : len(s)-1]
		}
	}
	return s
}

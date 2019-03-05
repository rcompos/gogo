package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"bufio"
	"path/filepath"
	_ "regexp"
	_ "strings"
	//"reflect"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

//
// List non-running Kubernetes pods
//
func main() {
	var ns string
	var nsall, Delete, Force bool
	flag.StringVar(&ns, "n", "default", "Define namespace")
	flag.BoolVar(&nsall, "a", false, "All namespaces")
	flag.BoolVar(&Delete, "d", false, "Delete pods")
	flag.BoolVar(&Force, "f", false, "Force without confirmation")
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

	if Force != true { prompt() }

	// Run for a single namespace or loop over all namespaces
	if nsall == true {
		// Get all namespaces in cluster
		namespaces, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
		if err != nil {
			log.Fatalln("failed to get namespaces:", err)
		}
		for _, namespace := range namespaces.Items {
			PodsAll := GetPods(namespace.Name, clientset)
			if Delete != true { 
				ListPods(namespace.Name, clientset, PodsAll)
			} else {
				DeletePods(namespace.Name, clientset, PodsAll)
			}
		}
	} else {
		PodsAll := GetPods(ns, clientset)
		if Delete != true { 
			ListPods(ns, clientset, PodsAll)
		} else {
			DeletePods(ns, clientset, PodsAll)
		}
	}

} // func main

func GetPods(ns string, c *kubernetes.Clientset) *v1.PodList {
	pods, err := c.CoreV1().Pods(string(ns)).List(metav1.ListOptions{
		FieldSelector: "status.phase!=Running",
	})
	if err != nil {
		log.Fatalln("failed to get pods:", err)
	}
	return pods
}

func ListPods(ns string, c *kubernetes.Clientset, pods *v1.PodList ) {
	for _, pod := range pods.Items {
		pname := pod.Name
		outc, err := json.Marshal(pname)
		if err != nil {
			panic(err)
		}
		TargetPod := trimQuote(string(outc))
		fmt.Println("To be Deleted! ", ns, TargetPod)
	}
}

func DeletePods(ns string, c *kubernetes.Clientset, pods *v1.PodList ) {
	var plist []string
	for _, pod := range pods.Items {
		pname := pod.Name
		outc, err := json.Marshal(pname)
		if err != nil {
			panic(err)
		}
		TargetPod := trimQuote(string(outc))
		//DeletePod(ns, TargetPod, c)
		//err := c.CoreV1().Pods(ns).Delete(TargetPod, &metav1.DeleteOptions{})
		//if err != nil {
		//	log.Printf("Error deleting pod %s, %s", TargetPod, err)
		//}
		fmt.Println("Deleted! ", ns, pod)
		//Todo:  Add if success then append
		plist = append(plist, TargetPod)
	}
}

func trimQuote(s string) string {
	if len(s) >= 2 {
		if s[0] == '"' && s[len(s)-1] == '"' {
			return s[1 : len(s)-1]
		}
	}
	return s
}

func prompt() {
    fmt.Printf("-> Press Return key to continue.")
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        break
    }
    if err := scanner.Err(); err != nil {
        panic(err)
    }
    fmt.Println()
}
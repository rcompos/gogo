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
// List non-running Kubernetes pods
//
func main() {
	var ns string
	var nsall bool
	var pl []string
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
			//fmt.Println(string(getpod(namespace.Name, clientset)))
			val := getpod(namespace.Name, clientset)
			if len(val) > 0 {
				pl = append(pl, val...)
			}
		}
	} else {
		//fmt.Println("# Value of namespaces: ", reflect.ValueOf(ns))
		//fmt.Println(string(getpod(ns, clientset)))
		pl = getpod(ns, clientset)
	}

	for j, _ := range pl {
		fmt.Println(string(pl[j]))
	}
}

func getpod(ns string, c *kubernetes.Clientset) []string {

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

	var plist []string
	for i, pod := range pods.Items {

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

	// Trim surrounding quotes
	//fmt.Printf("%v\n", string(outc))
	//if len(outc) > 0 && outc[0] == '"' {
	//	outc = outc[1:]
	//}
	//if len(outc) > 0 && outc[len(outc) - 1] == '"' {
	//	outc = outc[:len(outc) - 1]
	//}
	////outc = outc[1 : len(outc)-1]
	//fmt.Printf("%v\n", string(outc))
		fmt.Printf("%s %s\n", ns, trimQuote(string(outc)))
		fmt.Println("i: ", i)
		plist = append(plist, trimQuote(string(outc)))
	}
	return plist
}

func trimQuote(s string) string {
	if len(s) >= 2 {
		if s[0] == '"' && s[len(s)-1] == '"' {
			return s[1:len(s)-1]
		}
	}
	return s
}



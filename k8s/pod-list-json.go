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
	"reflect"

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
		fmt.Println("Type of namespaces: ", reflect.TypeOf(namespaces))
		fmt.Println("Value of namespaces: ", reflect.ValueOf(namespaces))
		if err != nil {
			log.Fatalln("failed to get namespaces:", err)
		}
		for _, namespace := range namespaces.Items {

			fmt.Println("Namespace: ", namespace.Name)
			getpod(namespace.Name, clientset)
		}
	} else {
		fmt.Println("Namespace: ", ns)
		fmt.Println("Value of namespaces: ", reflect.ValueOf(ns))
		getpod(ns, clientset)
	}

}

func getpod(ns string, c *kubernetes.Clientset) {

	//fmt.Println("####  ", ns, "  ####")
	pods, err := c.CoreV1().Pods(string(ns)).List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get pods:", err)
	}
	for _, pod := range pods.Items {

		b := pod.Status.ContainerStatuses
		fmt.Println("Type of b: ", reflect.TypeOf(b))
		fmt.Println("Value of b: ", reflect.ValueOf(b))

		fmt.Println("b> ", b)

		outb, err := json.Marshal(b)
		if err != nil {
			panic(err)
		}
		fmt.Println("B> ", string(outb))
		fmt.Println(" ")

		//c := pod.Status.Phase
		c := pod.Status.Conditions
		fmt.Println("Type of c: ", reflect.TypeOf(c))
		fmt.Println("Value of c: ", reflect.ValueOf(c))
		for j := 0; j < len(c); j++ {
			//d := c[j].Type
			d := c[j].Type
			outc, err := json.Marshal(d)
			if err != nil {
				panic(err)
			}
			fmt.Println("C> ", string(outc))
		}

		byt := []byte(`{"num":6.13,"strs":["alpha","beta","gamma","delta"]}`)
	    var dat map[string]interface{}
	    if err := json.Unmarshal(byt, &dat); err != nil {
			panic(err)
		}
		//fmt.Println(dat)
		strs := dat["strs"].([]interface{})
		str0 := strs[0].(string)
		str1 := strs[1].(string)
		fmt.Println("D> ", str0)
		fmt.Println("D> ", str1)

		//containsy := strings.Contains(string(outc), "\"lastState\":{\"terminated\"")
		////containsy := strings.Contains(string(out), "\"lastState\":{},\"ready\"")
		//if containsy == true {
		//	fmt.Println(string(out))
		//}
	}

}

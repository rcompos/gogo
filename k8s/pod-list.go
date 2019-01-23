package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	//"regexp"
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

//
//
func main() {
	var ns, nsall string
	flag.StringVar(&ns, "namespace", "default", "K8s namespace")
	flag.StringVar(&nsall, "all-namespaces", "", "All namespaces")
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

	if nsall != "" {

		fmt.Println("NSALL: ", nsall)

		// Get all namespaces in cluster
		namespaces, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
		if err != nil {
			log.Fatalln("failed to get namespaces:", err)
		}

		for _, namespace := range namespaces.Items {
			fmt.Println("####  ", namespace.Name, "  ####")
			//pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
			//pods, err := clientset.CoreV1().Pods(ns).List(metav1.ListOptions{})
			pods, err := clientset.CoreV1().Pods(string(namespace.Name)).List(metav1.ListOptions{})
			//pods, err := clientset.CoreV1().Pods(ns).List(metav1.ListOptions{
			//	LabelSelector: "function=isprime",
			//})
			if err != nil {
				log.Fatalln("failed to get pods:", err)
			}

			// print pods
			//for i, pod := range pods.Items {
			//	fmt.Printf("[%d] %s\n", i, pod.GetName())
			//}
			//reggie := regexp.MustCompile(`.*Terminated.*`)
			for _, pod := range pods.Items {
				//fmt.Println(pod.Name, pod.Status.PodIP)
				//fmt.Println(pod.Name, pod.Status)
				//fmt.Println(pod.Name, pod.Status.ContainerStatuses)
				//fmt.Printf("%q\n", reggie.FindString(pod.Status.ContainerStatuses))
				//fmt.Printf("%q\n", reggie.MatchString("asdfaTrminated kwik"))
				b := pod.Status.ContainerStatuses
				//a := &b
			    out, err := json.Marshal(b)
				if err != nil {
					panic (err)
				}
				//fmt.Println(string(out))
				fmt.Println("OUT: ", string(out))
				//fmt.Println(a)
				//fmt.Println(b)
				fmt.Println(strings.Contains(string(out), "terminated"))
				//fmt.Printf(strings.Contains(pod.Status.ContainerStatuses, "Terminated"))
				//fmt.Printf("%s\n", strings.Contains(out, "Terminated"))
			}
		}

	} else {

		fmt.Println("####  ", ns, "  ####")
		//pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		//pods, err := clientset.CoreV1().Pods(ns).List(metav1.ListOptions{})
		pods, err := clientset.CoreV1().Pods(string(ns)).List(metav1.ListOptions{})
		//pods, err := clientset.CoreV1().Pods(ns).List(metav1.ListOptions{
		//	LabelSelector: "function=isprime",
		//})
		if err != nil {
			log.Fatalln("failed to get pods:", err)
		}

		// print pods
		//for i, pod := range pods.Items {
		//	fmt.Printf("[%d] %s\n", i, pod.GetName())
		//}
		//reggie := regexp.MustCompile(`.*Terminated.*`)
		for _, pod := range pods.Items {
			//fmt.Println(pod.Name, pod.Status.PodIP)
			//fmt.Println(pod.Name, pod.Status)
			//fmt.Println(pod.Name, pod.Status.ContainerStatuses)
			//fmt.Printf("%q\n", reggie.FindString(pod.Status.ContainerStatuses))
			//fmt.Printf("%q\n", reggie.MatchString("asdfaTrminated kwik"))
			b := pod.Status.ContainerStatuses
			//a := &b
		    out, err := json.Marshal(b)
			if err != nil {
				panic (err)
			}
			//fmt.Println(string(out))
			//fmt.Println("OUT: ", string(out))
			//fmt.Println(a)
			//fmt.Println(b)
			//fmt.Println(strings.Contains(string(out), "\"lastState\":{\"terminated\""))
			containsy := strings.Contains(string(out), "\"lastState\":{\"terminated\"")
			if containsy == true {
				fmt.Println("OUT: ", string(out)) 
			}
			//fmt.Printf(strings.Contains(pod.Status.ContainerStatuses, "Terminated"))
			//fmt.Printf("%s\n", strings.Contains(out, "Terminated"))
		}

	}

}

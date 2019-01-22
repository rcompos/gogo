package main

import ( "flag"
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

// This program lists the pods in a cluster equivalent to
//
// kubectl get pods
//
func main() {
	var ns string
	flag.StringVar(&ns, "namespace", "default", "K8s namespace")
	flag.Parse()
	fmt.Println("K8s Namespace: ", ns)

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

	//pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	pods, err := clientset.CoreV1().Pods(ns).List(metav1.ListOptions{})
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
		a := &b
	    out, err := json.Marshal(a)
		if err != nil {
			panic (err)
		}
		fmt.Println(string(out))
		//fmt.Println(a)
		//fmt.Println(b)
		fmt.Println(strings.Contains(string(out), "heketi"))
		//fmt.Printf(strings.Contains(pod.Status.ContainerStatuses, "Terminated"))
		//fmt.Printf("%s\n", strings.Contains(out, "Terminated"))
	}
}

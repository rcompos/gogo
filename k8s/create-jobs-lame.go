package main

import (
	//"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	_ "regexp"
	_ "strings"
	//"reflect"

    appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

//
// Delete non-running Kubernetes pods
//
func main() {
	var ns string
	flag.StringVar(&ns, "n", "default", "Defined namespace")
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

	jobsClient := clientset.BatchV1().Jobs("hello")
	job := &batchV1.Job{
	    ObjectMeta: metav1.ObjectMeta{
	        Name: "demo-job",
	        Namespace: "demo",
	    },
	    Spec: batchv1.JobSpec{
	        Template: apiv1.PodTemplateSpec{
	            Spec: apiv1.PodSpec{
	                Containers: []apiv1.Container{
	                    {
	                        Name:  "demo-pod",
	                        Image: "nginx",
	                    },
	                },
	           },
	        },
	    },
	}

    fmt.Println("Creating job... ")
    result1, err1 := jobsClient.Create(job)
    if err != nil {
        fmt.Println(err1)
        panic(err1)
    }
    fmt.Printf("Created job %q.\n", result1)



} // func main

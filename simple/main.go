package main

import (
	"log"

	"github.com/johandry/klient"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	name := "apple"
	cm := []byte(`{"apiVersion": "v1", "kind": "ConfigMap", "metadata": { "name": "fruit" }, "data": {	"name": "` + name + `" } }`)

	c := klient.New("", "") // Take the Kubernetes config from the default location (~/.kube/config) and using the default context.
	if err := c.Apply(cm); err != nil {
		log.Fatal("failed to apply the ConfigMap")
	}

	cmFruit, err := c.Clientset.CoreV1().ConfigMaps("default").Get("fruit", metav1.GetOptions{})
	if err != nil {
		log.Fatal("Failed to get the ConfigMap fruits")
	}
	log.Printf("Fruit name: %s", cmFruit.Data["name"])

	if err := c.Delete(cm); err != nil {
		log.Fatal("failed to delete the ConfigMap")
	}
}

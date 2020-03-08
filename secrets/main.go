package main

import (
	"encoding/base64"
	"log"

	"github.com/johandry/klient"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	password := "Super5ecret0!"
	encPasswd := base64.StdEncoding.EncodeToString([]byte(password))
	secret := []byte(`{ "kind": "Secret", "apiVersion": "v1", "metadata": { "name": "appsecret" }, "data": { "password": "` + encPasswd + `" } }`)

	c := klient.NewClient("", "") // Take the Kubernetes config from the default location (~/.kube/config) and using the default context.
	if err := c.Apply(secret); err != nil {
		log.Fatal("failed to apply the Secret")
	}

	appSecret, err := c.Clientset.CoreV1().Secrets("default").Get("appsecret", metav1.GetOptions{})
	if err != nil {
		log.Fatal("Failed to get the Secret fruits")
	}
	log.Printf("Application Password: %s", appSecret.Data["password"])

	if err := c.Delete(secret); err != nil {
		log.Fatal("failed to delete the Secret")
	}
}

package tests

import (
	"io/ioutil"
	"log"

	"k8s.io/minikube/pkg/minikube/constants"
)

func MakeTempDir() string {
	tempDir, err := ioutil.TempDir("", "minipath")
	if err != nil {
		log.Fatal(err)
	}
	constants.Minipath = tempDir
	return tempDir
}

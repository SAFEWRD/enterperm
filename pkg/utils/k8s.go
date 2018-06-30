/*
Copyright 2018 Safewrd Ventures OÜ

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	"os"

	"github.com/golang/glog"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// GetClient gets the kubernetes config from inside a cluster
func GetClient() kubernetes.Interface {
	config, err := rest.InClusterConfig()
	if err != nil {
		glog.Errorln("Could not load in-cluster config")
	}

	client, err := kubernetes.NewForConfig(config)
	return client
}

// GetClientExternal gets the kubernetes client from outside a cluster
func GetClientExternal() kubernetes.Interface {
	config, err := buildConfig()
	if err != nil {
		glog.Fatalln(err)
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		glog.Fatalln(err)
	}
	return client
}

func buildConfig() (*rest.Config, error) {
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		kubeconfig = os.Getenv("HOME") + "/.kube/config"
	}
	return clientcmd.BuildConfigFromFlags("", kubeconfig)
}

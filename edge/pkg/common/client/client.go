/*
Copyright 2021 The KubeEdge Authors.

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

package client

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	"os"
	"sync"
)


var keClient *kubeEdgeClient
var once sync.Once


func InitKubeEdgeClient() {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("127.0.0.1:10550", "")
	if err != nil {
		klog.Errorf("Failed to build config, err: %v", err)
		os.Exit(1)
	}
	kubeConfig.ContentType = runtime.ContentTypeJSON
	kubeClient := kubernetes.NewForConfigOrDie(kubeConfig)


	once.Do(func() {
		keClient = &kubeEdgeClient{
			kubeClient:    kubeClient, //clientset
		}
	})
}


func GetKubeClient() kubernetes.Interface {
	return keClient.kubeClient
}

type kubeEdgeClient struct {
	kubeClient    *kubernetes.Clientset
}

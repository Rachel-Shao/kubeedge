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
	"sync"

	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

var (
	initOnce sync.Once
	keClient *kubeEdgeClient
)

func InitKubeEdgeClient() {
	initOnce.Do(func() {
		config, err := clientcmd.BuildConfigFromFlags("127.0.0.1:10550", "")
		if err != nil {
			klog.Errorf("Failed to build config, err: %v", err)
		}
		config.ContentType = runtime.ContentTypeJSON
		crdClient, err := clientset.NewForConfig(config)
		keClient = &kubeEdgeClient{
			crdClient: crdClient,
		}
	})
}

func GetCRDClient() *clientset.Clientset {
	return keClient.crdClient
}

type kubeEdgeClient struct {
	crdClient *clientset.Clientset
}

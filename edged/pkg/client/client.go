package client

import (
	edgedConfig "github.com/kubeedge/kubeedge/pkg/apis/componentconfig/edged/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	"os"
	"sync"
)

var edClient *EdgedClient
var once sync.Once

func InitEdgedClient(config *edgedConfig.KubeAPIConfig) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags(config.Master, config.KubeConfig)
	if err != nil {
		klog.Errorf("Failed to build config, err: %v", err)
		os.Exit(1)
	}
	kubeConfig.QPS = float32(config.QPS)
	kubeConfig.Burst = int(config.Burst)
	kubeConfig.ContentType = runtime.ContentTypeProtobuf
	kubeClient := kubernetes.NewForConfigOrDie(kubeConfig)

	once.Do(func() {
		edClient = &EdgedClient{
			kubeClient:  kubeClient,
		}
	})
}


func GetKubeClient() kubernetes.Interface {
	return edClient.kubeClient
}

type EdgedClient struct {
	kubeClient  *kubernetes.Clientset
}

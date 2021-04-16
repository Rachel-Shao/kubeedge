/*
Copyright 2019 The KubeEdge Authors.

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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)


const (
	CGroupDriverCGroupFS = "cgroupfs"
	CGroupDriverSystemd  = "systemd"
)


// EdgedConfig indicates the config of edged which get from edged config file
type EdgedConfig struct {
	metav1.TypeMeta

	// KubeAPIConfig indicates the kubernetes cluster info which edged will connected
	// +Required
	KubeAPIConfig *KubeAPIConfig `json:"kubeAPIConfig,omitempty"`

	/*
	// DataBase indicates database info
	// +Required
	DataBase *DataBase `json:"database,omitempty"`   // ???don't known whether to use

	 */

	// Modules indicates Edged modules config
	// +Required
	Modules *Modules `json:"modules,omitempty"`
}



// KubeAPIConfig indicates the configuration for interacting with k8s server
type KubeAPIConfig struct {
	// Master indicates the address of the Kubernetes API server (overrides any value in KubeConfig)
	// such as https://127.0.0.1:8443
	// default ""
	// Note: Can not use "omitempty" option,  It will affect the output of the default configuration file
	Master string `json:"master"`
	// ContentType indicates the ContentType of message transmission when interacting with k8s
	// default "application/vnd.kubernetes.protobuf"
	ContentType string `json:"contentType,omitempty"`
	// QPS to while talking with kubernetes apiserve
	// default 100
	QPS int32 `json:"qps,omitempty"`
	// Burst to use while talking with kubernetes apiserver
	// default 200
	Burst int32 `json:"burst,omitempty"`
	// KubeConfig indicates the path to kubeConfig file with authorization and master location information.
	// default "/root/.kube/config"
	// +Required
	KubeConfig string `json:"kubeConfig"`
}


// Modules indicates the modules of Edged will be use
type Modules struct {
	// Edged indicates edged module config
	// +Required
	Edged *Edged `json:"edged,omitempty"`
}




// Edged indicates the config fo edged module
// edged is lighted-kubelet
type Edged struct {
	// Enable indicates whether edged is enabled,
	// if set to false (for debugging etc.), skip checking other edged configs.
	// default true
	Enable bool `json:"enable,omitempty"`
	// Labels indicates current node labels
	Labels map[string]string `json:"labels,omitempty"`
	// Annotations indicates current node annotations
	Annotations map[string]string `json:"annotations,omitempty"`
	// Taints indicates current node taints
	Taints []v1.Taint `json:"taints,omitempty"`
	// NodeStatusUpdateFrequency indicates node status update frequency (second)
	// default 10
	NodeStatusUpdateFrequency int32 `json:"nodeStatusUpdateFrequency,omitempty"`
	// RuntimeType indicates cri runtime ,support: docker, remote
	// default "docker"
	RuntimeType string `json:"runtimeType,omitempty"`
	// DockerAddress indicates docker server address
	// default "unix:///var/run/docker.sock"
	DockerAddress string `json:"dockerAddress,omitempty"`
	// RemoteRuntimeEndpoint indicates remote runtime endpoint
	// default "unix:///var/run/dockershim.sock"
	RemoteRuntimeEndpoint string `json:"remoteRuntimeEndpoint,omitempty"`
	// RemoteImageEndpoint indicates remote image endpoint
	// default "unix:///var/run/dockershim.sock"
	RemoteImageEndpoint string `json:"remoteImageEndpoint,omitempty"`
	// NodeIP indicates current node ip
	// default get local host ip
	NodeIP string `json:"nodeIP"`
	// ClusterDNS indicates cluster dns
	// Note: Can not use "omitempty" option,  It will affect the output of the default configuration file
	// +Required
	ClusterDNS string `json:"clusterDNS"`
	// ClusterDomain indicates cluster domain
	// Note: Can not use "omitempty" option,  It will affect the output of the default configuration file
	ClusterDomain string `json:"clusterDomain"`
	// EdgedMemoryCapacity indicates memory capacity (byte)
	// default 7852396000
	EdgedMemoryCapacity int64 `json:"edgedMemoryCapacity,omitempty"`
	// PodSandboxImage is the image whose network/ipc namespaces containers in each pod will use.
	// +Required
	// kubeedge/pause:3.1 for x86 arch
	// kubeedge/pause-arm:3.1 for arm arch
	// kubeedge/pause-arm64 for arm64 arch
	// default kubeedge/pause:3.1
	PodSandboxImage string `json:"podSandboxImage,omitempty"`
	// ImagePullProgressDeadline indicates image pull progress dead line (second)
	// default 60
	ImagePullProgressDeadline int32 `json:"imagePullProgressDeadline,omitempty"`
	// RuntimeRequestTimeout indicates runtime request timeout (second)
	// default 2
	RuntimeRequestTimeout int32 `json:"runtimeRequestTimeout,omitempty"`
	// HostnameOverride indicates hostname
	// default os.Hostname()
	HostnameOverride string `json:"hostnameOverride,omitempty"`
	// RegisterNode enables automatic registration
	// default true
	RegisterNode bool `json:"registerNode,omitempty"`
	//RegisterNodeNamespace indicates register node namespace
	// default "default"
	RegisterNodeNamespace string `json:"registerNodeNamespace,omitempty"`
	// InterfaceName indicates interface name
	// default "eth0"
	// DEPRECATED after v1.5
	InterfaceName string `json:"interfaceName,omitempty"`
	// ConcurrentConsumers indicates concurrent consumers for pod add or remove operation
	// default 5
	ConcurrentConsumers int `json:"concurrentConsumers,omitempty"`
	// DevicePluginEnabled indicates enable device plugin
	// default false
	// Note: Can not use "omitempty" option, it will affect the output of the default configuration file
	DevicePluginEnabled bool `json:"devicePluginEnabled"`
	// GPUPluginEnabled indicates enable gpu plugin
	// default false,
	// Note: Can not use "omitempty" option, it will affect the output of the default configuration file
	GPUPluginEnabled bool `json:"gpuPluginEnabled"`
	// ImageGCHighThreshold indicates image gc high threshold (percent)
	// default 80
	ImageGCHighThreshold int32 `json:"imageGCHighThreshold,omitempty"`
	// ImageGCLowThreshold indicates image gc low threshold (percent)
	// default 40
	ImageGCLowThreshold int32 `json:"imageGCLowThreshold,omitempty"`
	// MaximumDeadContainersPerPod indicates max num dead containers per pod
	// default 1
	MaximumDeadContainersPerPod int32 `json:"maximumDeadContainersPerPod,omitempty"`
	// CGroupDriver indicates container cgroup driver, support: cgroupfs, systemd
	// default "cgroupfs"
	// +Required
	CGroupDriver string `json:"cgroupDriver,omitempty"`
	// NetworkPluginName indicates the name of the network plugin to be invoked,
	// if an empty string is specified, use noop plugin
	// default ""
	NetworkPluginName string `json:"networkPluginName,omitempty"`
	// CNIConfDir indicates the full path of the directory in which to search for CNI config files
	// default "/etc/cni/net.d"
	CNIConfDir string `json:"cniConfDir,omitempty"`
	// CNIBinDir indicates a comma-separated list of full paths of directories
	// in which to search for CNI plugin binaries
	// default "/opt/cni/bin"
	CNIBinDir string `json:"cniBinDir,omitempty"`
	// CNICacheDir indicates the full path of the directory in which CNI should store cache files
	// default "/var/lib/cni/cache"
	CNICacheDir string `json:"cniCacheDirs,omitempty"`
	// NetworkPluginMTU indicates the MTU to be passed to the network plugin
	// default 1500
	NetworkPluginMTU int32 `json:"networkPluginMTU,omitempty"`
	// CgroupsPerQOS enables QoS based Cgroup hierarchy: top level cgroups for QoS Classes
	// And all Burstable and BestEffort pods are brought up under their
	// specific top level QoS cgroup.
	// Default: true
	CgroupsPerQOS bool `json:"cgroupsPerQOS"`
	// CgroupRoot is the root cgroup to use for pods.
	// If CgroupsPerQOS is enabled, this is the root of the QoS cgroup hierarchy.
	// Default: ""
	CgroupRoot string `json:"cgroupRoot"`
	// EdgeCoreCgroups is the absolute name of cgroups to isolate the edgecore in
	// Dynamic Kubelet Config (beta): This field should not be updated without a full node
	// reboot. It is safest to keep this value the same as the local config.
	// Default: ""
	EdgeCoreCgroups string `json:"edgeCoreCgroups,omitempty"`
	// systemCgroups is absolute name of cgroups in which to place
	// all non-kernel processes that are not already in a container. Empty
	// for no container. Rolling back the flag requires a reboot.
	// Dynamic Kubelet Config (beta): This field should not be updated without a full node
	// reboot. It is safest to keep this value the same as the local config.
	// Default: ""
	SystemCgroups string `json:"systemCgroups,omitempty"`
	// How frequently to calculate and cache volume disk usage for all pods
	// Dynamic Kubelet Config (beta): If dynamically updating this field, consider that
	// shortening the period may carry a performance impact.
	// Default: "1m"
	VolumeStatsAggPeriod time.Duration `json:"volumeStatsAggPeriod,omitempty"`
	// EnableMetrics indicates whether enable the metrics
	// default true
	EnableMetrics bool `json:"enableMetrics,omitempty"`
}










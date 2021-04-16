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
	"github.com/kubeedge/kubeedge/common/constants"
	"github.com/kubeedge/kubeedge/pkg/util"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"path"
)

// NewDefaultEdgedConfig returns a full EdgedConfig object
func NewDefaultEdgedConfig() *EdgedConfig {
	hostnameOverride, err := os.Hostname()
	if err != nil {
		hostnameOverride = constants.DefaultHostnameOverride
	}
	localIP, _ := util.GetLocalIP(hostnameOverride)

	ed := &EdgedConfig{
		TypeMeta: metav1.TypeMeta{
			Kind: Kind,
			APIVersion: path.Join(GroupName, APIVersion),
		},
		KubeAPIConfig: &KubeAPIConfig{
			Master:      "127.0.0.1:10550", // metaserver端口
			ContentType: constants.DefaultKubeContentType,
			QPS:         constants.DefaultKubeQPS,
			Burst:       constants.DefaultKubeBurst,
			KubeConfig:  "", // ??? empty?
		},
		Modules: &Modules{
			Edged: &Edged{
				Enable:                      true,
				Labels:                      map[string]string{},
				Annotations:                 map[string]string{},
				Taints:                      []v1.Taint{},
				NodeStatusUpdateFrequency:   constants.DefaultNodeStatusUpdateFrequency,
				RuntimeType:                 constants.DefaultRuntimeType,
				DockerAddress:               constants.DefaultDockerAddress,
				RemoteRuntimeEndpoint:       constants.DefaultRemoteRuntimeEndpoint,
				RemoteImageEndpoint:         constants.DefaultRemoteImageEndpoint,
				NodeIP:                      localIP,
				ClusterDNS:                  "",
				ClusterDomain:               "",
				ConcurrentConsumers:         constants.DefaultConcurrentConsumers,
				EdgedMemoryCapacity:         constants.DefaultEdgedMemoryCapacity,
				PodSandboxImage:             util.GetPodSandboxImage(),
				ImagePullProgressDeadline:   constants.DefaultImagePullProgressDeadline,
				RuntimeRequestTimeout:       constants.DefaultRuntimeRequestTimeout,
				HostnameOverride:            hostnameOverride,
				RegisterNodeNamespace:       constants.DefaultRegisterNodeNamespace,
				RegisterNode:                true,
				DevicePluginEnabled:         false,
				GPUPluginEnabled:            false,
				ImageGCHighThreshold:        constants.DefaultImageGCHighThreshold,
				ImageGCLowThreshold:         constants.DefaultImageGCLowThreshold,
				MaximumDeadContainersPerPod: constants.DefaultMaximumDeadContainersPerPod,
				CGroupDriver:                CGroupDriverCGroupFS,
				CgroupsPerQOS:               true,
				CgroupRoot:                  constants.DefaultCgroupRoot,
				NetworkPluginName:           "",
				CNIConfDir:                  constants.DefaultCNIConfDir,
				CNIBinDir:                   constants.DefaultCNIBinDir,
				CNICacheDir:                 constants.DefaultCNICacheDir,
				NetworkPluginMTU:            constants.DefaultNetworkPluginMTU,
				VolumeStatsAggPeriod:        constants.DefaultVolumeStatsAggPeriod,
				EnableMetrics:               true,
			},
		},
	}
	return ed
}



// NewMinEdgedConfig returns a common EdgedConfig object
func NewMinEdgedConfig() *EdgedConfig {
	hostnameOverride, err := os.Hostname()
	if err != nil {
		hostnameOverride = constants.DefaultHostnameOverride
	}
	localIP, _ := util.GetLocalIP(hostnameOverride)

	ed := &EdgedConfig{
		TypeMeta: metav1.TypeMeta{
			Kind: Kind,
			APIVersion: path.Join(GroupName, APIVersion),
		},
		KubeAPIConfig: &KubeAPIConfig{
			Master:     "127.0.0.1:10550",
			KubeConfig: "", // constants.DefaultKubeConfig
		},
		Modules: &Modules{
			Edged: &Edged{
				RuntimeType:           constants.DefaultRuntimeType,
				RemoteRuntimeEndpoint: constants.DefaultRemoteRuntimeEndpoint,
				RemoteImageEndpoint:   constants.DefaultRemoteImageEndpoint,
				DockerAddress:         constants.DefaultDockerAddress,
				NodeIP:                localIP,
				ClusterDNS:            "",
				ClusterDomain:         "",
				PodSandboxImage:       util.GetPodSandboxImage(),
				HostnameOverride:      hostnameOverride,
				DevicePluginEnabled:   false,
				GPUPluginEnabled:      false,
				CGroupDriver:          CGroupDriverCGroupFS,
				CgroupsPerQOS:         true,
				CgroupRoot:            constants.DefaultCgroupRoot,
			},
		},
	}
	return ed
}
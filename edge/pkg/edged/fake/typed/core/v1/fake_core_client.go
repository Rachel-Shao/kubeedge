package v1

import (
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	fakecorev1 "k8s.io/client-go/kubernetes/typed/core/v1/fake"
)

type FakeCoreV1 struct {
	fakecorev1.FakeCoreV1
}

func (c *FakeCoreV1) Nodes() corev1.NodeInterface {
	return &FakeNodes{fakecorev1.FakeNodes{Fake: &c.FakeCoreV1}}
}

func (c *FakeCoreV1) PersistentVolumes() corev1.PersistentVolumeInterface {
	return &FakePersistentVolumes{fakecorev1.FakePersistentVolumes{Fake: &c.FakeCoreV1}}
}

func (c *FakeCoreV1) PersistentVolumeClaims(namespace string) corev1.PersistentVolumeClaimInterface {
	return &FakePersistentVolumeClaims{fakecorev1.FakePersistentVolumeClaims{Fake: &c.FakeCoreV1}, namespace}
}

func (c *FakeCoreV1) ConfigMaps(namespace string) corev1.ConfigMapInterface {
	return &FakeConfigMap{fakecorev1.FakeConfigMaps{Fake: &c.FakeCoreV1}, namespace}
}

package v1

import (
	"context"

	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fakestoragev1 "k8s.io/client-go/kubernetes/typed/storage/v1/fake"

	"github.com/kubeedge/kubeedge/edge/pkg/common/client"
)

// FakeVolumeAttachments implements PersistentVolumeInterface
type FakeVolumeAttachments struct {
	fakestoragev1.FakeVolumeAttachments
}

// Get takes name of the persistentVolume, and returns the corresponding persistentVolume object
func (c *FakeVolumeAttachments) Get(ctx context.Context, name string, options metav1.GetOptions) (result *storagev1.VolumeAttachment, err error) {
	return client.GetKubeClient().StorageV1().VolumeAttachments().Get(ctx, name, options)
}

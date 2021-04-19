package v1

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fakecorev1 "k8s.io/client-go/kubernetes/typed/core/v1/fake"

	"github.com/kubeedge/kubeedge/edge/pkg/common/client"
)

// FakeNodes implements NodeInterface
type FakeNodes struct {
	fakecorev1.FakeNodes
}

// Get takes name of the node, and returns the corresponding node object
func (c *FakeNodes) Get(ctx context.Context, name string, options metav1.GetOptions) (result *corev1.Node, err error) {
	return client.GetKubeClient().CoreV1().Nodes().Get(ctx, name, options)
}

// Update takes the representation of a node and updates it
func (c *FakeNodes) Update(ctx context.Context, node *corev1.Node, opts metav1.UpdateOptions) (result *corev1.Node, err error) {
	_, err = client.GetKubeClient().CoreV1().Nodes().Update(ctx, node, opts)
	if err != nil {
		return nil, err
	}
	return node, nil
}

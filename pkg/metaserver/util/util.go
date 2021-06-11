package util

import (
	"context"
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/klog/v2"

	beehiveModel "github.com/kubeedge/beehive/pkg/core/model"
	"github.com/kubeedge/kubeedge/edge/pkg/common/client"
)

const (
	EmptyString = ""
)

var (
	CRDResourceToKind = make(map[string]string)
	CRDKindToResource = make(map[string]string)
)

func UpdateCrdMap() error {
	list, err := client.GetCRDClient().ApiextensionsV1beta1().CustomResourceDefinitions().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, crd := range list.Items {
		kind := crd.Spec.Names.Kind
		plural := crd.Spec.Names.Plural
		CRDResourceToKind[plural] = kind
		CRDKindToResource[kind] = plural
	}
	klog.Infof("CRD Resource-Kind map updated")
	return nil
}

// MetaType is generally consisted of apiversion, kind like:
// {
// apiVersion: apps/v1
// kind: Deployments
// }
// TODO: support crd
func SetMetaType(obj runtime.Object) error {
	accessor, err := meta.Accessor(obj)
	if err != nil {
		return err
	}
	//gvr,_,_ := apiserverlite.ParseKey(accessor.GetSelfLink())
	kinds, _, err := scheme.Scheme.ObjectKinds(obj)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	gvk := kinds[0]
	obj.GetObjectKind().SetGroupVersionKind(gvk)
	klog.V(4).Infof("[metaserver]successfully set MetaType for obj %v, %+v", obj.GetObjectKind(), accessor.GetName())
	return nil
}

// Sometimes, we need guess kind according to resource:
// 1. In most cases, is like pods to Pod,
// 2. In some unusual cases, requires special treatment like endpoints to Endpoints
func UnsafeResourceToKind(r string) string {
	if len(r) == 0 {
		return r
	}
	unusualResourceToKind := map[string]string{
		"endpoints":                    "Endpoints",
		"nodes":                        "Node",
		"services":                     "Service",
		"customresourcedefinitions":    "CustomResourceDefinition",
		"customresourcedefinitionlist": "CustomResourceDefinitionList",
	}
	if v, isUnusual := unusualResourceToKind[r]; isUnusual {
		return v
	}
	if v, isCRD := CRDResourceToKind[r]; isCRD {
		return v
	}
	k := strings.Title(r)
	switch {
	case strings.HasSuffix(k, "ies"):
		return strings.TrimSuffix(k, "ies") + "y"
	case strings.HasSuffix(k, "es"):
		return strings.TrimSuffix(k, "es")
	case strings.HasSuffix(k, "s"):
		return strings.TrimSuffix(k, "s")
	}
	return k
}

func UnsafeKindToResource(k string) string {
	if len(k) == 0 || len(k) == 1 {
		return k
	}
	unusualKindToResource := map[string]string{
		"Endpoints":                    "endpoints",
		"CustomResourceDefinition":     "customresourcedefinitions",
		"CustomResourceDefinitionList": "customresourcedefinitionlist",
	}
	if v, isUnusual := unusualKindToResource[k]; isUnusual {
		return v
	}
	if v, isCRD := CRDKindToResource[k]; isCRD {
		return v
	}
	r := strings.ToLower(k)
	switch string(r[len(r)-1]) {
	case "s":
		return r + "es"
	case "y":
		if string(r[len(r)-2]) == "a" || string(r[len(r)-2]) == "e" || string(r[len(r)-2]) == "i" ||
			string(r[len(r)-2]) == "o" || string(r[len(r)-2]) == "u" {
			return r + "s"
		}
		return strings.TrimSuffix(r, "y") + "ies"
	}

	return r + "s"
}

func UnstructuredAttr(obj runtime.Object) (labels.Set, fields.Set, error) {
	switch obj.GetObjectKind().GroupVersionKind().Kind {
	case "Pod":
		metadata, err := meta.Accessor(obj)
		if err != nil {
			return nil, nil, err
		}
		setMap := make(fields.Set)
		if metadata.GetName() != "" {
			setMap["metadata.name"] = metadata.GetName()
		}
		if metadata.GetNamespace() != "" {
			setMap["metadata.namespaces"] = metadata.GetNamespace()
		}
		unstrObj, ok := obj.(*unstructured.Unstructured)
		if ok {
			value, found, err := unstructured.NestedString(unstrObj.Object, "spec", "nodeName")
			if found && err == nil {
				setMap["spec.nodeName"] = value
			}
		}
		return metadata.GetLabels(), setMap, nil
	default:
		return storage.DefaultNamespaceScopedAttr(obj)
	}
}

// GetMessageUID returns the UID of the object in message
func GetMessageAPIVerison(msg *beehiveModel.Message) string {
	obj, ok := msg.Content.(runtime.Object)
	if ok {
		return obj.GetObjectKind().GroupVersionKind().GroupVersion().String()
	}
	return ""
}

// GetMessageUID returns the UID of the object in message
func GetMessageResourceType(msg *beehiveModel.Message) string {
	obj, ok := msg.Content.(runtime.Object)
	if ok {
		return UnsafeKindToResource(obj.GetObjectKind().GroupVersionKind().Kind)
	}
	return ""
}

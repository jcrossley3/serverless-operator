/*
Copyright 2018 The Knative Authors

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
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"knative.dev/pkg/apis"
	"knative.dev/pkg/apis/duck"
	v1 "knative.dev/pkg/apis/duck/v1"
	"knative.dev/pkg/apis/duck/v1beta1"
)

// +genduck

// Addressable provides a generic mechanism for a custom resource
// definition to indicate a destination for message delivery.
//
// Addressable is the schema for the destination information. This is
// typically stored in the object's `status`, as this information may
// be generated by the controller.
type Addressable struct {
	v1beta1.Addressable `json:",omitempty"`

	Hostname string `json:"hostname,omitempty"`
}

var (
	// Addressable is an Implementable "duck type".
	_ duck.Implementable = (*Addressable)(nil)
	// Addressable is a Convertible type.
	_ apis.Convertible = (*Addressable)(nil)
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AddressableType is a skeleton type wrapping Addressable in the manner we expect
// resource writers defining compatible resources to embed it.  We will
// typically use this type to deserialize Addressable ObjectReferences and
// access the Addressable data.  This is not a real resource.
type AddressableType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status AddressStatus `json:"status"`
}

// AddressStatus shows how we expect folks to embed Addressable in
// their Status field.
type AddressStatus struct {
	Address *Addressable `json:"address,omitempty"`
}

var (
	// Verify AddressableType resources meet duck contracts.
	_ duck.Populatable = (*AddressableType)(nil)
	_ apis.Listable    = (*AddressableType)(nil)
)

// GetFullType implements duck.Implementable
func (*Addressable) GetFullType() duck.Populatable {
	return &AddressableType{}
}

// ConvertUp implements apis.Convertible
func (a *Addressable) ConvertUp(ctx context.Context, to apis.Convertible) error {
	url := a.GetURL()
	switch sink := to.(type) {
	case *v1.Addressable:
		sink.URL = url.DeepCopy()
		return nil
	case *v1beta1.Addressable:
		sink.URL = url.DeepCopy()
		return nil
	default:
		return fmt.Errorf("unknown version, got: %T", to)
	}
}

// ConvertDown implements apis.Convertible
func (a *Addressable) ConvertDown(ctx context.Context, from apis.Convertible) error {
	switch source := from.(type) {
	case *v1.Addressable:
		a.URL = source.URL.DeepCopy()
		return nil
	case *v1beta1.Addressable:
		a.URL = source.URL.DeepCopy()
		return nil
	default:
		return fmt.Errorf("unknown version, got: %T", from)
	}
}

// Populate implements duck.Populatable
func (t *AddressableType) Populate() {
	t.Status = AddressStatus{
		&Addressable{
			// Populate ALL fields
			Addressable: v1beta1.Addressable{
				URL: &apis.URL{
					Scheme: "http",
					Host:   "foo.bar.svc.cluster.local",
				},
			},
			Hostname: "this is not empty",
		},
	}
}

// GetURL returns the URL type for the Addressable.
func (a Addressable) GetURL() apis.URL {
	if a.URL != nil {
		return *a.URL
	}
	return apis.URL{
		Scheme: "http",
		Host:   a.Hostname,
	}
}

// GetListType implements apis.Listable
func (*AddressableType) GetListType() runtime.Object {
	return &AddressableTypeList{}
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AddressableTypeList is a list of AddressableType resources
type AddressableTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []AddressableType `json:"items"`
}

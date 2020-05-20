// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"time"

	v1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
	scheme "istio.io/client-go/pkg/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RequestAuthenticationsGetter has a method to return a RequestAuthenticationInterface.
// A group's client should implement this interface.
type RequestAuthenticationsGetter interface {
	RequestAuthentications(namespace string) RequestAuthenticationInterface
}

// RequestAuthenticationInterface has methods to work with RequestAuthentication resources.
type RequestAuthenticationInterface interface {
	Create(*v1beta1.RequestAuthentication) (*v1beta1.RequestAuthentication, error)
	Update(*v1beta1.RequestAuthentication) (*v1beta1.RequestAuthentication, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.RequestAuthentication, error)
	List(opts v1.ListOptions) (*v1beta1.RequestAuthenticationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.RequestAuthentication, err error)
	RequestAuthenticationExpansion
}

// requestAuthentications implements RequestAuthenticationInterface
type requestAuthentications struct {
	client rest.Interface
	ns     string
}

// newRequestAuthentications returns a RequestAuthentications
func newRequestAuthentications(c *SecurityV1beta1Client, namespace string) *requestAuthentications {
	return &requestAuthentications{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the requestAuthentication, and returns the corresponding requestAuthentication object, and an error if there is any.
func (c *requestAuthentications) Get(name string, options v1.GetOptions) (result *v1beta1.RequestAuthentication, err error) {
	result = &v1beta1.RequestAuthentication{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("requestauthentications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RequestAuthentications that match those selectors.
func (c *requestAuthentications) List(opts v1.ListOptions) (result *v1beta1.RequestAuthenticationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.RequestAuthenticationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("requestauthentications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested requestAuthentications.
func (c *requestAuthentications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("requestauthentications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a requestAuthentication and creates it.  Returns the server's representation of the requestAuthentication, and an error, if there is any.
func (c *requestAuthentications) Create(requestAuthentication *v1beta1.RequestAuthentication) (result *v1beta1.RequestAuthentication, err error) {
	result = &v1beta1.RequestAuthentication{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("requestauthentications").
		Body(requestAuthentication).
		Do().
		Into(result)
	return
}

// Update takes the representation of a requestAuthentication and updates it. Returns the server's representation of the requestAuthentication, and an error, if there is any.
func (c *requestAuthentications) Update(requestAuthentication *v1beta1.RequestAuthentication) (result *v1beta1.RequestAuthentication, err error) {
	result = &v1beta1.RequestAuthentication{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("requestauthentications").
		Name(requestAuthentication.Name).
		Body(requestAuthentication).
		Do().
		Into(result)
	return
}

// Delete takes name of the requestAuthentication and deletes it. Returns an error if one occurs.
func (c *requestAuthentications) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("requestauthentications").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *requestAuthentications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("requestauthentications").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched requestAuthentication.
func (c *requestAuthentications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.RequestAuthentication, err error) {
	result = &v1beta1.RequestAuthentication{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("requestauthentications").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

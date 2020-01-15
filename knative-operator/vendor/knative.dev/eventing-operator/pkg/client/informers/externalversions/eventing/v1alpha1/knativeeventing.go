/*
Copyright 2019 The Knative Authors

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	eventingv1alpha1 "knative.dev/eventing-operator/pkg/apis/eventing/v1alpha1"
	versioned "knative.dev/eventing-operator/pkg/client/clientset/versioned"
	internalinterfaces "knative.dev/eventing-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "knative.dev/eventing-operator/pkg/client/listers/eventing/v1alpha1"
)

// KnativeEventingInformer provides access to a shared informer and lister for
// KnativeEventings.
type KnativeEventingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.KnativeEventingLister
}

type knativeEventingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewKnativeEventingInformer constructs a new informer for KnativeEventing type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKnativeEventingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKnativeEventingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredKnativeEventingInformer constructs a new informer for KnativeEventing type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKnativeEventingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1alpha1().KnativeEventings(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1alpha1().KnativeEventings(namespace).Watch(options)
			},
		},
		&eventingv1alpha1.KnativeEventing{},
		resyncPeriod,
		indexers,
	)
}

func (f *knativeEventingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKnativeEventingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *knativeEventingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&eventingv1alpha1.KnativeEventing{}, f.defaultInformer)
}

func (f *knativeEventingInformer) Lister() v1alpha1.KnativeEventingLister {
	return v1alpha1.NewKnativeEventingLister(f.Informer().GetIndexer())
}

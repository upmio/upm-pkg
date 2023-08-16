/*
Copyright The Kubernetes Authors.

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
	"context"
	time "time"

	unitsetv1alpha1 "github.com/upmio/upm-pkg/pkg/apis/unitset/v1alpha1"
	versioned "github.com/upmio/upm-pkg/pkg/client/unitset/clientset/versioned"
	internalinterfaces "github.com/upmio/upm-pkg/pkg/client/unitset/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/upmio/upm-pkg/pkg/client/unitset/listers/unitset/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// UnitsetInformer provides access to a shared informer and lister for
// Unitsets.
type UnitsetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.UnitsetLister
}

type unitsetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewUnitsetInformer constructs a new informer for Unitset type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewUnitsetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredUnitsetInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredUnitsetInformer constructs a new informer for Unitset type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredUnitsetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrdV1alpha1().Unitsets(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrdV1alpha1().Unitsets(namespace).Watch(context.TODO(), options)
			},
		},
		&unitsetv1alpha1.Unitset{},
		resyncPeriod,
		indexers,
	)
}

func (f *unitsetInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredUnitsetInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *unitsetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&unitsetv1alpha1.Unitset{}, f.defaultInformer)
}

func (f *unitsetInformer) Lister() v1alpha1.UnitsetLister {
	return v1alpha1.NewUnitsetLister(f.Informer().GetIndexer())
}

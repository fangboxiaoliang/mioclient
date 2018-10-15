/*
Copyright 2018 The Kubernetes Authors.

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

	mio_v1alpha1 "github.com/hidevopsio/mioclient/pkg/apis/mio/v1alpha1"
	versioned "github.com/hidevopsio/mioclient/pkg/client/clientset/versioned"
	internalinterfaces "github.com/hidevopsio/mioclient/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/hidevopsio/mioclient/pkg/client/listers/mio/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PipelineConfigInformer provides access to a shared informer and lister for
// PipelineConfigs.
type PipelineConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PipelineConfigLister
}

type pipelineConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPipelineConfigInformer constructs a new informer for PipelineConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPipelineConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPipelineConfigInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPipelineConfigInformer constructs a new informer for PipelineConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPipelineConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MioV1alpha1().PipelineConfigs(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MioV1alpha1().PipelineConfigs(namespace).Watch(options)
			},
		},
		&mio_v1alpha1.PipelineConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *pipelineConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPipelineConfigInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *pipelineConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&mio_v1alpha1.PipelineConfig{}, f.defaultInformer)
}

func (f *pipelineConfigInformer) Lister() v1alpha1.PipelineConfigLister {
	return v1alpha1.NewPipelineConfigLister(f.Informer().GetIndexer())
}

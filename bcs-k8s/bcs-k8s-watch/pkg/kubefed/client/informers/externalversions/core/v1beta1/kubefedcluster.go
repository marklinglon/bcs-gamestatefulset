/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	core_v1beta1 "github.com/Tencent/bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/apis/core/v1beta1"
	versioned "github.com/Tencent/bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/client/clientset/versioned"
	internalinterfaces "github.com/Tencent/bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/Tencent/bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/client/listers/core/v1beta1"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KubeFedClusterInformer provides access to a shared informer and lister for
// KubeFedClusters.
type KubeFedClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.KubeFedClusterLister
}

type kubeFedClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewKubeFedClusterInformer constructs a new informer for KubeFedCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKubeFedClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKubeFedClusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredKubeFedClusterInformer constructs a new informer for KubeFedCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKubeFedClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1beta1().KubeFedClusters(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1beta1().KubeFedClusters(namespace).Watch(options)
			},
		},
		&core_v1beta1.KubeFedCluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *kubeFedClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKubeFedClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kubeFedClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&core_v1beta1.KubeFedCluster{}, f.defaultInformer)
}

func (f *kubeFedClusterInformer) Lister() v1beta1.KubeFedClusterLister {
	return v1beta1.NewKubeFedClusterLister(f.Informer().GetIndexer())
}